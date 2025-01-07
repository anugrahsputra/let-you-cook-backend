package jwt

import (
	"fmt"
	"let-you-cook/domain/model"
	"let-you-cook/utils/helper"

	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(helper.GetENV("JWT_SECRET", "secret"))

func GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (model.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return model.User{}, err
	}

	claims := token.Claims.(*jwt.MapClaims)

	// Safely extract values from claims
	var user model.User
	if id, ok := (*claims)["id"].(string); ok {
		user.Id = id
	} else {
		return model.User{}, fmt.Errorf("id claim is missing or not a string")
	}

	if username, ok := (*claims)["username"].(string); ok {
		user.Username = username
	} else {
		return model.User{}, fmt.Errorf("username claim is missing or not a string")
	}

	if email, ok := (*claims)["email"].(string); ok {
		user.Email = email
	} else {
		return model.User{}, fmt.Errorf("email claim is missing or not a string")
	}

	return user, nil
}
