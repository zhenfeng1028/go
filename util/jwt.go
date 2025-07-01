package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lunny/log"
)

// https://supertokens.com/blog/what-is-jwt

func main() {
	signedToken := EncJwt("abcd123")
	userId, err := DecJwt(signedToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(userId)
}

func EncJwt(uid string) string {
	key := []byte("jwt_key")

	// 序列化时会对key排序
	claims := jwt.MapClaims{
		"userId": uid,
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		log.Error("Failed to sign token:", uid, err)
		return ""
	}
	return signedToken
}

func DecJwt(signedToken string) (uid string, err error) {
	key := []byte("jwt_key")

	// 验证JWT
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		log.Error("Failed to parse token:", signedToken, err)
		return
	}

	// 验证Claim
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		uid = claims["userId"].(string)
	} else {
		log.Error("Failed to parse claims")
	}
	return
}
