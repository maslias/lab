package user

import (
	"context"
	"database/sql"
)

type Store struct {
	*sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(ctx context.Context, user *User) error {
	stmt := `insert into users (username, password, email) values($1, $2, $3) returning id, created_at`
	return s.DB.QueryRowContext(ctx, stmt, user.Username, user.Password.Hash, user.Email).
		Scan(&user.Id, &user.CreatedAt)
}

func (s *Store) GetById(ctx context.Context, userId int64) (error, *User) {
	stmt := `select id, username, password, email, created_at from users where id = $1`

	user := &User{}
	if err := s.DB.QueryRowContext(ctx, stmt, userId).Scan(&user.Id, &user.Username, &user.Password.Hash, &user.Email, &user.CreatedAt); err != nil {
		return err, nil
	}

	return nil, user
}

func (s *Store) GetByEmail(ctx context.Context, userEmail string) (error, *User) {
	stmt := `select id, username, password, email, created_at from users where email = $1`

	user := &User{}
	if err := s.DB.QueryRowContext(ctx, stmt, userEmail).Scan(&user.Id, &user.Username, &user.Password.Hash, &user.Email, &user.CreatedAt); err != nil {
		return err, nil
	}

	return nil, user
}
