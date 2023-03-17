package resourcetest

import (
	"bytes"
	"testing"

	"github.com/BetterWorks/gosk-api/internal/resolver"
	utils "github.com/BetterWorks/gosk-api/test/testutils"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type CreateSuite struct {
	suite.Suite
	method   string
	app      *fiber.App
	db       *pgxpool.Pool
	resolver *resolver.Resolver
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, &CreateSuite{})
}

// SetupSuite runs setup before all suite tests
func (s *CreateSuite) SetupSuite() {
	s.T().Log("SetupSuite")

	app, db, resolver, err := utils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "POST"
	s.app = app
	s.db = db
	s.resolver = resolver
}

// TearDownSuite runs teardown after all suite tests
func (s *CreateSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

// SetupTest runs setup before each test
func (s *CreateSuite) SetupTest() {
	s.T().Log("SetupTest")
}

// TearDownTest runs teardown after each test
func (s *CreateSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}

func (s *CreateSuite) TestResourceCreate() {
	tests := []utils.Setup{
		{
			Description: "resource create succeeds (201) with valid payload",
			Route:       routePrefix,
			Request: utils.Request{
				Body: bytes.NewBuffer([]byte(`{
					"data": {
						"type": "resource",
						"properties": {
							"title": "Resource Title",
							"description": "Resource Description",
							"enabled": true,
							"status": 1
						}
					}
				}`)),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			Expected: utils.Expected{Code: 201},
		},
	}

	for _, test := range tests {
		req := utils.SetRequestData(s.method, test.Route, test.Request.Body, test.Request.Headers)
		msTimeout := 1000

		res, err := s.app.Test(req, msTimeout)
		if err != nil {
			s.T().Log(err)
		}

		s.Equalf(test.Expected.Code, res.StatusCode, test.Description)
	}
}
