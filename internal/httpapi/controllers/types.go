package controllers

import "github.com/jasonsites/gosk-api/internal/core/types"

// MovieRequestBody
type MovieRequestBody struct {
	Data MovieResource `json:"data"`
}

// MovieResource
type MovieResource struct {
	Type       string      `json:"type"`
	ID         string      `json:"id"`
	Properties types.Movie `json:"properties"`
}

// BookRequestBody
type BookRequestBody struct {
	Data BookResource `json:"data"`
}

// BookResource
type BookResource struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties types.Book `json:"properties"`
}
