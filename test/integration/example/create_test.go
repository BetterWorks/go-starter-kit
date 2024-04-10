package exampletest

import (
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

type CreateSuite struct {
	suite.Suite
	method   string
	handler  http.Handler
	db       *pgxpool.Pool
	resolver *resolver.Resolver
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, &CreateSuite{})
}

func (s *CreateSuite) SetupSuite() {
	server, db, resolver, err := testutils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "POST"
	s.handler = server.Server.Handler
	s.db = db
	s.resolver = resolver
}

func (s *CreateSuite) TearDownTest() {
	testutils.Cleanup(s.resolver)
}

func (s *CreateSuite) TestResourceDetail() {
	tests := []testutils.Setup{
		{
			Description: "resource create succeeds (200)",
			Route:       "/domain/examples/",
			Request: testutils.Request{
				Body: fx.ComposeJSONBody(
					&models.ExampleRequest{
						Data: &models.ExampleRequestResource{
							Attributes: *fx.NewExampleRequestAttributesBuilder().Build(),
						},
					}),
			},
			Expected: testutils.Expected{Code: 201},
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
