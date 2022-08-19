package daos

import (
	"context"
	"database/sql"

	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// FindNoteById
func FindNoteById(noteId int, ctx context.Context) (*models.Note, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindNote(ctx, contextExecutor, noteId)
}

// Get all the notes with a count
func GetNotesWithCount(queryMods []qm.QueryMod, ctx context.Context) (models.NoteSlice, int64, error) {
	contextExecutor := getContextExecutor(nil)
	notes, err := models.Notes(queryMods...).All(ctx, contextExecutor)
	if err != nil {
		return models.NoteSlice{}, 0, err
	}
	queryMods = append(queryMods, qm.Offset(0))
	count, err := models.Notes(queryMods...).Count(ctx, contextExecutor)
	return notes, count, err
}

// Create notes transaction
func CreateNoteTx(note models.Note, ctx context.Context, tx *sql.Tx) (models.Note, error) {
	contextExecutor := getContextExecutor(tx)

	err := note.Insert(ctx, contextExecutor, boil.Infer())
	return note, err
}

// Create note
func CreateNote(note models.Note, ctx context.Context, tx *sql.Tx) (models.Note, error) {
	return CreateNoteTx(note, ctx, tx)
}

// Update notes transaction
func UpdateNoteTx(note models.Note, ctx context.Context, tx *sql.Tx) (models.Note, error) {
	contextExecutor := getContextExecutor(tx)

	_, err := note.Update(ctx, contextExecutor, boil.Infer())
	return note, err
}

// Update note
func UpdateNote(note models.Note, ctx context.Context, tx *sql.Tx) (models.Note, error) {
	return UpdateNoteTx(note, ctx, tx)
}

// Delete note
func DeleteNote(note models.Note, ctx context.Context) (int64, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := note.Delete(ctx, contextExecutor)
	return rowsAffected, err
}
