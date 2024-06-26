package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/jobcontractor"
	"gqlgen-ent/graph/model"
	"gqlgen-ent/middlewares"
)

// CreateContractor is the resolver for the createContractor field.
func (r *mutationResolver) CreateContractor(ctx context.Context, input model.JobContractorInput) (*ent.JobContractor, error) {
	panic(fmt.Errorf("not implemented: CreateContractor - createContractor"))
}

// UpdateContractor is the resolver for the updateContractor field.
func (r *mutationResolver) UpdateContractor(ctx context.Context, id string, input model.JobContractorInput) (*ent.JobContractor, error) {
	panic(fmt.Errorf("not implemented: UpdateContractor - updateContractor"))
}

// AllContractor is the resolver for the allContractor field.
func (r *queryResolver) AllContractor(ctx context.Context) ([]*ent.JobContractor, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.JobContractor.Query().Where(jobcontractor.DeletedEQ(0)).All(ctx)
}
