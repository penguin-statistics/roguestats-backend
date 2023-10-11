package directives

import (
	"entgo.io/contrib/entgql"
	"github.com/vektah/gqlparser/v2/ast"
)

func Private(userIdFieldName string) entgql.Directive {
	return entgql.Directive{
		Name: "private",
		Arguments: []*ast.Argument{
			{
				Name: "userIdFieldName",
				Value: &ast.Value{
					Raw:  userIdFieldName,
					Kind: ast.StringValue,
				},
			},
		},
	}
}
