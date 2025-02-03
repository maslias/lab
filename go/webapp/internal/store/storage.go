package store

import (
	"context"
	"database/sql"
	"time"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
	}
	Users interface {
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		ActivateFromInvitation(context.Context, string, time.Duration) error
		GetByEmail(context.Context, string) (*User, error)
		GetById(context.Context, int64) (*User, error)
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetById(context.Context, int64) (*Comment, error)
		GetByPostId(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Posts:    &PostsStore{db: db},
		Users:    &UserStore{db: db},
		Comments: &CommentsStore{db: db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
