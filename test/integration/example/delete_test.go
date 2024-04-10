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

type DeleteSuite struct {
	suite.Suite
	method   string
	handler  http.Handler
	db       *pgxpool.Pool
	resolver *resolver.Resolver
	record   *models.ExampleDomainModel
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, &DeleteSuite{})
}

func (s *DeleteSuite) SetupSuite() {
	server, db, resolver, err := testutils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "DELETE"
	s.handler = server.Server.Handler
	s.db = db
	s.resolver = resolver
}

func (s *DeleteSuite) SetupTest() {
	attrs := fx.NewExampleRequestAttributesBuilder().Build()

	repo := s.resolver.ExampleRepository()
	record, err := repo.Create(context.Background(), attrs)

	if err != nil {
		s.T().Log(err)
	}

	s.record = record
}

func (s *DeleteSuite) TearDownTest() {
	testutils.Cleanup(s.resolver)
}

func (s *DeleteSuite) TestResourceDetail() {
	tests := []testutils.Setup{
		{
			Description: "resource delete succeeds (204)",
			Route:       fmt.Sprintf("/domain/examples/%s", s.record.Data[0].Attributes.ID.String()),
			Request:     testutils.Request{},
			Expected:    testutils.Expected{Code: 204},
		},
	}

	for _, test := range tests {
		req := testutils.SetRequestData(s.method, test.Route, test.Request.Body, nil)
		rec := httptest.NewRecorder()

		s.handler.ServeHTTP(rec, req)

		res := rec.Result()
		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
	}
}
