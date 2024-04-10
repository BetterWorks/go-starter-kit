package exampletest

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/http/httpserver"
	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type DeleteSuite struct {
	suite.Suite
	method     string
	httpServer *httpserver.Server
	db         *pgxpool.Pool
	resolver   *resolver.Resolver
	record     *models.ExampleDomainModel
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
	s.httpServer = server
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
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		if err := s.httpServer.Serve(); err != nil {
			return err
		}

		return nil
	})
	s.T().Log("polling until health check passes")
	assert.Eventually(s.T(), func() bool {
		req := testutils.SetRequestData("GET", "/domain/health", nil, nil)
		if res, err := http.DefaultClient.Do(req); err != nil {
			s.T().Log(err)
		} else if res.StatusCode == 200 {
			return true
		}
		return false
	}, 5*time.Second, 500*time.Millisecond)
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
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			s.T().Log(err)
		}

		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
	}
}
