package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserInfo struct {
	jwt.StandardClaims
	UserName string `json:"userName"`
	Role     string `json:"role"`
}
