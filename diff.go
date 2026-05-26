package godump

import (
	"io"
)

// Diff prints a diff between two values to stdout.
// @group Diff
//
// Example: print diff
//
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	godump.Diff(a, b)
//	// <#diff // path:line
//	// - #map[string]int {
//	// -   a => 1 #int
//	// - }
//	// + #map[string]int {
//	// +   a => 2 #int
//	// + }
func Diff(a, b any) { _ = "STUB: not implemented"; return }

// Diff prints a diff between two values to the configured writer.
// @group Diff
//
// Example: print diff with a custom dumper
//
//	d := godump.NewDumper()
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	d.Diff(a, b)
//	// <#diff // path:line
//	// - #map[string]int {
//	// -   a => 1 #int
//	// - }
//	// + #map[string]int {
//	// +   a => 2 #int
//	// + }
func (d *Dumper) Diff(a, b any) { _ = "STUB: not implemented"; return }

// DiffStr returns a string diff between two values.
// @group Diff
//
// Example: diff string
//
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	out := godump.DiffStr(a, b)
//	_ = out
//	// <#diff // path:line
//	// - #map[string]int {
//	// -   a => 1 #int
//	// - }
//	// + #map[string]int {
//	// +   a => 2 #int
//	// + }
func DiffStr(a, b any) string { _ = "STUB: not implemented"; return "" }

// DiffStr returns a string diff between two values.
// @group Diff
//
// Example: diff string with a custom dumper
//
//	d := godump.NewDumper()
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	out := d.DiffStr(a, b)
//	_ = out
//	// <#diff // path:line
//	// - #map[string]int {
//	// -   a => 1 #int
//	// - }
//	// + #map[string]int {
//	// +   a => 2 #int
//	// + }
func (d *Dumper) DiffStr(a, b any) string { _ = "STUB: not implemented"; return "" }

// DiffHTML returns an HTML diff between two values.
// @group Diff
//
// Example: HTML diff
//
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	html := godump.DiffHTML(a, b)
//	_ = html
//	// (html diff)
func DiffHTML(a, b any) string { _ = "STUB: not implemented"; return "" }

// DiffHTML returns an HTML diff between two values.
// @group Diff
//
// Example: HTML diff with a custom dumper
//
//	d := godump.NewDumper()
//	a := map[string]int{"a": 1}
//	b := map[string]int{"a": 2}
//	html := d.DiffHTML(a, b)
//	_ = html
//	// (html diff)
func (d *Dumper) DiffHTML(a, b any) string { _ = "STUB: not implemented"; return "" }

type diffDumpPair struct {
	left  string
	right string
}

// diffDumps builds the left and right dump strings, aligning reference ids.
func (d *Dumper) diffDumps(a, b any) diffDumpPair {
	_ = "STUB: not implemented"
	return *new(diffDumpPair)
}

// dumpStrNoHeader renders a dump without the header line.
func (d *Dumper) dumpStrNoHeader(vs ...any) string { _ = "STUB: not implemented"; return "" }

// printDiffHeader writes the diff header line when a caller frame is available.
func (d *Dumper) printDiffHeader(out io.Writer) { _ = "STUB: not implemented"; return }

// typeStringForAny returns a displayable type for a value.
func (d *Dumper) typeStringForAny(v any) string { _ = "STUB: not implemented"; return "" }

type diffKind int

const (
	diffEqual diffKind = iota
	diffDelete
	diffInsert
)

type diffLine struct {
	kind diffKind
	text string
}

// diffLines computes a line-level diff with insert/delete operations.
func diffLines(a, b []string) []diffLine { _ = "STUB: not implemented"; return nil }

// diffPrefix returns the colored diff marker prefix.
func (d *Dumper) diffPrefix(kind diffKind) string { _ = "STUB: not implemented"; return "" }

// diffTintLine tints a full diff line based on change type.
func (d *Dumper) diffTintLine(line string, kind diffKind) string {
	_ = "STUB: not implemented"
	return ""
}

// tintBackgroundLine applies a full-line background while preserving text colors.
func (d *Dumper) tintBackgroundLine(line, bgCode, bgHex string) string {
	_ = "STUB: not implemented"
	return ""
}

const ansiEscape = '\x1b'
const ansiEraseLine = "\x1b[K"

// stripANSI removes ANSI escape sequences from a string.
func stripANSI(s string) string { _ = "STUB: not implemented"; return "" }

// Skip ANSI CSI sequences like "\x1b[31m" or "\x1b[K".

// consume final byte

// Drop lone escape byte.

// stripHTMLSpans removes color span tags while preserving content.
func stripHTMLSpans(s string) string { _ = "STUB: not implemented"; return "" }

// isHTMLLine reports whether the line contains HTML color spans.
func isHTMLLine(line string) bool { _ = "STUB: not implemented"; return false }

// splitLines splits a string into lines while normalizing CRLF and trimming a trailing newline.
func splitLines(s string) []string { _ = "STUB: not implemented"; return nil }
