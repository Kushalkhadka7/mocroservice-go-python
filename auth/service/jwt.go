package service

import (
	"auth/model"
	"fmt"
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

func NewJWTMnanager(secretKey string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

func (jwtManager *JWTManager) GenerateToken(user *model.User) (string, error) {
	claims := UserInfo{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtManager.tokenDuration).Unix(),
		},
		UserName: user.Name,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtManager.secretKey))
}

func (jwtManager *JWTManager) VerifyToken(accessToken string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserInfo{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected singing token")
			}

			return []byte(jwtManager.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Invalid token %w", err)
	}

	claims, ok := token.Claims.(*UserInfo)
	if !ok {
		return nil, fmt.Errorf("Unexpected singing token")
	}

	return claims, nil
}
