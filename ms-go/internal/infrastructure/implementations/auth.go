package implementations

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secret []byte
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{
		secret: []byte(secret),
	}
}

func (p *JWTProvider) Generate(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(p.secret)
}

func (p *JWTProvider) Validate(tokenString string) error {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return p.secret, nil
		},
	)
	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrTokenInvalidClaims
	}

	return nil
}
