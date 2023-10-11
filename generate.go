package main

//go:generate go run -mod=mod ./internal/ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate cd tools/fullschema && npm run generate
