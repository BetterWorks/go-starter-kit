package episodetest

import (
	"bytes"
	"testing"

	utils "github.com/jasonsites/gosk-api/test/testutils"
	"github.com/stretchr/testify/assert"
)

func TestEpisodeCreate(t *testing.T) {
	var (
		routePrefix = "/domain/episodes"
		method      = "POST"
	)

	tests := []utils.Setup{
		{
			Description: "create episode succeeds (201) with valid payload",
			Route:       routePrefix,
			Request: utils.Request{
				Body: bytes.NewBuffer([]byte(`{
					"data": {
						"type": "episode",
						"properties": {
							"title": "Episode 1",
							"description": "Episode 1 Description...",
							"enabled": true,
							"season_id": "c544d628-1db5-4ce7-9f13-1a2109e96989",
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

	app, _, err := utils.InitializeApp(nil)
	if err != nil {
		t.Log(err)
	}

	for _, test := range tests {
		req := utils.SetRequestData(method, test.Route, test.Request.Body, test.Request.Headers)

		res, err := app.Test(req, 1000)
		if err != nil {
			t.Log(err)
		}

		assert.Equalf(t, test.Expected.Code, res.StatusCode, test.Description)
	}
}
