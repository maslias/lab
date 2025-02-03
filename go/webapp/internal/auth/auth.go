package auth

import "github.com/golang-jwt/jwt"

type Authenticator interface {
	GenerateToken(jwt.Claims) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}
