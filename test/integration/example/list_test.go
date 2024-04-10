package exampletest

import (
	"context"
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

type ListSuite struct {
	suite.Suite
	method   string
	handler  http.Handler
	db       *pgxpool.Pool
	resolver *resolver.Resolver
	records  []*models.ExampleDomainModel
}

func TestListSuite(t *testing.T) {
	suite.Run(t, &ListSuite{})
}

func (s *ListSuite) SetupSuite() {
	server, db, resolver, err := testutils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "GET"
	s.handler = server.Server.Handler
	s.db = db
	s.resolver = resolver
}

func (s *ListSuite) SetupTest() {
	records := make([]*models.ExampleDomainModel, 0, 4)
	for i := 0; i < 4; i++ {
		attrs := fx.NewExampleRequestAttributesBuilder().Build()

		repo := s.resolver.ExampleRepository()
		record, err := repo.Create(context.Background(), attrs)

		if err != nil {
			s.T().Log(err)
		}

		records = append(records, record)
	}
	s.records = records
}

func (s *ListSuite) TearDownTest() {
	testutils.Cleanup(s.resolver)
}

func (s *ListSuite) TestResourceList() {
	tests := []testutils.Setup{
		{
			Description: "resource list succeeds (200)",
			Route:       "/domain/examples",
			Request:     testutils.Request{},
			Expected:    testutils.Expected{Code: 200},
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
