package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/BetterWorks/go-starter-kit/test/mock"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	htu "github.com/BetterWorks/go-starter-kit/test/testutils/http"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Test_HTTPControllers_Example_NewExampleController(t *testing.T) {
	t.Parallel()

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: &mock.ExampleService{},
	}
	_, err := NewExampleController(config)
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}

	config = &ExampleControllerConfig{
		Logger:      nil,
		QueryConfig: nil,
		Service:     mock.NewExampleService(nil),
	}
	_, err = NewExampleController(config)
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_HTTPControllers_Example_NewExampleController_Error(t *testing.T) {
	t.Parallel()

	config := &ExampleControllerConfig{
		Logger: nil,
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: nil,
	}
	_, err := NewExampleController(config)
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_HTTPControllers_Example_List_Success(t *testing.T) {
	t.Parallel()

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: mock.NewExampleService(nil),
	}
	ctrl, err := NewExampleController(config)
	if err != nil {
		t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
	}
	handler := ctrl.List()

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	expected := htu.Expected{
		ResponseBody: "{\"meta\":{\"page\":{\"limit\":0,\"offset\":0,\"total\":0}},\"data\":[]}\n",
		StatusCode:   http.StatusOK,
	}

	if res.StatusCode != expected.StatusCode {
		te.NewLineErrorf(t, expected.StatusCode, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	responseBody := string(b)

	if responseBody != expected.ResponseBody {
		te.NewLineErrorf(t, expected.ResponseBody, responseBody)
	}
}

func Test_HTTPControllers_Example_List_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		handler     func() http.HandlerFunc
		expected    htu.Expected
	}{{
		name:        "ServiceError",
		description: "tenant service returns an internal server error",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: &mock.ExampleService{
					ListExampleError: true,
				},
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.List()
		},
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(fx.NewJSONAPIErrorResponseBuilder().Build()).String()),
			StatusCode:   http.StatusInternalServerError,
		},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			handler := tc.handler()

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatal(err)
			}

			handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expected.StatusCode {
				te.NewLineErrorf(t, tc.expected.StatusCode, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			responseBody := string(b)

			if responseBody != tc.expected.ResponseBody {
				te.NewLineErrorf(t, tc.expected.ResponseBody, responseBody)
			}
		})
	}
}

func Test_HTTPControllers_Example_Detail_Success(t *testing.T) {
	t.Parallel()

	mockService := mock.NewExampleService(nil)

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: mockService,
	}
	ctrl, err := NewExampleController(config)
	if err != nil {
		t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
	}
	handler := ctrl.Detail()

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New().String()
	req = AddChiURLParams(req, map[string]string{"id": id})

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	expected := htu.Expected{
		ResponseBody: BuildExampleObjectResponseBody(mockService.DetailResult.Data[0]),
		StatusCode:   http.StatusOK,
	}

	if res.StatusCode != expected.StatusCode {
		te.NewLineErrorf(t, expected.StatusCode, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	responseBody := string(b)

	if responseBody != expected.ResponseBody {
		te.NewLineErrorf(t, expected.ResponseBody, responseBody)
	}
}

func Test_HTTPControllers_Example_Detail_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		id          string
		handler     func() http.HandlerFunc
		expected    htu.Expected
	}{{
		name:        "ServiceError",
		description: "example service returns an internal server error",
		id:          uuid.New().String(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: &mock.ExampleService{
					DetailExampleError: true,
				},
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Detail()
		},
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(fx.NewJSONAPIErrorResponseBuilder().Build()).String()),
			StatusCode:   http.StatusInternalServerError,
		},
	}, {
		name:        "InvalidUUID",
		description: "example service returns an invalid UUID error",
		id:          "invalid-uuid",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: &mock.ExampleService{
					DetailExampleError: true,
				},
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Detail()
		},
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("error parsing resource id").
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusInternalServerError,
		},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			handler := tc.handler()

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Error(err)
			}
			req = AddChiURLParams(req, map[string]string{"id": tc.id})

			handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expected.StatusCode {
				te.NewLineErrorf(t, tc.expected.StatusCode, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			responseBody := string(b)

			if responseBody != tc.expected.ResponseBody {
				te.NewLineErrorf(t, tc.expected.ResponseBody, responseBody)
			}
		})
	}
}

func Test_HTTPControllers_Example_Create_Success(t *testing.T) {
	t.Parallel()
	createResult := models.ExampleDomainModel{Data: []models.ExampleObject{fx.NewExampleObjectBuilder().Build()}, Solo: true}

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: mock.NewExampleService(&mock.ExampleService{CreateResult: &createResult}),
	}
	ctrl, err := NewExampleController(config)
	if err != nil {
		t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
	}
	handler := ctrl.Create()

	rd := &htu.RequestData{
		Body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{
					Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
				},
			},
		),
		Method: http.MethodPost,
	}

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "", rd.Body)
	if err != nil {
		t.Fatal(err)
	}

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	expected := htu.Expected{
		ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(createResult.FormatResponse()).String()),
		StatusCode:   http.StatusCreated,
	}

	if res.StatusCode != expected.StatusCode {
		te.NewLineErrorf(t, expected.StatusCode, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	responseBody := string(b)

	if responseBody != expected.ResponseBody {
		te.NewLineErrorf(t, expected.ResponseBody, responseBody)
	}
}

func Test_HTTPControllers_Example_Create_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		handler     func() http.HandlerFunc
		body        io.Reader
		expected    htu.Expected
	}{{
		name:        "ValidationError",
		description: "example service returns a validation error",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(nil),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Create()
		},
		body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{},
			},
		),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("cannot be blank").
						Code(http.StatusBadRequest).
						Title("ValidationError").
						Source(&models.ErrorSource{Pointer: "/title"}).
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusBadRequest,
		},
	}, {
		name:        "ServiceError",
		description: "example service returns an internal server error",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					CreateExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Create()
		},
		body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{
					Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
				},
			},
		),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(fx.NewJSONAPIErrorResponseBuilder().Build()).String()),
			StatusCode:   http.StatusInternalServerError,
		},
	}, {
		name:        "Unparseable Body",
		description: "example service returns a validation error",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					CreateExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Create()
		},
		body: fx.ComposeJSONBody(fx.NewExampleRequestAttributesBuilder().Build()),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("request body decode error").
						Code(http.StatusBadRequest).
						Title("ValidationError").
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusBadRequest,
		},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			handler := tc.handler()

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "", tc.body)
			if err != nil {
				t.Error(err)
			}

			handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expected.StatusCode {
				te.NewLineErrorf(t, tc.expected.StatusCode, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			responseBody := string(b)

			if responseBody != tc.expected.ResponseBody {
				te.NewLineErrorf(t, tc.expected.ResponseBody, responseBody)
			}
		})
	}
}

func Test_HTTPControllers_Example_Update_Success(t *testing.T) {
	t.Parallel()
	updateResult := models.ExampleDomainModel{Data: []models.ExampleObject{fx.NewExampleObjectBuilder().Build()}, Solo: true}

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: mock.NewExampleService(&mock.ExampleService{UpdateResult: &updateResult}),
	}
	ctrl, err := NewExampleController(config)
	if err != nil {
		t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
	}
	handler := ctrl.Update()

	rd := &htu.RequestData{
		Body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{
					Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
				},
			},
		),
		Method: http.MethodPatch,
	}

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(rd.Method, "", rd.Body)
	if err != nil {
		t.Fatal(err)
	}

	id := uuid.New().String()
	req = AddChiURLParams(req, map[string]string{"id": id})

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	expected := htu.Expected{
		ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(updateResult.FormatResponse()).String()),
		StatusCode:   http.StatusOK,
	}

	if res.StatusCode != expected.StatusCode {
		te.NewLineErrorf(t, expected.StatusCode, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	responseBody := string(b)

	if responseBody != expected.ResponseBody {
		te.NewLineErrorf(t, expected.ResponseBody, responseBody)
	}
}

func Test_HTTPControllers_Example_Update_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		id          string
		handler     func() http.HandlerFunc
		body        io.Reader
		expected    htu.Expected
	}{{
		name:        "ValidationError invalid uuid for id",
		description: "unparseable uuid for id",
		id:          fake.Word(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(nil),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Update()
		},
		body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{
					Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
				},
			},
		),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("resource id parse error").
						Code(http.StatusBadRequest).
						Title("ValidationError").
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusBadRequest,
		},
	}, {
		name:        "ValidationError",
		description: "example service returns a validation error",
		id:          uuid.New().String(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(nil),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Update()
		},
		body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{},
			},
		),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("cannot be blank").
						Code(http.StatusBadRequest).
						Title("ValidationError").
						Source(&models.ErrorSource{Pointer: "/title"}).
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusBadRequest,
		},
	}, {
		name:        "ServiceError",
		description: "example service returns an internal server error",
		id:          uuid.New().String(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					UpdateExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Update()
		},
		body: fx.ComposeJSONBody(
			&models.ExampleRequest{
				Data: &models.ExampleRequestResource{
					Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
				},
			},
		),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(fx.NewJSONAPIErrorResponseBuilder().Build()).String()),
			StatusCode:   http.StatusInternalServerError,
		},
	}, {
		name:        "Unparseable Body",
		description: "example service returns a validation error",
		id:          uuid.New().String(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					UpdateExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Update()
		},
		body: fx.ComposeJSONBody(fx.NewExampleRequestAttributesBuilder().Build()),
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("request body decode error").
						Code(http.StatusBadRequest).
						Title("ValidationError").
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusBadRequest,
		},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			handler := tc.handler()

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPatch, "", tc.body)
			if err != nil {
				t.Error(err)
			}
			req = AddChiURLParams(req, map[string]string{"id": tc.id})

			handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expected.StatusCode {
				te.NewLineErrorf(t, tc.expected.StatusCode, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			responseBody := string(b)

			if responseBody != tc.expected.ResponseBody {
				te.NewLineErrorf(t, tc.expected.ResponseBody, responseBody)
			}
		})
	}
}

func Test_HTTPControllers_Example_Delete_Success(t *testing.T) {
	t.Parallel()

	config := &ExampleControllerConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		QueryConfig: &QueryConfig{
			Defaults: &QueryDefaults{
				Paging:  &query.QueryPaging{},
				Sorting: &query.QuerySorting{},
			},
		},
		Service: mock.NewExampleService(nil),
	}
	ctrl, err := NewExampleController(config)
	if err != nil {
		t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
	}
	handler := ctrl.Delete()

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodDelete, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New().String()
	req = AddChiURLParams(req, map[string]string{"id": id})

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	expected := htu.Expected{
		ResponseBody: "",
		StatusCode:   http.StatusNoContent,
	}

	if res.StatusCode != expected.StatusCode {
		te.NewLineErrorf(t, expected.StatusCode, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	responseBody := string(b)

	if responseBody != expected.ResponseBody {
		te.NewLineErrorf(t, expected.ResponseBody, responseBody)
	}
}

func Test_HTTPControllers_Example_Delete_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		id          string
		handler     func() http.HandlerFunc
		expected    htu.Expected
	}{{
		name:        "ServiceError",
		description: "example service returns an internal server error",
		id:          uuid.New().String(),
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					DeleteExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Delete()
		},
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(fx.NewJSONAPIErrorResponseBuilder().Build()).String()),
			StatusCode:   http.StatusInternalServerError,
		},
	}, {
		name:        "InvalidUUID",
		description: "example service returns an invalid UUID error",
		id:          "invalid-uuid",
		handler: func() http.HandlerFunc {
			config := &ExampleControllerConfig{
				Logger: &logger.CustomLogger{
					Enabled: false,
					Level:   logger.LevelDebug,
					Log:     mock.Logger(),
				},
				QueryConfig: &QueryConfig{
					Defaults: &QueryDefaults{
						Paging:  &query.QueryPaging{},
						Sorting: &query.QuerySorting{},
					},
				},
				Service: mock.NewExampleService(&mock.ExampleService{
					DeleteExampleError: true,
				}),
			}
			ctrl, err := NewExampleController(config)
			if err != nil {
				t.Fatalf("\nExpected:\nNo error\nActual:\n'%s'", err.Error())
			}
			return ctrl.Delete()
		},
		expected: htu.Expected{
			ResponseBody: fmt.Sprintf("%s\n", fx.ComposeJSONBody(
				fx.NewJSONAPIErrorResponseBuilder().
					Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
						Detail("error parsing resource id").
						Build()}).
					Build(),
			).String()),
			StatusCode: http.StatusInternalServerError,
		},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			handler := tc.handler()

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodDelete, "", nil)
			if err != nil {
				t.Error(err)
			}
			req = AddChiURLParams(req, map[string]string{"id": tc.id})

			handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expected.StatusCode {
				te.NewLineErrorf(t, tc.expected.StatusCode, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			responseBody := string(b)

			if responseBody != tc.expected.ResponseBody {
				te.NewLineErrorf(t, tc.expected.ResponseBody, responseBody)
			}
		})
	}
}

func AddChiURLParams(r *http.Request, params map[string]string) *http.Request {
	ctx := chi.NewRouteContext()
	for k, v := range params {
		ctx.URLParams.Add(k, v)
	}

	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
}

func BuildExampleObjectResponseBody(obj models.ExampleObject) string {
	return fmt.Sprintf(
		"{\"meta\":null,\"data\":{\"type\":\"example\",\"id\":\"%s\",\"attributes\":{\"title\":\"%s\",\"description\":\"%s\",\"status\":%d,\"enabled\":%t,\"created_on\":\"%s\",\"created_by\":%d,\"modified_on\":\"%s\",\"modified_by\":%d}}}\n",
		obj.Attributes.ID.String(),
		obj.Attributes.Title,
		*obj.Attributes.Description,
		*obj.Attributes.Status,
		obj.Attributes.Enabled,
		obj.Attributes.CreatedOn.Format(time.RFC3339Nano),
		obj.Attributes.CreatedBy,
		obj.Attributes.ModifiedOn.Format(time.RFC3339Nano),
		*obj.Attributes.ModifiedBy,
	)
}
