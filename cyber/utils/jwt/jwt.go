package cutils_jwt

import (
	"errors"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/DrReMain/cyber-base-server/cyber"
	cutils_hd "github.com/DrReMain/cyber-base-server/cyber/utils/h_duration"
)

var (
	TokenInvalid = errors.New("登录授权无效, 请重新登录")
	TokenBuffer  = errors.New("登录授权需要刷新")
	TokenExpired = errors.New("登录授权已过期")
)

type BaseClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type CyberClaims struct {
	BufferTime *jwtv5.NumericDate
	BaseClaims
	jwtv5.RegisteredClaims
}

type JsonWebToken struct {
	SigningKey []byte
}

func NewJsonWebToken() *JsonWebToken {
	return &JsonWebToken{
		SigningKey: []byte(cyber.Config.Jwt.SigningKey),
	}
}

func (j *JsonWebToken) CreateClaims(bc BaseClaims) *CyberClaims {
	expire, _ := cutils_hd.ParseDuration(cyber.Config.Jwt.ExpiresTime)
	bf, _ := cutils_hd.ParseDuration(cyber.Config.Jwt.BufferTime)
	return &CyberClaims{
		BufferTime: jwtv5.NewNumericDate(time.Now().Add(bf)),
		BaseClaims: bc,
		RegisteredClaims: jwtv5.RegisteredClaims{
			Issuer:    cyber.Config.Jwt.Issuer,
			Audience:  jwtv5.ClaimStrings{"cyber"},
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(expire)),
			NotBefore: jwtv5.NewNumericDate(time.Now().Add(-1000)),
		},
	}
}

func (j *JsonWebToken) GenToken(claims *CyberClaims) (string, error) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, *claims)
	return token.SignedString(j.SigningKey)
}

func (j *JsonWebToken) ParseToken(tokenString string) (*CyberClaims, error) {
	t, err := jwtv5.ParseWithClaims(tokenString, &CyberClaims{}, func(token *jwtv5.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwtv5.ErrTokenExpired):
			if claims, ok := t.Claims.(*CyberClaims); ok && claims.BufferTime.After(time.Now()) {
				return nil, TokenBuffer
			}
			return nil, TokenExpired
		default:
			return nil, TokenInvalid
		}
	}

	if t != nil {
		if claims, ok := t.Claims.(*CyberClaims); ok && t.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}
