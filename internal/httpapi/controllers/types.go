package controllers

import "github.com/jasonsites/gosk-api/internal/types"

// MovieResource
type MovieRequestResource struct {
	Type       string      `json:"type"`
	ID         string      `json:"id"`
	Properties types.Movie `json:"properties"`
}

// MovieRequestBody
type MovieRequestBody struct {
	Data MovieRequestResource `json:"data"`
}

// BookResource
type BookRequestResource struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Properties types.Book `json:"properties"`
}

// BookRequestBody
type BookRequestBody struct {
	Data BookRequestResource `json:"data"`
}
