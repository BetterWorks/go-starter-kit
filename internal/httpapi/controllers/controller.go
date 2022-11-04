package controllers

import (
	"github.com/jasonsites/gosk-api/internal/core/types"
)

type Config struct {
	Application types.Application
	Logger      *types.Logger
}

type Controller struct {
	application types.Application
	logger      *types.Logger
}

func NewController(c *Config) *Controller {
	return &Controller{
		application: c.Application,
		logger:      c.Logger,
	}
}
