package service

import (
	"fmt"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/configs"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(role string) (string, error) {

	claims := &models.Claims{
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(configs.SecretJWTKey)
	if err != nil {
		return "", fmt.Errorf("error signed string for jwt token")
	}

	return stringToken, nil
}
