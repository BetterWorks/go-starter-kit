package httpserver

import (
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
)

var Error_Validator_NilServerConfig = "validator: (nil *httpserver.ServerConfig)"

func TestNewServer_nilConfig(t *testing.T) {
	t.Parallel()

	_, err := NewServer(nil)

	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}

	if err.Error() != Error_Validator_NilServerConfig {
		te.NewLineErrorf(t, Error_Validator_NilServerConfig, err.Error())
	}
}
