package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// 解码token
type DecodedToken struct {
	Iat      int    `json:"iat"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	IsAdmin  int    `json:"isAdmin"`
}

// 生成token
func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}

// 验证token
func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	jsonErr := json.Unmarshal(jsonString, &decodedToken)
	if jsonErr != nil {
		log.Print(jsonErr)
	}

	return &decodedToken
}
