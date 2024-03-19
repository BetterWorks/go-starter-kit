package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/go-chi/chi/v5"
)

type testRoutes struct {
}

func (t *testRoutes) Routes() []chi.Route {
	return []chi.Route{}
}

func (t *testRoutes) Middlewares() chi.Middlewares {
	return chi.Middlewares{}
}

func (t *testRoutes) Match(rctx *chi.Context, method, path string) bool {
	return path == "/match"
}

func Test_NotFound(t *testing.T) {
	t.Parallel()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("next handler called"))
	})

	handlerToTest := NotFound(nextHandler)

	req := getRequest(fake.Word())
	w := httptest.NewRecorder()

	handlerToTest.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 404 {
		te.NewLineErrorf(t, 404, resp.StatusCode)
	}

	strBody := strings.ReplaceAll(string(body), "\n", "")
	expectedErrorBody := fmt.Sprintf(`{"errors":[{"status":%d,"title":"%s","detail":"%s"}]}`, 404, cerror.ErrorType.NotFound, "path not found")

	if strBody != expectedErrorBody {
		te.NewLineErrorf(t, expectedErrorBody, strBody)
	}
}

func Test_NotFound_WithMatch(t *testing.T) {
	t.Parallel()

	expectedBodyMessage := "next handler called"

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBodyMessage))
	})

	// create the handler to test, using our custom "next" handler
	handlerToTest := NotFound(nextHandler)

	// create a mock request to use
	req := getRequest("match")

	w := httptest.NewRecorder()

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		te.NewLineErrorf(t, 200, resp.StatusCode)
	}

	strBody := strings.ReplaceAll(string(body), "\n", "")
	if strBody != expectedBodyMessage {
		te.NewLineErrorf(t, expectedBodyMessage, strBody)
	}
}

func getRequest(path string) *http.Request {
	req := httptest.NewRequest("GET", fmt.Sprintf("http://testing/%s", path), nil)
	routeContext := chi.NewRouteContext()
	routeContext.Routes = &testRoutes{}

	context := context.WithValue(context.Background(), chi.RouteCtxKey, routeContext)
	req = req.WithContext(context)
	return req
}
