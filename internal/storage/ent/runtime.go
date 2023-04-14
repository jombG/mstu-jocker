// Code generated by ent, DO NOT EDIT.

package ent

import (
	"jocer/internal/storage/ent/jock"
	"jocer/internal/storage/schema"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	jockFields := schema.Jock{}.Fields()
	_ = jockFields
	// jockDescName is the schema descriptor for name field.
	jockDescName := jockFields[1].Descriptor()
	// jock.NameValidator is a validator for the "name" field. It is called by the builders before save.
	jock.NameValidator = jockDescName.Validators[0].(func(string) error)
	// jockDescID is the schema descriptor for id field.
	jockDescID := jockFields[0].Descriptor()
	// jock.DefaultID holds the default value on creation for the id field.
	jock.DefaultID = jockDescID.Default.(func() uuid.UUID)
}
