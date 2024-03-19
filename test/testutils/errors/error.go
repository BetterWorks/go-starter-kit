package error

import "testing"

func NewLineErrorf(t *testing.T, expected any, actual any) {
	t.Errorf("\nExpected:\n'%+v'\nActual:\n'%+v'", expected, actual)
}
