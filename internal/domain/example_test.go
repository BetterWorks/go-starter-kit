package domain

import (
	"context"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/BetterWorks/go-starter-kit/test/mock"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	"github.com/google/uuid"
)

func Test_Example_NewExampleService(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	_, err := NewExampleService(config)
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_NewExampleService_Error(t *testing.T) {
	t.Parallel()

	config := &ExampleServiceConfig{
		Logger: nil,
		Repo:   nil,
	}
	_, err := NewExampleService(config)
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_Example_Create_Success(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	service, _ := NewExampleService(config)

	_, err := service.Create(context.Background(), fx.NewExampleRequestAttributesBuilder().Build())
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_Create_Error(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{CreateExampleError: true},
	}
	service, _ := NewExampleService(config)

	_, err := service.Create(context.Background(), fx.NewExampleRequestAttributesBuilder().Build())
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_Example_Delete_Success(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	service, _ := NewExampleService(config)

	err := service.Delete(context.Background(), uuid.New())
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_Delete_Error(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{DeleteExampleError: true},
	}
	service, _ := NewExampleService(config)

	err := service.Delete(context.Background(), uuid.New())
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_Example_Detail_Success(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	service, _ := NewExampleService(config)

	_, err := service.Detail(context.Background(), uuid.New())
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_Detail_Error(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{DetailExampleError: true},
	}
	service, _ := NewExampleService(config)

	_, err := service.Detail(context.Background(), uuid.New())
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_Example_List_Success(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	service, _ := NewExampleService(config)

	_, err := service.List(context.Background(), query.QueryData{})
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_List_Error(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{ListExampleError: true},
	}
	service, _ := NewExampleService(config)

	_, err := service.List(context.Background(), query.QueryData{})
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}

func Test_Example_Update_Success(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{},
	}
	service, _ := NewExampleService(config)

	_, err := service.Update(context.Background(), fx.NewExampleRequestAttributesBuilder().Build(), uuid.New())
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}
}

func Test_Example_Update_Error(t *testing.T) {
	t.Parallel()

	newRelicClient, _ := mock.NewRelicClient(t)
	config := &ExampleServiceConfig{
		Logger: &logger.CustomLogger{
			Enabled: false,
			Level:   logger.LevelDebug,
			Log:     mock.Logger(),
		},
		NewRelicClient: newRelicClient,
		Repo:           &mock.ExampleRepository{UpdateExampleError: true},
	}
	service, _ := NewExampleService(config)

	_, err := service.Update(context.Background(), fx.NewExampleRequestAttributesBuilder().Build(), uuid.New())
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}
