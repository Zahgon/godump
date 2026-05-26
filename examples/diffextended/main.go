//go:build ignore
// +build ignore

package main

import (
	"time"

	"github.com/goforj/godump"
)

type FriendlyDuration time.Duration

// String renders the duration as HH:MM:SS.
func (fd FriendlyDuration) String() string { _ = "STUB: not implemented"; return "" }

type IsZeroer interface {
	IsZero() bool
}

type Inner struct {
	ID    int
	Notes []string
	Blob  []byte
}

type Ref struct {
	Self *Ref
}

type Everything struct {
	String        string
	Bool          bool
	Int           int
	Float         float64
	Time          time.Time
	Duration      time.Duration
	Friendly      FriendlyDuration
	PtrString     *string
	PtrDuration   *time.Duration
	SliceInts     []int
	ArrayStrings  [2]string
	MapValues     map[string]int
	Nested        Inner
	NestedPtr     *Inner
	Interface     any
	InterfaceImpl IsZeroer
	Recursive     *Ref
	privateField  string
	privateStruct Inner
}

// makeEverything builds a populated sample struct.
func makeEverything(now time.Time, label string) Everything {
	_ = "STUB: not implemented"
	return *new(Everything)
}

// main demonstrates diffing two complex structures.
func main() {
	now := time.Date(2025, 12, 18, 16, 34, 37, 0, time.FixedZone("CST", -6*60*60))

	before := makeEverything(now, "v1")
	after := makeEverything(now.Add(time.Hour), "v2")
	after.Int = 99
	after.Duration = time.Minute * 45
	after.Friendly = FriendlyDuration(after.Duration)
	after.SliceInts = []int{1, 3, 4}
	after.MapValues["c"] = 3
	after.Nested.Notes = []string{"alpha", "gamma"}
	after.NestedPtr = nil
	after.Interface = map[string]bool{"ok": false}
	after.privateField = "changed"

	godump.Diff(before, after)

	diff := godump.DiffStr(before, after)
	_ = diff
}
