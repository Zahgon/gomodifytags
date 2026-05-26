package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"

	"github.com/fatih/gomodifytags/modifytags"
)

// structType contains a structType node and it's name. It's a convenient
// helper type, because *ast.StructType doesn't contain the name of the struct.
// Also, we want to be able to handle structs selected by their variable name,
// so we can't simply use a TypeSpec.
type structType struct {
	name string
	node *ast.StructType
}

// output is used usually by editors
type output struct {
	Start  int      `json:"start"`
	End    int      `json:"end"`
	Lines  []string `json:"lines"`
	Errors []string `json:"errors,omitempty"`
}

// A config defines how the input is parsed and formatted.
type config struct {
	file     string
	output   string
	quiet    bool
	write    bool
	modified io.Reader

	offset     int
	structName string
	fieldName  string
	line       string
	start, end int
	all        bool

	fset *token.FileSet
}

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func realMain() error { _ = "STUB: not implemented"; return nil }

func parseFlags(args []string) (*config, *modifytags.Modification, error) {
	_ = "STUB: not implemented"

	// file flags
	return nil, nil, nil
}

// processing modes

// tag flags

// formatting

// option flags

// this fails if there are flags re-defined with the same name.

func run(cfg *config, mod *modifytags.Modification) error { _ = "STUB: not implemented"; return nil }

func (cfg *config) parse() (*ast.File, error) { _ = "STUB: not implemented"; return nil, nil }

func parseTransform(input string) (modifytags.Transform, error) {
	_ = "STUB: not implemented"
	return *new(modifytags.Transform), nil
}

func parseOptions(options string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findSelection returns the start and end positions of the fields that are
// subject to change. It depends on the line, struct or offset selection.
func (cfg *config) findSelection(file *ast.File) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

// collectStructs collects and maps structType nodes to their positions
func collectStructs(node ast.Node) map[token.Pos]*structType { _ = "STUB: not implemented"; return nil }

// this case also catches struct fields and the structName
// therefore might contain the field name (which is wrong)
// because `x.Type` in this case is not a *ast.StructType.
//
// We're OK with it, because, in our case *ast.Field represents
// a parameter declaration, i.e:
//
//   func test(arg struct {
//   	Field int
//   }) {
//   }
//
// and hence the struct name will be `arg`.

// if expression is in form "*T" or "[]T", dereference to check if "T"
// contains a struct expression

func (cfg *config) format(file ast.Node, rwErr *modifytags.RewriteErrors) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NOTE(arslan): print first the whole file and then cut out our
// selection. The reason we don't directly print the struct is that the
// printer is not capable of printing loosy comments, comments that are
// not part of any field inside a struct. Those are part of *ast.File
// and only printed inside a struct if we print the whole file. This
// approach is the sanest and simplest way to get a struct printed
// back. Second, our cursor might intersect two different structs with
// other declarations in between them. Printing the file and cutting
// the selection is the easier and simpler to do.

// this is the default config from `format.Node()`, but we add
// `printer.SourcePos` to get the original source position of the
// modified lines

// prevent selection to be larger than the actual number of lines

func (cfg *config) lineSelection(file *ast.File) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

// Convert start and end line numbers to token.Pos.

// Get the position of the end of the line

func (cfg *config) structSelection(file *ast.File) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

// if field name has been specified as well, only select the given field

func (cfg *config) fieldSelection(st *ast.StructType) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

func (cfg *config) offsetSelection(file *ast.File) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

// offset selects all fields

// allSelection selects all structs inside a file
func (cfg *config) allSelection(file *ast.File) (token.Pos, token.Pos, error) {
	_ = "STUB: not implemented"
	return *new(token.Pos), *new(token.Pos), nil
}

// validate determines whether the config is valid or not
func (cfg *config) validate() error { _ = "STUB: not implemented"; return nil }

// parseLines parses the given buffer and returns a slice of lines
func parseLines(buf io.Reader) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// check for any line directive and store it for next iteration to
// re-construct the original file. If it's not a line directive,
// continue consturcting the original file

// split splits the given line directive and returns the line number
// see https://golang.org/cmd/compile/#hdr-Compiler_Directives for more
// information
// NOTE(arslan): this only splits the line directive that the go.Parser
// outputs. If the go parser changes the format of the line directive, make
// sure to fix it in the below function
func split(line string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// deref takes an expression, and removes all its leading "*" and "[]"
// operator. Uuse case : if found expression is a "*t" or "[]t", we need to
// check if "t" contains a struct expression.
func deref(x ast.Expr) ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }
