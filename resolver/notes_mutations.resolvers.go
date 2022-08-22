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
	"go-template/pkg/utl/throttle"
	"strconv"
	"time"

	null "github.com/volatiletech/null/v8"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input gqlmodels.NotesCreateInput) (*gqlmodels.Notes, error) {
	// ensure we cannot create new notes within less than 3 seconds
	err := throttle.Check(ctx, 1, 3*time.Second)
	if err != nil {
		return nil, err
	}
	// get the current user from context
	user := auth.FromContext(ctx)
	// create the model for our sql boiler transaction
	modelNote := models.Note{
		UserID:    null.IntFrom(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Title:     null.StringFrom(input.Title),
		Note:      null.StringFromPtr(input.Note),
	}
	// use our dao to insert a note
	note, err := daos.CreateNote(modelNote, ctx)
	// validate errors
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "create note")
	}
	// convert our modelnoteto graphql note
	graphNote := convert.NoteToGraphqlNote(&note)
	// return our saved note
	return graphNote, err
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, input *gqlmodels.NotesUpdateInput) (*gqlmodels.Notes, error) {
	// ensure we cannot updates notes within less than 5 seconds
	err := throttle.Check(ctx, 3, 5*time.Second)
	if err != nil {
		return nil, err
	}
	// First let's validate if the note user is trying to update exists or not
	noteId, convErr := strconv.Atoi(input.ID)
	if convErr != nil {
		// If our ID string is invalid, we throw error
		return nil, resultwrapper.ResolverWrapperFromMessage(400, "Not a valid id")
	}
	// query our note & see if it exists or not
	_, queryErr := daos.FindNoteById(noteId, ctx)
	if queryErr != nil {
		return nil, resultwrapper.ResolverSQLError(err, "update note")
	}
	// get the current user from context
	user := auth.FromContext(ctx)
	// create the model for our sql boiler transaction
	modelNote := models.Note{
		ID:        noteId,
		UserID:    null.IntFrom(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Title:     null.StringFromPtr(input.Title),
		Note:      null.StringFromPtr(input.Note),
	}
	// use our dao to update a note
	note, err := daos.UpdateNote(modelNote, ctx)
	// validate errors
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "update note")
	}
	// convert our modelnoteto graphql note
	graphNote := convert.NoteToGraphqlNote(&note)
	// return our updated note
	return graphNote, err
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (*gqlmodels.NotesDelete, error) {
	// Lets get our id from string
	noteId, convErr := strconv.Atoi(id)
	if convErr != nil {
		//if the id passed was not a valid string integer, we throw error
		return nil, resultwrapper.ResolverWrapperFromMessage(400, "Not a valid id")
	}
	// We get our note from the note's id
	modelNote, queryErr := daos.FindNoteById(noteId, ctx)
	// If we didn't find a note, we throw error
	if queryErr != nil {
		return nil, resultwrapper.ResolverSQLError(queryErr, "note")
	}
	// Now that we found our note, we will try to delete it
	_, err := daos.DeleteNote(*modelNote, ctx)
	if err != nil {
		// if some failure happens to delete the note, we throw error
		return nil, resultwrapper.ResolverSQLError(err, "note")
	}
	return &gqlmodels.NotesDelete{ID: fmt.Sprint(noteId)}, nil
}
