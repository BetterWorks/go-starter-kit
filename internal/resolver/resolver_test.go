package resolver

import (
	"context"
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
)

func TestDefaultNewResolver(t *testing.T) {
	r := NewResolver(context.Background(), nil)

	if r.appContext == nil {
		t.Errorf("r.appContext expected to be non-nil: actual: '%+v'", r.appContext)
	}

	if r.config != nil {
		t.Errorf("resolver.config expected to be nil: actual: '%+v'", r.config)
	}
	if r.domain != nil {
		t.Errorf("resolver.domain expected to be nil: actual: '%+v'", r.domain)
	}
	if r.log != nil {
		t.Errorf("resolver.log expected to be nil: actual: '%+v'", r.log)
	}
	if r.metadata != nil {
		t.Errorf("resolver.metadata expected to be nil: actual: '%+v'", r.metadata)
	}
}

func TestLoadHTTPEntry(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	r.Load("http")

	if r.httpServer == nil {
		t.Errorf("resolver.httpServer expected to be non-nil: actual: '%+v'", r.httpServer)
	}
}

func TestLoadInvalidEntry(t *testing.T) {
	r := NewResolver(context.Background(), nil)

	defer func() {
		if r := recover(); r == nil {
			te.NewLineErrorf(t, "panic on invalid entry", nil)
		}
	}()

	r.Load("invalid")
}
