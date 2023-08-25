package entid

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func New(entity string) string {
	return entity + "_" + strings.ToLower(ulid.Make().String())
}

func generator(entity string) func() string {
	return func() string {
		return New(entity)
	}
}

var (
	User  = generator("usr")
	Event = generator("evt")
)
