package exampletest

import (
	"context"
	"net/http"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/http/httpserver"
	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type BaseSuite struct {
	suite.Suite
	method     string
	httpServer *httpserver.Server
	db         *pgxpool.Pool
	resolver   *resolver.Resolver
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
	s.httpServer = server
	s.db = db
	s.resolver = resolver
}

func (s *BaseSuite) SetupTest() {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		if err := s.httpServer.Serve(); err != nil {
			return err
		}

		return nil
	})
	pollUntilServerStartup(s.T(), 5, 500)
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
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			s.T().Log(err)
		}

		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
	}
}
