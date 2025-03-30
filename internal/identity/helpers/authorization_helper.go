package helpers

import (
	"fmt"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var mySigningKey = []byte(envService.GetEnvServiceInstance().Env.JWTSecret)

func GenerateToken(userid string) (string, error) {
	// Kullanıcı ID boş mu?
	if userid == "" {
		return "", fmt.Errorf("userid cannot be empty")
	}

	// Özel talepler (claims)
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token'ın geçerlilik süresi (24 saat)
		"userid":   userid,
		"role":     "user",
		"username": userid,
		"iat":      time.Now().Unix(),
	}

	// Token oluşturma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// İmzalama anahtarını kontrol et
	if mySigningKey == nil || len(mySigningKey) == 0 {
		return "", fmt.Errorf("signing key cannot be empty")
	}

	// Token'ı imzala ve string'e dönüştür
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenHandler(c *gin.Context, userid string) {

	if userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID gereklidir"})
		return
	}

	tokenString, err := GenerateToken(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token oluşturulamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func VerifyToken(tokenString string, allowedRoles []enum.Role) (jwt.MapClaims, error) {
	// Token'ı ayrıştır
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// İmzalama yöntemini kontrol et
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Beklenmeyen imzalama yöntemi: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Rolü claims'den al
		role, ok := claims["role"].(string)
		if !ok {
			return nil, fmt.Errorf("Rol bilgisi bulunamadı")
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
