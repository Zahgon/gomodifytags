package modifytags

import (
	"go/ast"
	"go/token"

	"github.com/fatih/structtag"
)

// A Transform determines how Go field names will be translated into
// names used in struct tags. For example, the [SnakeCase] transform
// converts the field MyField into the json tag json:"my_field".
type Transform int

const (
	SnakeCase  = iota // MyField -> my_field
	CamelCase         // MyField -> myField
	LispCase          // MyField -> my-field
	PascalCase        // MyField -> MyField
	TitleCase         // MyField -> My Field
	Keep              // keep the existing field name
)

// A Modification defines how struct tags should be modified for a given input struct.
type Modification struct {
	Add        []string            // tags to add
	AddOptions map[string][]string // options to add, per tag

	Remove        []string            // tags to remove
	RemoveOptions map[string][]string // options to remove, per tag

	Overwrite            bool // if set, replace existing tags when adding
	SkipUnexportedFields bool // if set, do not modify tags on unexported struct fields

	Transform    Transform // transform rule for adding tags
	Sort         bool      // if set, sort tags in ascending order by key name
	ValueFormat  string    // format for the tag's value, after transformation; for example "column:{field}"
	Clear        bool      // if set, clear all tags. tags are cleared before any new tags are added
	ClearOptions bool      // if set, clear all tag options; options are cleared before any new options are added
}

// Apply applies the struct tag modifications of the receiver to all
// struct fields contained within the given node between start and end position, modifying its input.
func (mod *Modification) Apply(fset *token.FileSet, node ast.Node, start, end token.Pos) error {
	_ = "STUB: not implemented"
	return nil
}

// rewrite rewrites the node for structs between the start and end positions
func (mod *Modification) rewrite(fset *token.FileSet, node ast.Node, start, end token.Pos) error {
	_ = "STUB: not implemented"
	return nil
}

// not in range

// anonymous field

// nothing to process, continue with next field

// processField returns the new struct tag value for the given field
func (mod *Modification) processField(fieldName, tagVal string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (mod *Modification) removeTags(tags *structtag.Tags) *structtag.Tags {
	_ = "STUB: not implemented"
	return nil
}

func (mod *Modification) clearTags(tags *structtag.Tags) *structtag.Tags {
	_ = "STUB: not implemented"
	return nil
}

func (mod *Modification) clearOptions(tags *structtag.Tags) *structtag.Tags {
	_ = "STUB: not implemented"
	return nil
}

func (mod *Modification) removeTagOptions(tags *structtag.Tags) (*structtag.Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mod *Modification) addTagOptions(tags *structtag.Tags) (*structtag.Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mod *Modification) addTags(fieldName string, tags *structtag.Tags) (*structtag.Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use snakecase as the default.

// support old style for backward compatibility

// tag doesn't exist, create a new one

func isPublicName(name string) bool { _ = "STUB: not implemented"; return false }

// validate determines whether the Modification is valid or not.
func (mod *Modification) validate() error { _ = "STUB: not implemented"; return nil }

func quote(tag string) string { _ = "STUB: not implemented"; return "" }

// RewriteErrors are errors that occurred while rewriting struct field tags.
type RewriteErrors struct {
	Errs []error
}

func (r *RewriteErrors) Error() string { _ = "STUB: not implemented"; return "" }
