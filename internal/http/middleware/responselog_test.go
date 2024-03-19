package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

func Test_ResponseLogger(t *testing.T) {
	t.Parallel()
	bodyObj := map[string]interface{}{
		fake.Word(): fake.Word(),
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(bodyObj)
		w.Write(body)
		w.Header().Set(fake.Word(), fake.Word())
	})

	var (
		buff       bytes.Buffer
		testLogger = slog.New(slog.NewJSONHandler(&buff, nil))
	)

	cLogger := &logger.CustomLogger{
		Enabled: true,
		Level:   "debug",
		Log:     testLogger,
	}

	config := &ResponseLoggerConfig{Logger: cLogger, Next: nil}

	handlerToTest := ResponseLogger(config)

	method := "POST"
	path := fmt.Sprintf("/%s", fake.Word())
	traceID := fake.UUID()

	req := buildRequest(method, path, nil, traceID)
	w := httptest.NewRecorder()

	handlerToTest(nextHandler).ServeHTTP(w, req)

	logEntries := parseLogEntries(t, buff.Bytes())
	if len(logEntries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(logEntries))
	}

	entry := logEntries[0]
	if entry["level"] != "INFO" {
		te.NewLineErrorf(t, "INFO", entry["level"])
	}

	if entry["msg"] != "response" {
		te.NewLineErrorf(t, "response", entry["msg"])
	}

	if entry["response_time"] == nil {
		te.NewLineErrorf(t, "response_time", nil)
	}

	if entry["status"] != json.Number("200") {
		te.NewLineErrorf(t, json.Number("200"), entry["status"])
	}

	if entry["time"] == nil {
		te.NewLineErrorf(t, "time", nil)
	}

	if entry["trace_id"] != traceID {
		te.NewLineErrorf(t, traceID, entry["trace_id"])
	}

	assertHeadersFromWriter(t, entry, w)
	assertBody(t, entry, bodyObj)
}

func Test_ResponseLogger_Info(t *testing.T) {
	t.Parallel()
	bodyObj := map[string]interface{}{
		fake.Word(): fake.Word(),
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(bodyObj)
		w.Write(body)
		w.Header().Set(fake.Word(), fake.Word())
	})

	var (
		buff       bytes.Buffer
		testLogger = slog.New(slog.NewJSONHandler(&buff, nil))
	)

	cLogger := &logger.CustomLogger{
		Enabled: true,
		Level:   "info",
		Log:     testLogger,
	}

	config := &ResponseLoggerConfig{Logger: cLogger, Next: nil}

	handlerToTest := ResponseLogger(config)

	method := "POST"
	path := fmt.Sprintf("/%s", fake.Word())
	traceID := fake.UUID()

	req := buildRequest(method, path, nil, traceID)
	w := httptest.NewRecorder()

	handlerToTest(nextHandler).ServeHTTP(w, req)

	logEntries := parseLogEntries(t, buff.Bytes())
	if len(logEntries) != 1 {
		te.NewLineErrorf(t, 1, len(logEntries))
	}

	entry := logEntries[0]
	if entry["level"] != "INFO" {
		te.NewLineErrorf(t, "INFO", entry["level"])
	}

	if entry["msg"] != "response" {
		te.NewLineErrorf(t, "response", entry["msg"])
	}

	if entry["response_time"] == nil {
		te.NewLineErrorf(t, "response_time", nil)
	}

	if entry["status"] != json.Number("200") {
		te.NewLineErrorf(t, json.Number("200"), entry["status"])
	}

	if entry["time"] == nil {
		te.NewLineErrorf(t, "time", nil)
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

func assertHeadersFromWriter(t *testing.T, entry map[string]any, w http.ResponseWriter) {
	if entry["headers"] == nil {
		te.NewLineErrorf(t, "headers", nil)
	}

	headers := entry["headers"].(map[string]any)

	for k, v := range headers {
		if v.([]any)[0] != w.Header().Get(k) {
			te.NewLineErrorf(t, w.Header().Get(k), v)
		}
	}
}
