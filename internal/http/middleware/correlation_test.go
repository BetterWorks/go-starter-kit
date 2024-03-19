package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

var Error_Validator_NilServerConfig = "validator: (nil *httpserver.ServerConfig)"

func Test_Correlation_Default(t *testing.T) {
	t.Parallel()

	handler := Correlation(&CorrelationConfig{})
	headerName := "X-Request-Id"

	if handler == nil {
		te.NewLineErrorf(t, handler, nil)
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header[headerName][0]
		if val == "" {
			t.Error("traceID not present")
		}
		traceID, err := uuid.Parse(val)

		if traceID.String() == "" || err != nil {
			t.Error("traceID is not a valid uuid")
		}
	})

	handlerToTest := handler(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func Test_Correlation_CustomOverrides(t *testing.T) {
	t.Parallel()

	headerName := fake.Word()
	myUUID := uuid.New().String()
	contextKey := trace.ContextKey(fake.Word())
	next := func(r *http.Request) bool { return false }
	handler := Correlation(&CorrelationConfig{
		Header: headerName,
		Generator: func() string {
			return myUUID
		},
		ContextKey: contextKey,
		Next:       next,
	})

	if handler == nil {
		te.NewLineErrorf(t, handler, nil)
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header[headerName]
		if len(header) == 0 {
			t.Error("traceID header not present")
		}
		val := header[0]
		if val == "" {
			t.Error("traceID not present")
		}
		traceID, err := uuid.Parse(val)

		if traceID.String() != myUUID || err != nil {
			t.Error("traceID is not a valid uuid")
		}
	})

	handlerToTest := handler(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func Test_Correlation_CustomNextSkipsCorrelation(t *testing.T) {
	t.Parallel()

	headerName := "X-Request-Id"
	next := func(r *http.Request) bool { return true }
	handler := Correlation(&CorrelationConfig{Next: next})

	if handler == nil {
		te.NewLineErrorf(t, handler, nil)
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header[headerName]
		if len(header) > 0 {
			t.Error("traceID header present")
		}
	})

	handlerToTest := handler(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}
