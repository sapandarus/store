package util

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"store/model/user"
)

func init() {}

func VerifyToken(tokenString string) jwt.MapClaims {
	if tokenString == "" {
		return nil
	}

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	hmacSampleSecret := []byte(config.Secret)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

func CreateToken(user user.User) string {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	hmacSampleSecret := []byte([]byte(config.Secret))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  user.ID,
		"role":    user.Role,
		"storeId": user.StoreId,
		"nbf":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	return tokenString
}
