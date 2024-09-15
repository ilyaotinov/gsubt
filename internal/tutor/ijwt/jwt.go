package ijwt

import (
	"fmt"
	"multiApp/internal/tutor/config"
	"multiApp/internal/tutor/model/user"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	cfg config.JWTConfig
}

func New(cfg config.JWTConfig) *JWT {
	return &JWT{cfg: cfg}
}

func (j *JWT) Generate(u *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		user.LoginKey: u.Login,
		user.IDKey:    u.ID,
	})

	tokenString, err := token.SignedString(j.cfg.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed Sign token: %w", err)
	}

	return tokenString, nil
}
