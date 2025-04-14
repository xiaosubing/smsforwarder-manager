package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	JwtKey = []byte("admin")
	//TokenExpireDuration = time.Hour * 24
	AccessTokenExpire = time.Hour * 2
	//RefreshTokenExpire  = time.Hour * 24 * 7
)

type UserClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

//	func GenerateToken(id uint, name string, expireTime int64) (string, error) {
//		uc := UserClaims{
//			ID:   id,
//			Name: name,
//			RegisteredClaims: jwt.RegisteredClaims{
//				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireTime))),
//			},
//		}
//
//		token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
//		tokenString, err := token.SignedString(JwtKey)
//		if err != nil {
//			panic(err)
//		}
//		return tokenString, nil
//	}
func GenerateToken(id uint, name string) (string, error) {
	claims := UserClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
