package testrequire

import (
	"testing"
)

// True fails the test immediately if value is false.
//
//nolint:revive // bool flag parameter is intentional for require-style API parity.
func True(t *testing.T, value bool, msgAndArgs ...any) { _ = "STUB: not implemented"; return }

// NoError fails the test immediately if err is non-nil.
func NoError(t *testing.T, err error, msgAndArgs ...any) { _ = "STUB: not implemented"; return }

// Contains fails the test immediately if s does not contain contains.
func Contains(t *testing.T, s, contains any, msgAndArgs ...any) { _ = "STUB: not implemented"; return }
