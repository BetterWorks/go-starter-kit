package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

var nextHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("next handler called")
})

func Test_RequestLogger(t *testing.T) {
	t.Parallel()

	var (
		buff       bytes.Buffer
		testLogger = slog.New(slog.NewJSONHandler(&buff, nil))
	)

	cLogger := &logger.CustomLogger{
		Enabled: true,
		Level:   "debug",
		Log:     testLogger,
	}

	config := &RequestLoggerConfig{Logger: cLogger, Next: nil}

	// create the handler to test, using our custom "next" handler
	handlerToTest := RequestLogger(config)

	bodyObj := map[string]interface{}{
		fake.Word(): fake.Word(),
	}
	method := "POST"
	path := fmt.Sprintf("/%s", fake.Word())
	traceID := fake.UUID()

	// create a mock request to use
	req := buildRequest(method, path, bodyObj, traceID)

	w := httptest.NewRecorder()
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest(nextHandler).ServeHTTP(w, req)

	// check the log output
	logEntries := parseLogEntries(t, buff.Bytes())
	if len(logEntries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(logEntries))
	}

	entry := logEntries[0]
	if entry["level"] != "INFO" {
		te.NewLineErrorf(t, "INFO", entry["level"])
	}

	if entry["msg"] != "request" {
		te.NewLineErrorf(t, "request", entry["msg"])
	}

	if entry["ip"] != req.RemoteAddr {
		te.NewLineErrorf(t, req.RemoteAddr, entry["ip"])
	}

	if entry["method"] != method {
		te.NewLineErrorf(t, method, entry["method"])
	}

	if entry["path"] != path {
		te.NewLineErrorf(t, path, entry["path"])
	}

	if entry["trace_id"] != traceID {
		te.NewLineErrorf(t, traceID, entry["trace_id"])
	}

	assertHeaders(t, entry, req)
	assertBody(t, entry, bodyObj)
}

func Test_RequestLogger_Info(t *testing.T) {
	t.Parallel()

	var (
		buff       bytes.Buffer
		testLogger = slog.New(slog.NewJSONHandler(&buff, nil))
	)

	cLogger := &logger.CustomLogger{
		Enabled: true,
		Level:   "info",
		Log:     testLogger,
	}

	config := &RequestLoggerConfig{Logger: cLogger, Next: nil}

	// create the handler to test, using our custom "next" handler
	handlerToTest := RequestLogger(config)

	bodyObj := map[string]interface{}{
		fake.Word(): fake.Word(),
	}
	method := "POST"
	path := fmt.Sprintf("/%s", fake.Word())
	traceID := fake.UUID()
	// create a mock request to use
	req := buildRequest(method, path, bodyObj, traceID)

	w := httptest.NewRecorder()
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest(nextHandler).ServeHTTP(w, req)

	// check the log output
	logEntries := parseLogEntries(t, buff.Bytes())
	if len(logEntries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(logEntries))
	}

	entry := logEntries[0]
	if entry["level"] != "INFO" {
		te.NewLineErrorf(t, "INFO", entry["level"])
	}

	if entry["msg"] != "request" {
		te.NewLineErrorf(t, "request", entry["msg"])
	}

	if entry["ip"] != req.RemoteAddr {
		te.NewLineErrorf(t, req.RemoteAddr, entry["ip"])
	}

	if entry["method"] != method {
		te.NewLineErrorf(t, method, entry["method"])
	}

	if entry["path"] != path {
		te.NewLineErrorf(t, path, entry["path"])
	}

	if entry["trace_id"] != traceID {
		te.NewLineErrorf(t, traceID, entry["trace_id"])
	}

	if entry["headers"] != nil {
		te.NewLineErrorf(t, nil, entry["headers"])
	}

	if entry["body"] != nil {
		te.NewLineErrorf(t, nil, entry["body"])
	}
}

func assertBody(t *testing.T, entry map[string]any, bodyObj map[string]interface{}) {
	if entry["body"] == nil {
		te.NewLineErrorf(t, "body", nil)
	}

	actualBody, _ := json.Marshal(entry["body"].(map[string]any))
	var bodyCopy map[string]any
	if err := json.Unmarshal(actualBody, &bodyCopy); err != nil {
		t.Fatal(err)
	}

	if len(bodyCopy) != len(bodyObj) {
		te.NewLineErrorf(t, len(bodyObj), len(bodyCopy))
	}

	for k, v := range bodyCopy {
		if v != bodyObj[k] {
			te.NewLineErrorf(t, bodyObj[k], v)
		}
	}
}

func assertHeaders(t *testing.T, entry map[string]any, req *http.Request) {
	if entry["headers"] == nil {
		te.NewLineErrorf(t, "headers", nil)
	}

	headers := entry["headers"].(map[string]any)

	for k, v := range headers {
		if v.([]any)[0] != req.Header.Get(k) {
			te.NewLineErrorf(t, req.Header.Get(k), v)
		}
	}
}

func buildRequest(method string, path string, bodyObj map[string]interface{}, traceID string) *http.Request {
	body, _ := json.Marshal(bodyObj)

	req := httptest.NewRequest(method, fmt.Sprintf("http://testing%s", path), bytes.NewReader(body))
	ctx := trace.CreateOpContext(context.Background(), traceID)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(fake.Word(), fake.Word())
	req.RemoteAddr = fmt.Sprintf("%s:%d", fake.IPv4Address(), fake.Number(1000, 9999))
	return req
}

func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	t.Helper()

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()

	var entries []map[string]any
	for {
		var entry map[string]any
		if err := dec.Decode(&entry); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		entries = append(entries, entry)
	}

	if len(entries) == 0 {
		t.Fatal("no log entries found")
	}

	return entries
}
