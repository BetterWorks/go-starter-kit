package resolver

import (
	"context"
	"testing"
)

func TestConfigComponent(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	actual := r.Config()
	if actual != r.config {
		t.Errorf("resolver.config expected to be a singleton: actual: '%v'", actual)
	}
}

func TestDomainComponent(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	actual := r.Domain()
	if actual != r.domain {
		t.Errorf("resolver.domain expected to be a singleton: actual: '%v'", actual)
	}
}

func TestFlagsClientComponent(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	actual := r.FlagsClient()
	if actual != r.flagsClient {
		t.Errorf("resolver.flagsClient expected to be a singleton: actual: '%v'", actual)
	}
}

func TestFlagsClientComponentError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic during FlagsClient initialization, but there was none.")
		}
	}()

	r := NewResolver(context.Background(), nil)
	if r.Config().Flags.SDKKey == "" {
		t.Errorf("Flags.SDKKey should have been pipulated but it was empty")
	}
	r.config.Flags.SDKKey = ""
	r.flagsClient = nil
	if r.Config().Flags.SDKKey != "" {
		t.Errorf("Flags.SDKKey should have been empty but it was not")
	}
	r.FlagsClient()
}

func TestLogComponent(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	actual := r.Log()
	if actual != r.log {
		t.Errorf("resolver.log expected to be a singleton: actual: '%v'", actual)
	}
}

func TestMetadataComponent(t *testing.T) {
	r := NewResolver(context.Background(), nil)
	actual := r.Metadata()
	if actual != r.metadata {
		t.Errorf("resolver.metadata expected to be a singleton: actual: '%v'", actual)
	}
}
