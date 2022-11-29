package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/grin-ch/dever-box-api/cfg"
)

type RoleBase struct {
	Id       int    `json:"id"`
	UUID     string `json:"uuid"`
	Avatar   string `json:"avator"` // 头像
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Time     int64  `json:"time"`
	Ip       string `json:"ip"`
}

type Cliams struct {
	RoleBase
	jwt.StandardClaims
}

func GenerateJWT(rBase RoleBase) (string, error) {
	now := time.Now()
	expire := now.Add(time.Duration(cfg.Config.Token.Expire) * time.Second)
	rBase.Time = now.Unix()
	claims := Cliams{
		RoleBase: rBase,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    cfg.Config.Token.Issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Config.Token.Signed))
	return token, err
}

func ParseJWT(token string) (*Cliams, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Config.Token.Signed), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Cliams); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
