// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         string                 `json:"id" bun:"user_id"`
	Name       string                 `json:"name"`
	Email      *string                `json:"email,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// User's encrypted credential
	Credential string `json:"-"`
}
