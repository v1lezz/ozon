// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package notes_repository

import (
	"context"
)

type Querier interface {
	InsertNote(ctx context.Context, arg *InsertNoteParams) (int32, error)
	InsertNoteTag(ctx context.Context, arg *InsertNoteTagParams) error
	InsertTag(ctx context.Context, note string) error
	ListNotes(ctx context.Context, author string) ([]*ListNotesRow, error)
}

var _ Querier = (*Queries)(nil)
