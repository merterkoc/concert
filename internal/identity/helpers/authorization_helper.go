package helpers

import (
	"fmt"
	"net/http"
	"time"

	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var mySigningKey = []byte(envService.GetEnvServiceInstance().Env.JWTSecret)

func GenerateToken(userid string, role enum.Role) (string, error) {
	if userid == "" {
		return "", fmt.Errorf("userid cannot be empty")
	}

	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"userid":   userid,
		"role":     role.String(),
		"username": userid,
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if mySigningKey == nil || len(mySigningKey) == 0 {
		return "", fmt.Errorf("signing key cannot be empty")
	}

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	bearerToken := "Bearer " + tokenString
	return bearerToken, nil
}

func GenerateTokenHandler(c *gin.Context, userid uuid.UUID, role enum.Role) {

	if userid == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID required"})
		return
	}

	tokenString, err := GenerateToken(userid.String(), role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func VerifyToken(tokenString string, allowedRoles []enum.Role) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role, ok := claims["role"].(string)
		if !ok {
			return nil, fmt.Errorf("Role not found.")
		}

		if len(allowedRoles) == 0 {
			return claims, nil
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole.String() {
				return claims, nil
			}
		}
	}

	return nil, fmt.Errorf("Invalid token")
}
