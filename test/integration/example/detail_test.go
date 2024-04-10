package exampletest

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type DetailSuite struct {
	suite.Suite
	method   string
	handler  http.Handler
	db       *pgxpool.Pool
	resolver *resolver.Resolver
	record   *models.ExampleDomainModel
}

func TestDetailSuite(t *testing.T) {
	suite.Run(t, &DetailSuite{})
}

func (s *DetailSuite) SetupSuite() {
	server, db, resolver, err := testutils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "GET"
	s.handler = server.Server.Handler
	s.db = db
	s.resolver = resolver
}

func (s *DetailSuite) SetupTest() {
	attrs := fx.NewExampleRequestAttributesBuilder().Build()

	repo := s.resolver.ExampleRepository()
	record, err := repo.Create(context.Background(), attrs)

	if err != nil {
		s.T().Log(err)
	}

	s.record = record
}

func (s *DetailSuite) TearDownTest() {
	testutils.Cleanup(s.resolver)
}

func (s *DetailSuite) TestResourceDetail() {
	tests := []testutils.Setup{
		{
			Description: "resource detail succeeds (200)",
			Route:       fmt.Sprintf("/domain/examples/%s", s.record.Data[0].Attributes.ID.String()),
			Request:     testutils.Request{},
			Expected:    testutils.Expected{Code: 200},
		},
	}

	for _, test := range tests {
		req := testutils.SetRequestData(s.method, test.Route, test.Request.Body, nil)
		rec := httptest.NewRecorder()

		s.handler.ServeHTTP(rec, req)

		res := rec.Result()

		// b, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	s.T().Error(err)
		// }
		// responseBody := string(b)

		// expectedResponseBody := fmt.Sprintf("%s\n", fx.ComposeJSONBody(
		// 	fx.NewJSONAPIErrorResponseBuilder().
		// 		Errors([]models.ErrorData{*fx.NewJSONAPIErrorDataBuilder().
		// 			Detail("error parsing resource id").
		// 			Build()}).
		// 		Build(),
		// ).String())

		// fmt.Printf("res: %v\n", responseBody)

		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
	}
}
