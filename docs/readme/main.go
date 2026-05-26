//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"regexp"
)

const (
	apiStart       = "<!-- api:embed:start -->"
	apiEnd         = "<!-- api:embed:end -->"
	testCountStart = "<!-- test-count:embed:start -->"
	testCountEnd   = "<!-- test-count:embed:end -->"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("✔ API section updated in README.md")
}

func run() error { _ = "STUB: not implemented"; return nil }

//
// ------------------------------------------------------------
// Data model
// ------------------------------------------------------------
//

type FuncDoc struct {
	Name        string
	Group       string
	Behavior    string
	Fluent      string
	Description string
	Examples    []Example
}

type Example struct {
	Label string
	Code  string
	Line  int
}

//
// ------------------------------------------------------------
// Parsing
// ------------------------------------------------------------
//

var (
	groupHeader    = regexp.MustCompile(`(?i)^\s*@group\s+(.+)$`)
	behaviorHeader = regexp.MustCompile(`(?i)^\s*@behavior\s+(.+)$`)
	fluentHeader   = regexp.MustCompile(`(?i)^\s*@fluent\s+(.+)$`)
	exampleHeader  = regexp.MustCompile(`(?i)^\s*Example:\s*(.*)$`)
)

func parseFuncs(root string) ([]*FuncDoc, error) { _ = "STUB: not implemented"; return nil, nil }

func extractGroup(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func extractBehavior(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func extractFluent(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func extractDescription(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func extractExamples(fset *token.FileSet, fn *ast.FuncDecl) []Example {
	_ = "STUB: not implemented"
	return nil
}

// selectPackage picks the primary package to document.
// Strategy:
//  1. If only one package exists, use it.
//  2. Prefer the non-"main" package with the most files.
//  3. Fall back to the first package alphabetically.
func selectPackage(pkgs map[string]*ast.Package) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//
// ------------------------------------------------------------
// Rendering
// ------------------------------------------------------------
//

func renderAPI(funcs []*FuncDoc) string { _ = "STUB: not implemented"; return "" }

// ---------------- Index ----------------

// ---------------- Details ----------------

//
// ------------------------------------------------------------
// README replacement
// ------------------------------------------------------------
//

func replaceAPISection(readme, api string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func countTests(root string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var testsBadgePattern = regexp.MustCompile(`tests-\d+-brightgreen`)

func updateTestsSection(readme string, tests int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//
// ------------------------------------------------------------
// Helpers
// ------------------------------------------------------------
//

func findRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func fileExists(p string) bool { _ = "STUB: not implemented"; return false }

func normalizeIndent(lines []string) []string { _ = "STUB: not implemented"; return nil }
