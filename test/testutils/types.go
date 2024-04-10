package testutils

import "io"

type Expected struct {
	Code int
}

type Request struct {
	Body    io.Reader
	Headers map[string]string
}

type Setup struct {
	Description string
	Method      string
	Route       string
	Request     Request
	Expected    Expected
}
