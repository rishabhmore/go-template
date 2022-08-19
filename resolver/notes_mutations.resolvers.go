package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input gqlmodels.NotesCreateInput) (*gqlmodels.Notes, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, input *gqlmodels.NotesUpdateInput) (*gqlmodels.Notes, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context) (*gqlmodels.NotesDelete, error) {
	panic(fmt.Errorf("not implemented"))
}
