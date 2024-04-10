package exampletest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type BaseSuite struct {
	suite.Suite
	method   string
	handler  http.Handler
	db       *pgxpool.Pool
	resolver *resolver.Resolver
}

func TestBaseSuite(t *testing.T) {
	suite.Run(t, &BaseSuite{})
}

func (s *BaseSuite) SetupSuite() {
	server, db, resolver, err := testutils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "GET"
	s.handler = server.Server.Handler
	s.db = db
	s.resolver = resolver
}

func (s *BaseSuite) TearDownTest() {
	testutils.Cleanup(s.resolver)
}

func (s *BaseSuite) TestResourceList() {
	tests := []testutils.Setup{
		{
			Description: "resource list succeeds (200)",
			Route:       "/domain/",
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
