package exampletest

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/BetterWorks/go-starter-kit/internal/resolver"
// 	"github.com/BetterWorks/go-starter-kit/internal/types"
// 	utils "github.com/BetterWorks/go-starter-kit/test/testutils"
// 	"github.com/stretchr/testify/suite"
// )

// type DetailSuite struct {
// 	suite.Suite
// 	method   string
// 	app      *fiber.App
// 	db       *pgxpool.Pool
// 	resolver *resolver.Resolver
// 	record   *types.ExampleEntity
// }

// func TestDetailSuite(t *testing.T) {
// 	suite.Run(t, &DetailSuite{})
// }

// // SetupSuite runs setup before all suite tests
// func (s *DetailSuite) SetupSuite() {
// 	app, db, resolver, err := utils.InitializeApp(nil)
// 	if err != nil {
// 		s.T().Log(err)
// 	}

// 	s.method = "GET"
// 	s.app = app
// 	s.db = db
// 	s.resolver = resolver
// }

// // TearDownSuite runs teardown after all suite tests
// func (s *DetailSuite) TearDownSuite() {
// 	//
// }

// // SetupTest runs setup before each test
// func (s *DetailSuite) SetupTest() {
// 	record, err := insertRecord(s.db)
// 	if err != nil {
// 		s.T().Log(err)
// 	}
// 	s.record = record
// }

// // TearDownTest runs teardown after each test
// func (s *DetailSuite) TearDownTest() {
// 	utils.Cleanup(s.resolver)
// }

// func (s *DetailSuite) TestResourceDetail() {
// 	tests := []utils.Setup{
// 		{
// 			Description: "resource detail succeeds (200)",
// 			Route:       fmt.Sprintf("%s/%s", routePrefix, s.record.ID.String()),
// 			Request:     utils.Request{},
// 			Expected:    utils.Expected{Code: 200},
// 		},
// 	}

// 	for _, test := range tests {
// 		req := utils.SetRequestData(s.method, test.Route, nil, nil)
// 		msTimeout := 1000

// 		res, err := s.app.Test(req, msTimeout)
// 		if err != nil {
// 			s.T().Log(err)
// 		}

// 		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
// 	}
// }
