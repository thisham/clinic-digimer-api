package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Role string

const (
	DOCTOR Role = "Doctor"
	ADMIN  Role = "Admin"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Role Role
}

func GenerateJwt(userId string, role Role) (token string, err error) {
	claims := JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        userId,
		},
		Role: role,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = rawToken.SignedString([]byte("configs.GetJwtSecret().SecretKey"))
	return
}

func SetJwtCookie(ec echo.Context, token string) {
	authCookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	ec.SetCookie(&authCookie)
}
