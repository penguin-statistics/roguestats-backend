package controller

import (
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/graph"
	"exusiai.dev/roguestats-backend/internal/middleware"
	"exusiai.dev/roguestats-backend/internal/service"
)

type GraphQL struct {
	fx.In

	ResolverDeps     graph.ResolverDeps
	Middleware       middleware.Middleware
	DirectiveService service.Directive
	Ent              *ent.Client
	Route            fiber.Router `name:"root"`
}

func RegisterGraphQL(c GraphQL) {
	graphConfig := graph.Config{
		Resolvers: &graph.Resolver{
			ResolverDeps: c.ResolverDeps,
		},
	}
	graphConfig.Directives.Admin = c.DirectiveService.Admin
	graphConfig.Directives.Authenticated = c.DirectiveService.Authenticated
	graphConfig.Directives.Private = c.DirectiveService.Private

	srv := handler.New(graph.NewExecutableSchema(graphConfig))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(200))
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.Use(entgql.Transactioner{TxOpener: c.Ent})

	c.Route.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/graphql")))
	c.Route.Post("/graphql", c.Middleware.InjectFiberCtx(), c.Middleware.CurrentUser(), adaptor.HTTPHandler(srv))
}
