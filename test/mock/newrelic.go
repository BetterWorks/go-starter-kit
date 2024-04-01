package mock

import (
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicClient(t *testing.T) (*newrelic.Application, error) {
	client, err := newrelic.NewApplication(
		newrelic.ConfigAppName("test"),
		newrelic.ConfigLicense(fake.DigitN(40)),
	)
	if err != nil {
		t.Fatalf("mock.NewRelicClient instantiation error")
	}

	return client, nil
}
