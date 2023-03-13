package resourcetest

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasonsites/gosk-api/internal/resolver"
	"github.com/jasonsites/gosk-api/internal/types"
	utils "github.com/jasonsites/gosk-api/test/testutils"
	"github.com/stretchr/testify/suite"
)

type UpdateSuite struct {
	suite.Suite
	method   string
	app      *fiber.App
	db       *pgxpool.Pool
	resolver *resolver.Resolver
	record   *types.ResourceEntity
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, &UpdateSuite{})
}

// SetupSuite runs setup before all suite tests
func (s *UpdateSuite) SetupSuite() {
	s.T().Log("SetupSuite")

	app, db, resolver, err := utils.InitializeApp(nil)
	if err != nil {
		s.T().Log(err)
	}

	s.method = "PUT"
	s.app = app
	s.db = db
	s.resolver = resolver
}

// TearDownSuite runs teardown after all suite tests
func (s *UpdateSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

// SetupTest runs setup before each test
func (s *UpdateSuite) SetupTest() {
	record, err := insertRecord(s.db)
	if err != nil {
		s.T().Log(err)
	}
	s.record = record
}

// TearDownTest runs teardown after each test
func (s *UpdateSuite) TearDownTest() {
	utils.Cleanup(s.resolver)
}

func (s *UpdateSuite) TestResourceUpdate() {
	tests := []utils.Setup{
		{
			Description: "resource update succeeds (200) with valid payload",
			Route:       fmt.Sprintf("%s/%s", routePrefix, s.record.ID.String()),
			Request: utils.Request{
				Body: bytes.NewBuffer([]byte(fmt.Sprintf(`{
					"data": {
						"type": "resource",
						"id": "%s",
						"properties": {
							"title": "Resource Title",
							"description": "Updated resource description",
							"enabled": true,
							"status": 1
						}
					}
				}`, s.record.ID.String()))),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			Expected: utils.Expected{Code: 200},
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
