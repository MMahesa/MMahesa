package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Target   string        `json:"target"`
	Kind     string        `json:"kind"`
	OK       bool          `json:"ok"`
	Latency  time.Duration `json:"latency"`
	Attempts int           `json:"attempts"`
	Detail   string        `json:"detail"`
}

type Summary struct {
	Total int `json:"total"`
	Up    int `json:"up"`
	Down  int `json:"down"`
}

type Target struct {
	Address string `json:"address"`
}

type Config struct {
	Targets []Target `json:"targets"`
}

type Checker struct {
	timeout     time.Duration
	retries     int
	concurrency int
	client      *http.Client
	dial        func(context.Context, string, string) (net.Conn, error)
}

func New(timeout time.Duration, retries, concurrency int) *Checker {
	if retries < 1 {
		retries = 1
	}
	if concurrency < 1 {
		concurrency = 1
	}

	dialer := net.Dialer{Timeout: timeout}
	return &Checker{
		timeout:     timeout,
		retries:     retries,
		concurrency: concurrency,
		client: &http.Client{
			Timeout: timeout,
		},
		dial: dialer.DialContext,
	}
}

func (c *Checker) Run(ctx context.Context, targets []string) []Result {
	results := make([]Result, len(targets))
	type job struct {
		index  int
		target string
	}

	jobs := make(chan job)
	var wg sync.WaitGroup

	for range make([]struct{}, c.concurrency) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results[job.index] = c.check(ctx, job.target)
			}
		}()
	}

	for index, target := range targets {
		jobs <- job{index: index, target: target}
	}
	close(jobs)
	wg.Wait()
	return results
}

func (c *Checker) check(ctx context.Context, target string) Result {
	var result Result
	for attempt := 1; attempt <= c.retries; attempt++ {
		if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
			result = c.checkHTTP(ctx, target, attempt)
		} else {
			result = c.checkTCP(ctx, target, attempt)
		}
		if result.OK {
			return result
		}
	}
	return result
}

func (c *Checker) checkHTTP(ctx context.Context, target string, attempt int) Result {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return Result{Target: target, Kind: "http", Attempts: attempt, Detail: err.Error()}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Result{Target: target, Kind: "http", Attempts: attempt, Latency: time.Since(start), Detail: err.Error()}
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	ok := resp.StatusCode >= 200 && resp.StatusCode < 400
	return Result{
		Target:   target,
		Kind:     "http",
		OK:       ok,
		Latency:  time.Since(start),
		Attempts: attempt,
		Detail:   fmt.Sprintf("status=%d", resp.StatusCode),
	}
}

func (c *Checker) checkTCP(ctx context.Context, target string, attempt int) Result {
	start := time.Now()
	conn, err := c.dial(ctx, "tcp", target)
	if err != nil {
		return Result{Target: target, Kind: "tcp", Attempts: attempt, Latency: time.Since(start), Detail: err.Error()}
	}
	_ = conn.Close()

	return Result{
		Target:   target,
		Kind:     "tcp",
		OK:       true,
		Latency:  time.Since(start),
		Attempts: attempt,
		Detail:   "connection established",
	}
}

func BuildSummary(results []Result) Summary {
	summary := Summary{Total: len(results)}
	for _, result := range results {
		if result.OK {
			summary.Up++
		} else {
			summary.Down++
		}
	}
	return summary
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
