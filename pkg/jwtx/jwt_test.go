package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	TOKEN_EXPIRED_TIME = time.Hour * 10
	WITHINLIMIT        = time.Second * 3600 * 20
)

var (
	_ CustomClaims = &CustomClaimsO{}
	_ SigningKeyer = (*SignKey)(nil)
)

type CustomClaimsO struct {
	UID   uint   `json:"uid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (CustomClaimsO) GetClaims() jwt.Claims {
	return &CustomClaimsO{
		UID:   1,
		Email: "123",
	}
}

func (cc CustomClaimsO) GetUid() uint {
	return cc.UID
}

func (cc CustomClaimsO) GetEmail() string {
	return cc.Email
}

type SignKey struct{}

func (s SignKey) SigningKey() []byte {
	return []byte("123123")
}

func TestGenerateToken(t *testing.T) {

	cc := &CustomClaimsO{}
	sk := &SignKey{}

	token, err := GenerateToken(cc, sk)
	assert.Equal(t, err, nil)
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	cc := &CustomClaimsO{}
	sk := &SignKey{}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImVtYWlsIjoiMTIzIn0.aNJPFlf_9HwByk_FSIJoj23ZAAzxavAk4iLJk0-L86M"
	claims, err := ParseToken(token, cc, sk)
	assert.Equal(t, err, nil)

	claimsObj, ok := claims.(*CustomClaimsO)

	assert.Equal(t, ok, true)

	assert.Equal(t, claimsObj.UID, uint(1))
	assert.Equal(t, claimsObj.Email, "123")

	t.Log(claimsObj.UID)
	t.Log(claimsObj.Email)
}

// TODO Refresh Token not test
