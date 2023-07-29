package service

import (
	"context"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/appcontext"
)

type Directive struct {
	fx.In

	User User
}

// Private directive is used to check if the current user is the owner of the object.
// if not, return null for the field.
func (s Directive) Private(ctx context.Context, obj any, next graphql.Resolver, userIDFieldName *string) (res any, err error) {
	currentUser := appcontext.CurrentUser(ctx)
	fieldCtx := graphql.GetFieldContext(ctx)
	if currentUser == nil {
		// for anonymous user, return null as well
		return nil, nil
	}

	// if the field is not specified, use the default field name
	fieldName := "userId"
	if userIDFieldName != nil {
		fieldName = *userIDFieldName
	}

	// get the user ID from the object
	// you might need to use reflection to get the field value
	var found reflect.Value
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		thisFieldName := val.Type().Field(i).Tag.Get("json")
		if thisFieldName == fieldName {
			found = field
			break
		}
	}
	if !found.IsValid() {
		log.Warn().Msgf("field %s not found in %s", fieldName, fieldCtx.Field.Name)
		return nil, nil
	}

	// get the actual string
	var userID string
	if found.Kind() == reflect.Ptr {
		userID = found.Elem().String()
	} else {
		userID = found.String()
	}

	if userID == currentUser.ID {
		return next(ctx)
	}

	return nil, nil
}
