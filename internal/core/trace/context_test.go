package trace

import (
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

func Test_ContextKey_String(t *testing.T) {
	t.Parallel()

	s := fake.Word()
	key := ContextKey(s)

	result := key.String()

	if result != s {
		te.NewLineErrorf(t, s, result)
	}
}
