package entid

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func New(entity string) string {
	return entity + "_" + strings.ToLower(ulid.Make().String())
}

func NewGenerator(entity string) func() string {
	return func() string {
		return New(entity)
	}
}

var (
	User  = NewGenerator("usr")
	Event = NewGenerator("evt")
)
