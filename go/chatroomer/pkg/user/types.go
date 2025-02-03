package user

import (
	"context"

	"github.com/maslias/chatroomer/pkg/common"
)

type User struct {
	Id        int64           `json:"id"`
	Email     string          `json:"email"`
	Username  string          `json:"username"`
	Password  common.Password `json:"password"`
	CreatedAt string          `json:"created_at"`
}

// type password struct {
// 	text *string
// 	hash []byte
// }
//
// func (p *password) Set(text string) error {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
//
// 	p.text = &text
// 	p.hash = hash
//
// 	return nil
// }

type UserStore interface {
	Create(context.Context, *User) error
	GetById(context.Context, int64) (error, *User)
	GetByEmail(context.Context, string) (error, *User)
}
