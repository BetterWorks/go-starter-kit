package controllers

import (
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/types"
)

// Config
type Config struct {
	Application *application.Application
	Logger      *types.Logger
}

// Controller
type Controller struct {
	application *application.Application
	logger      *types.Logger
}

// NewController
func NewController(c *Config) *Controller {
	return &Controller{
		application: c.Application,
		logger:      c.Logger,
	}
}
