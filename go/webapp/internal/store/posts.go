package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"

	"github.com/maslias/webapp/cmd/customerror"
)

type Post struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Tags      []string  `json:"tags"`
	Comments  []Comment `json:"comments"`
	// User      User      `json:"user"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, p *Post) error {
	stmt := `insert into posts( content, title, user_id, tags) values($1, $2, $3, $4) returning id, created_at, updated_at`
	err := s.db.QueryRowContext(ctx, stmt, p.Content, p.Title, p.UserId, pq.Array(p.Tags)).
		Scan(&p.Id, &p.CreatedAt, &p.UpdatedAt)

	return err
}

func (s *PostsStore) GetById(ctx context.Context, postId int64) (*Post, error) {
	var post Post
	// post.User = User{}
	stmt := `select p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, p.tags
    from posts p
    where p.id = $1;
    `

	err := s.db.QueryRowContext(ctx, stmt, postId).
		Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt, pq.Array(&post.Tags))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, customerror.ErrNotFound
		default:
			return nil, err

		}
	}

	return &post, nil
}
