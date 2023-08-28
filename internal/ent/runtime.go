// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/ent/schema"
	"exusiai.dev/roguestats-backend/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[1].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() string)
	researchFields := schema.Research{}.Fields()
	_ = researchFields
	// researchDescName is the schema descriptor for name field.
	researchDescName := researchFields[1].Descriptor()
	// research.NameValidator is a validator for the "name" field. It is called by the builders before save.
	research.NameValidator = researchDescName.Validators[0].(func(string) error)
	// researchDescID is the schema descriptor for id field.
	researchDescID := researchFields[0].Descriptor()
	// research.DefaultID holds the default value on creation for the id field.
	research.DefaultID = researchDescID.Default.(func() string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescCredential is the schema descriptor for credential field.
	userDescCredential := userFields[3].Descriptor()
	// user.CredentialValidator is a validator for the "credential" field. It is called by the builders before save.
	user.CredentialValidator = userDescCredential.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
}