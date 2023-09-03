package jwtx

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrTokenNotExpired = errors.New("Token not expired.")
)

type SigningKeyer interface {
	SigningKey() []byte
}

type CustomClaims interface {
	GetClaims() jwt.Claims
}

func GenerateToken(claims CustomClaims, sk SigningKeyer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.GetClaims())
	tokenStr, err := token.SignedString(sk.SigningKey())
	return tokenStr, err
}

func ParseToken(tokenStr string, claims CustomClaims, sk SigningKeyer) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims.GetClaims(),
		func(token *jwt.Token) (interface{}, error) {
			return sk.SigningKey(), nil
		})
	claims, ok := token.Claims.(CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func RefreshToken(claims CustomClaims, sk SigningKeyer, limit time.Duration) (string, error) {
	expiredTime, err := claims.GetClaims().GetExpirationTime()
	if err != nil {
		return "", err
	}
	if time.Since(expiredTime.Time) < limit {
		return GenerateToken(claims, sk)
	}
	return "", ErrTokenNotExpired
}

//func (j *JWT) GetType() string {
//	return "brear"
//}
