package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"
)

// PostMessageTo is the resolver for the postMessageTo field.
func (r *mutationResolver) PostMessageTo(ctx context.Context, subscriber string, content string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Subscribe is the resolver for the subscribe field.
func (r *subscriptionResolver) Subscribe(ctx context.Context, subscriber string) (<-chan string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type (
	mutationResolver     struct{ *Resolver }
	subscriptionResolver struct{ *Resolver }
)
