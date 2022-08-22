package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/internal/middleware/auth"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"go-template/pkg/utl/resultwrapper"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context, pagination *gqlmodels.NotesPagination) (*gqlmodels.NotesPayload, error) {
	// Let's grab the user id from the context
	userId := auth.UserIDFromContext(ctx)
	// Define a QueryMod slice with base query to get notes for a user id
	queryMods := []qm.QueryMod{
		qm.Where(fmt.Sprintf("%s=?", models.NoteColumns.UserID), userId),
	}
	// If pagination is enabled
	if pagination != nil {
		if pagination.Limit != 0 {
			// Add a new query specifying the page limit & the offset
			queryMods = append(queryMods, qm.Limit(pagination.Limit), qm.Offset(pagination.Page*pagination.Limit))
		}
	}

	notes, count, err := daos.GetNotesWithCount(queryMods, ctx)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	return &gqlmodels.NotesPayload{
		Total: int(count),
		Notes: convert.NotesToGraphqlNotes(notes),
	}, nil
}

// AllNotes is the resolver for the allNotes field.
func (r *queryResolver) AllNotes(ctx context.Context, pagination *gqlmodels.NotesPagination) (*gqlmodels.NotesPayload, error) {
	queryMods := []qm.QueryMod{}
	// If pagination is enabled
	if pagination != nil {
		if pagination.Limit != 0 {
			// Add a new query specifying the page limit & the offset
			queryMods = append(queryMods, qm.Limit(pagination.Limit), qm.Offset(pagination.Page*pagination.Limit))
		}
	}

	notes, count, err := daos.GetNotesWithCount(queryMods, ctx)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	return &gqlmodels.NotesPayload{
		Total: int(count),
		Notes: convert.NotesToGraphqlNotes(notes),
	}, nil
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
