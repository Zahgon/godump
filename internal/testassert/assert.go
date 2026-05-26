package testassert

import (
	"testing"
)

func message(msgAndArgs ...any) string { _ = "STUB: not implemented"; return "" }

func fail(t *testing.T, msg string, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func containsString(container string, item any) bool { _ = "STUB: not implemented"; return false }

// Contains fails the test if s does not contain contains.
func Contains(t *testing.T, s, contains any, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotContains fails the test if s contains contains.
func NotContains(t *testing.T, s, contains any, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Equal fails the test if expected and actual are not deeply equal.
func Equal(t *testing.T, expected, actual any, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// True fails the test if value is false.
//
//nolint:revive // bool flag parameter is intentional for assert-style API parity.
func True(t *testing.T, value bool, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// False fails the test if value is true.
//
//nolint:revive // bool flag parameter is intentional for assert-style API parity.
func False(t *testing.T, value bool, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func isNil(v any) bool { _ = "STUB: not implemented"; return false }

// Nil fails the test if v is not nil.
func Nil(t *testing.T, v any, msgAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// JSONEq fails the test if expected and actual are not equivalent JSON values.
func JSONEq(t *testing.T, expected, actual string, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NoError fails the test if err is non-nil.
func NoError(t *testing.T, err error, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}
