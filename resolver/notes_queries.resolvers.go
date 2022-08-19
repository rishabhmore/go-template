package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
)

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context, pagination *gqlmodels.NotesPagination) (*gqlmodels.NotesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
