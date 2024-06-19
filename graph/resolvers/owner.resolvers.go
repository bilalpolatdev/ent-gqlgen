package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/jobowner"
	"gqlgen-ent/graph/model"
	"gqlgen-ent/middlewares"
)

// CreateOwner is the resolver for the createOwner field.
func (r *mutationResolver) CreateOwner(ctx context.Context, input model.JobOwnerInput) (*ent.JobOwner, error) {
	panic(fmt.Errorf("not implemented: CreateOwner - createOwner"))
}

// UpdateOwner is the resolver for the updateOwner field.
func (r *mutationResolver) UpdateOwner(ctx context.Context, id string, input model.JobOwnerInput) (*ent.JobOwner, error) {
	panic(fmt.Errorf("not implemented: UpdateOwner - updateOwner"))
}

// AllOwner is the resolver for the allOwner field.
func (r *queryResolver) AllOwner(ctx context.Context) ([]*ent.JobOwner, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.JobOwner.Query().Where(jobowner.DeletedEQ(0)).All(ctx)
}