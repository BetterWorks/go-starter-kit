package domain

import (
	"testing"

	"github.com/BetterWorks/go-starter-kit/test/mock"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
)

func TestNewDomain(t *testing.T) {
	t.Parallel()

	services := &Services{
		Example: &mock.ExampleService{},
	}
	_, err := NewDomain(services)
	if err != nil {
		te.NewLineErrorf(t, nil, err.Error())
	}

	services = &Services{}
	_, err = NewDomain(services)
	if err == nil {
		te.NewLineErrorf(t, "Error", nil)
	}
}
