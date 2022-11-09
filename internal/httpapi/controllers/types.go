package controllers

import "github.com/jasonsites/gosk-api/internal/core/types"

// Movie
type MovieRequestBody struct {
	Data MovieResource `json:"data"`
}

type MovieResource struct {
	Type       string      `json:"type"`
	ID         string      `json:"id"`
	Properties types.Movie `json:"properties"`
}

// Book
type BookRequestBody struct {
	Data BookResource `json:"data"`
}

type BookResource struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties types.Book `json:"properties"`
}
