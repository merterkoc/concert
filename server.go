package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
	identityservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"
	"golang.org/x/net/context"

	"time"

	_ "gilab.com/pragmaticreviews/golang-gin-poc/docs"
	externalController "gilab.com/pragmaticreviews/golang-gin-poc/external/controller"
	internalController "gilab.com/pragmaticreviews/golang-gin-poc/internal/controller"

	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/event-service"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	eventservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/event-service"
	"github.com/gin-contrib/cors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gilab.com/pragmaticreviews/golang-gin-poc/boot"
	eventDTO "gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	db                      = boot.DbStart()
	firebase                = boot.FirebaseStart()
	storageClient           = boot.FirebaseStorageStart()
	identityService         = identityservice.NewIdentityService(repository.NewIdentityRepository(db, firebase, storageClient), firebase)
	newExternalEventService = externalEventService.NewEventService(
		envService.GetEnvServiceInstance(),
	)
	newEventService = eventservice.NewEventService(
		repository.NewEventRepository(db),
	)
	NewIdentityController      = internalController.NewIdentityController(identityService)
	NewExternalEventController = externalController.NewEventController(newExternalEventService)
	eventController            = internalController.NewEventController(newEventService)
)

// @title GigBuddy API
// @version 1.0.0
// @description GigBuddy API Documentation
// @host localhost:8080
// @BasePath /v1

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization
// @tokenUrl https://example.com/oauth/token
// @scope.user Grants write access
// @scope.admin Grants read and write access to administrative information
func main() {
	// Start the server
	server := gin.Default()

	// CORS Middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authClient, err := firebase.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase auth client oluşturulamadı: %v", err)
	}
	envValue := envService.GetEnvServiceInstance().GetEnv()
	if envValue == "stage" {
		//authClient.UseEmulator("localhost", 9099)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(false)))
	server.POST("/v1/identity/verify", func(c *gin.Context) {
		var req dto.VerifyTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		NewIdentityController.VerifyToken(c, req)
	})
	server.POST("/v1/identity/create", func(c *gin.Context) {

		request := dto.CreateUserRequest{
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
			Username: c.PostForm("username"),
		}

		file, _ := c.FormFile("image")
		if err != nil {
			log.Printf("failed to get image file: %v", err)
		}

		request.Image = file
		user, err := NewIdentityController.CreateUser(c, request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Return the created user
		c.JSON(200, user)
	})
	server.GET("/v1/identity/userinfo", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
			return
		}
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		NewIdentityController.GetUserInfoById(c, parsedUID)
	})
	server.GET("/v1/events", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		var req eventDTO.GetEventRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		events, err := NewExternalEventController.FindByKeywordOrLocation(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, events)
	})
	server.POST("/v1/events/:eventId/join", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
			return
		}

		parsedUID, err := uuid.Parse(uid.(string))

		eventID := c.Param("eventId")

		if err != nil {
			fmt.Println("Unexpected UUID:", err)
			return
		}

		err = eventController.JoinEvent(parsedUID, eventID)
		if err != nil {
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "joined"})
	})
	server.POST("/v1/events/:eventId/leave", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
			return
		}
		eventID := c.Param("eventId")
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			fmt.Println("Unexpected UUID:", err)
			return
		}

		leaveErr := eventController.LeaveEvent(parsedUID, eventID)
		if leaveErr != nil {
			if errors.Is(leaveErr, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": leaveErr.Error()})
				return
			}
			c.JSON(500, gin.H{"error": leaveErr.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "left"})
	})
	server.GET("/v1/events/user", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
			return
		}

		parsedUID, err := uuid.Parse(uid.(string))
		event, err := eventController.GetEventByUser(parsedUID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, event)
	})
	env := envService.GetEnvServiceInstance()
	err = server.Run(":" + env.Env.AppPort)
	if err != nil {
		return
	}
}
func tokenMiddleware(firebaseAuth *auth.Client, allowedRoles []enum.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Context'i al
		ctx := c.Request.Context()

		// Token'ı header'dan al
		token, err := getTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}

		// Token'ı doğrula
		_, err = identityService.VerifyCustomToken(ctx, firebaseAuth, token, allowedRoles)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(token, "Bearer ")
		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["userid"])

		// Sonraki handler'a yönlendir
		c.Next()
	}
}
func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Algoritmayı doğrula
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("geçersiz imzalama yöntemi")
		}
		secretKey := []byte(envService.GetEnvServiceInstance().Env.JWTSecret)
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Claims'i döndür
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("geçersiz token")
}

func getTokenFromHeader(r *gin.Context) (string, error) {
	authHeader := r.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is missing")
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	return parts[1], nil
}
