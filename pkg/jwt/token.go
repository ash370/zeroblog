package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func BuildTokens(opt TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix() //令牌签发时间需要稍早于当前时刻，避免时间差导致验证问题
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = now + opt.AccessExpire

	return token, nil
}

func genToken(iat int64, sk string, payloads map[string]interface{}, exp int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256) //HMAC SHA-256
	token.Claims = claims
	return token.SignedString([]byte(sk))
}
