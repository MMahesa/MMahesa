package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MMahesa/uptime-probe/internal/monitor"
)

func main() {
	timeout := flag.Duration("timeout", 3*time.Second, "request timeout")
	format := flag.String("format", "table", "output format: table or json")
	file := flag.String("file", "", "optional file containing one target per line")
	configPath := flag.String("config", "", "optional JSON config file")
	retries := flag.Int("retries", 2, "number of attempts per target")
	concurrency := flag.Int("concurrency", 4, "number of concurrent workers")
	flag.Parse()

	targets := flag.Args()
	if *configPath != "" {
		config, err := monitor.LoadConfig(*configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, target := range config.Targets {
			targets = append(targets, strings.TrimSpace(target.Address))
		}
	}
	if *file != "" {
		loaded, err := loadTargets(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		targets = append(targets, loaded...)
	}
	if len(targets) == 0 {
		fmt.Fprintln(os.Stderr, "usage: uptime-probe [--timeout=3s] [--format=table|json] [--file=targets.txt] <target> [target...]")
		os.Exit(1)
	}

	checker := monitor.New(*timeout, *retries, *concurrency)
	results := checker.Run(context.Background(), targets)
	summary := monitor.BuildSummary(results)

	switch strings.ToLower(strings.TrimSpace(*format)) {
	case "json":
		if err := json.NewEncoder(os.Stdout).Encode(map[string]any{
			"summary": summary,
			"results": results,
		}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		printTable(results)
		printSummary(summary)
	}

	if summary.Down > 0 {
		os.Exit(2)
	}
}

func printTable(results []monitor.Result) {
	fmt.Printf("%-8s %-7s %-28s %-10s %-8s %s\n", "TYPE", "STATUS", "TARGET", "LATENCY", "ATTEMPTS", "DETAIL")
	for _, result := range results {
		status := "down"
		if result.OK {
			status = "up"
		}
		fmt.Printf("%-8s %-7s %-28s %-10s %-8d %s\n",
			result.Kind,
			status,
			result.Target,
			result.Latency.Round(time.Millisecond),
			result.Attempts,
			result.Detail,
		)
	}
}

func printSummary(summary monitor.Summary) {
	fmt.Println()
	fmt.Printf("Summary: total=%d up=%d down=%d\n", summary.Total, summary.Up, summary.Down)
}

func loadTargets(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var targets []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		targets = append(targets, line)
	}
	return targets, scanner.Err()
}
