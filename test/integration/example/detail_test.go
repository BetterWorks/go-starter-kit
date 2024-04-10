package exampletest

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/http/httpserver"
	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type DetailSuite struct {
	suite.Suite
	method     string
	httpServer *httpserver.Server
	db         *pgxpool.Pool
	resolver   *resolver.Resolver
	record     *models.ExampleDomainModel
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
	s.httpServer = server
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
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		if err := s.httpServer.Serve(); err != nil {
			return err
		}

		return nil
	})
	pollUntilServerStartup(s.T(), 5, 500)
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
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			s.T().Log(err)
		}

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
