package monitor

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRunChecksTargets(t *testing.T) {
	checker := New(2*time.Second, 2, 2)
	checker.client.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("ok")),
			Header:     make(http.Header),
		}, nil
	})
	checker.dial = func(context.Context, string, string) (net.Conn, error) {
		return testConn{}, nil
	}

	results := checker.Run(context.Background(), []string{"https://status.internal.local", "10.10.10.10:443"})

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	for _, result := range results {
		if !result.OK {
			t.Fatalf("expected target %s to be up, got detail %s", result.Target, result.Detail)
		}
	}
}

func TestBuildSummary(t *testing.T) {
	summary := BuildSummary([]Result{
		{OK: true},
		{OK: false},
		{OK: true},
	})

	if summary.Total != 3 || summary.Up != 2 || summary.Down != 1 {
		t.Fatalf("unexpected summary: %+v", summary)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

type testConn struct{}

func (testConn) Read([]byte) (int, error)         { return 0, io.EOF }
func (testConn) Write([]byte) (int, error)        { return 0, nil }
func (testConn) Close() error                     { return nil }
func (testConn) LocalAddr() net.Addr              { return testAddr("local") }
func (testConn) RemoteAddr() net.Addr             { return testAddr("remote") }
func (testConn) SetDeadline(time.Time) error      { return nil }
func (testConn) SetReadDeadline(time.Time) error  { return nil }
func (testConn) SetWriteDeadline(time.Time) error { return nil }

type testAddr string

func (addr testAddr) Network() string { return "tcp" }
func (addr testAddr) String() string  { return string(addr) }
