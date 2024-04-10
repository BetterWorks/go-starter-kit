package http

import (
	"errors"
	"io"
	"net/http"
)

// Expected
type Expected struct {
	ResponseBody string
	StatusCode   int
}

// RequestData
type RequestData struct {
	Body    io.Reader
	Headers map[string]string
	Method  string
	URL     string
}

// RequestOptions
type RequestOptions struct {
	JSON bool
}

// SetRequestData creates a new HTTP Request instance from the given data
func (r *RequestData) SetRequestData(opts *RequestOptions) (*http.Request, error) {
	if r == nil {
		return nil, errors.New("request data must be non-nil")
	}
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		return nil, err
	}
	req = r.SetRequestHeaders(req, r.Headers, opts)
	return req, nil
}

// SetRequestHeaders set all headers on the given request
func (r *RequestData) SetRequestHeaders(req *http.Request, headers map[string]string, opts *RequestOptions) *http.Request {
	if opts == nil || opts.JSON {
		req.Header.Add("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}
