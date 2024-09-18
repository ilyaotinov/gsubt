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
		"exp":         j.cfg.Exp,
	})

	tokenString, err := token.SignedString(j.cfg.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed Sign token: %w", err)
	}

	return tokenString, nil
}

func (j *JWT) Validate(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.cfg.SecretKey, nil
	})

	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("invalid token provided: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, fmt.Errorf("jwt claims expect to be map type")
	}

	return claims, nil
}
