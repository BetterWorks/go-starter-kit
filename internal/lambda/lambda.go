package lambda

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/aws/aws-lambda-go/events"
	lambdaSDK "github.com/aws/aws-lambda-go/lambda"
)

type LambdaServiceConfig struct {
	Logger *logger.CustomLogger `validate:"required"`
}

type LambdaService struct {
	logger *logger.CustomLogger
}

type LambdaEvent struct {
	EventType string `json:"event_type"`
}

// NewExampleService returns a new exampleService instance
func NewLambdaService(c *LambdaServiceConfig) (*LambdaService, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	service := &LambdaService{
		logger: c.Logger,
	}

	return service, nil
}

func (s *LambdaService) Start() {
	lambdaSDK.Start(s.baseLambdaHandler)
}

func (s *LambdaService) baseLambdaHandler(e LambdaEvent) (events.APIGatewayProxyResponse, error) {
	switch e.EventType {
	case "example":
		return s.exampleLambdaHandler(e)
	default:
		return s.unhandledMethod(e)
	}
}

func (s *LambdaService) exampleLambdaHandler(e LambdaEvent) (events.APIGatewayProxyResponse, error) {
	s.logger.Log.Info(fmt.Sprintf("Received event: %v", e))

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}

	return res, nil
}

func (s *LambdaService) unhandledMethod(e LambdaEvent) (events.APIGatewayProxyResponse, error) {
	s.logger.Log.Info(fmt.Sprintf("Received event: %v", e))
	var errorMethodNotAllowed = http.StatusText(http.StatusMethodNotAllowed)

	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       errorMethodNotAllowed,
	}

	return res, errors.New(errorMethodNotAllowed)
}
