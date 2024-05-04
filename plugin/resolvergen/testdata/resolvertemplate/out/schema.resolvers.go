package out

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44-dev

import (
	"context"
	"fmt"

	customresolver "github.com/99designs/gqlgen/plugin/resolvergen/testdata/singlefile/out"
)

// Resolver is the resolver for the resolver field.
func (r *queryCustomResolverType) Resolver(ctx context.Context) (*customresolver.Resolver, error) {
	// Custom Resolver implementation
	panic(fmt.Errorf("custom Resolver not implemented: Resolver - resolver"))
}

// Name is the resolver for the name field.
func (r *resolverCustomResolverType) Name(ctx context.Context, obj *customresolver.Resolver) (string, error) {
	// Custom Resolver implementation
	panic(fmt.Errorf("custom Resolver not implemented: Name - name"))
}

// Query returns customresolver.QueryResolver implementation.
func (r *CustomResolverType) Query() customresolver.QueryResolver { return &queryCustomResolverType{r} }

// Resolver returns customresolver.ResolverResolver implementation.
func (r *CustomResolverType) Resolver() customresolver.ResolverResolver {
	return &resolverCustomResolverType{r}
}

type (
	queryCustomResolverType    struct{ *CustomResolverType }
	resolverCustomResolverType struct{ *CustomResolverType }
)
