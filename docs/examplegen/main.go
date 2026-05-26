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

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("✔ Examples generated in ./examples/")
}

func run() error { _ = "STUB: not implemented"; return nil }

// Debug / inspection hook (optional)
//env.Dump(fd)

func findRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func fileExists(p string) bool { _ = "STUB: not implemented"; return false }

func modulePath(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }

//
// ------------------------------------------------------------
// Data models
// ------------------------------------------------------------
//

type FuncDoc struct {
	Name        string
	Group       string
	Description string
	Examples    []Example
}

type Example struct {
	FuncName string
	File     string
	Label    string
	Line     int
	Code     string
}

//
// ------------------------------------------------------------
// Example extraction
// ------------------------------------------------------------
//

var exampleHeader = regexp.MustCompile(`(?i)^\s*Example:\s*(.*)$`)
var groupHeader = regexp.MustCompile(`(?i)^\s*@group\s+(.+)$`)

type docLine struct {
	text string
	pos  token.Pos
}

func extractFuncDocs(
	fset *token.FileSet,
	filename string,
	file *ast.File,
) map[string]*FuncDoc {
	_ = "STUB: not implemented"
	return nil
}

func extractGroup(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func extractFuncDescription(group *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

// Stop before Example or @group

func docLines(group *ast.CommentGroup) []docLine { _ = "STUB: not implemented"; return nil }

func extractBlocks(
	fset *token.FileSet,
	filename, funcName string,
	fn *ast.FuncDecl,
) []Example {
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
// Write ./examples/<func>/main.go
// ------------------------------------------------------------
//

func writeMain(base string, fd *FuncDoc, importPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Build tag

// Description

// Examples

func containsCodeUsage(code, token string) bool { _ = "STUB: not implemented"; return false }
