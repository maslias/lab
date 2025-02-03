package store

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
	RoleId    int64    `json:"role_id"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	stmt := `insert into users (username, email, password, role_id) values($1, $2, $3, $4) returning id, created_at`
	err := tx.QueryRowContext(ctx, stmt, u.Username, u.Email, u.Password.hash, u.RoleId).
		Scan(&u.Id, &u.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) CreateUserInvitation(
	ctx context.Context,
	tx *sql.Tx,
	userId int64,
	token string,
	exp time.Duration,
) error {
	stmt := `insert into user_invitations (token, user_id, expire) values($1, $2, $3)`

	_, err := tx.ExecContext(ctx, stmt, token, userId, time.Now().Add(exp))
	return err
}

func (s *UserStore) DeleteUserInvitation(ctx context.Context, tx *sql.Tx, userId int64) error {
	stmt := `delete from user_invitations where user_id = $1;`
	_, err := tx.ExecContext(ctx, stmt, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) CreateAndInvite(
	ctx context.Context,
	user *User,
	token string,
	exp time.Duration,
) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		if err := s.CreateUserInvitation(ctx, tx, user.Id, token, exp); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) ActivateFromInvitation(
	ctx context.Context,
	token string,
	exp time.Duration,
) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		user, err := s.GetUserFromInvitationToken(ctx, tx, token, exp)
		if err != nil {
			return err
		}

		user.IsActive = true
		if err := s.Update(ctx, tx, user); err != nil {
			return nil
		}
		if err := s.DeleteUserInvitation(ctx, tx, user.Id); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) GetUserFromInvitationToken(
	ctx context.Context,
	tx *sql.Tx,
	token string,
	exp time.Duration,
) (*User, error) {
	stmt := `select id, username, email, password, created_at, is_active from users join user_invitations ui on ui.user_id = id where ui.token = $1 and ui.expire > $2`

	user := &User{}
	err := tx.QueryRowContext(ctx, stmt, token, time.Now()).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password.hash, &user.CreatedAt, &user.IsActive)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserStore) Update(ctx context.Context, tx *sql.Tx, user *User) error {
	stmt := `update users set username = $1, email = $2, is_active = $3 where id = $4`
	_, err := tx.ExecContext(ctx, stmt, user.Username, user.Email, user.IsActive, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	stmt := `select id, username, email, password, created_at, role_id from users where email = $1 and is_active = true;`

	user := &User{}

	if err := s.db.QueryRowContext(ctx, stmt, email).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password.hash, &user.CreatedAt, &user.RoleId); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserStore) GetById(ctx context.Context, userId int64) (*User, error) {
	stmt := `select id, username, email, password, created_at, role_id from users where id = $1 and is_active = true;`
	user := &User{}
	if err := s.db.QueryRowContext(ctx, stmt, userId).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password.hash, &user.CreatedAt, &user.RoleId); err != nil {
		return nil, err
	}

	return user, nil
}
