package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maslias/webapp/cmd/customerror"
)

type Comment struct {
	Id        int64  `json:"id"`
	PostId    int64  `json:"post_id"`
	UserId    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type CommentsStore struct {
	db *sql.DB
}

func (s *CommentsStore) Create(ctx context.Context, comment *Comment) error {
	stmt := `insert into comments (post_id, user_id, content) values($1, $2, $3) returning id, created_at;`

	return s.db.QueryRowContext(ctx, stmt, comment.PostId, comment.UserId, comment.Content).
		Scan(&comment.Id, &comment.CreatedAt)
}

func (s *CommentsStore) GetById(ctx context.Context, commentId int64) (*Comment, error) {
	stmt := `select id, post_id, user_id, content, created_at
        from comments
        where id = $1;
    `

	var comment Comment

	err := s.db.QueryRowContext(ctx, stmt, commentId).
		Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, customerror.ErrNotFound
		default:
			return nil, err

		}
	}

	return &comment, nil
}

func (s *CommentsStore) GetByPostId(ctx context.Context, postId int64) ([]Comment, error) {
	stmt := `select c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id
            from comments c
            join users on users.id = c.user_id
            where c.post_id = $1
            order by c.created_at desc;
    `

	rows, err := s.db.QueryContext(ctx, stmt, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var c Comment
        c.User = User{}

		err := rows.Scan(
			&c.Id,
			&c.PostId,
			&c.UserId,
			&c.Content,
			&c.CreatedAt,
            &c.User.Username,
            &c.User.Id,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
