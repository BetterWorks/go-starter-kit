package jsonio

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/trace"
)

// DecodeRequest
func DecodeRequest(w http.ResponseWriter, r *http.Request, dest any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(1048576))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dest); err != nil {
		return err
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("request body must contain only one json object")
	}

	return nil
}

// EncodeResponse
func EncodeResponse(w http.ResponseWriter, r *http.Request, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		w.Write([]byte("internal server error"))
	}
}
