package exampletest

// import (
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/BetterWorks/go-starter-kit/internal/resolver"
// 	"github.com/BetterWorks/go-starter-kit/internal/types"
// 	utils "github.com/BetterWorks/go-starter-kit/test/testutils"
// 	"github.com/stretchr/testify/suite"
// )

// type ListSuite struct {
// 	suite.Suite
// 	method   string
// 	app      *fiber.App
// 	db       *pgxpool.Pool
// 	resolver *resolver.Resolver
// 	records  []*types.ExampleEntity
// }

// func TestListSuite(t *testing.T) {
// 	suite.Run(t, &ListSuite{})
// }

// // SetupSuite runs setup before all suite tests
// func (s *ListSuite) SetupSuite() {
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
// func (s *ListSuite) TearDownSuite() {
// 	//
// }

// // SetupTest runs setup before each test
// func (s *ListSuite) SetupTest() {
// 	records := make([]*types.ExampleEntity, 0, 4)

// 	for range records {
// 		record, err := insertRecord(s.db)
// 		if err != nil {
// 			s.T().Log(err)
// 		}
// 		records = append(records, record)
// 		s.T().Log("\n\nHERE\n\n")
// 	}
// }

// // TearDownTest runs teardown after each test
// func (s *ListSuite) TearDownTest() {
// 	utils.Cleanup(s.resolver)
// }

// func (s *ListSuite) TestResourceList() {
// 	tests := []utils.Setup{
// 		{
// 			Description: "resource list succeeds (200)",
// 			Route:       routePrefix,
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
