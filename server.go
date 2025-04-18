package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"

	dto "gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity/enum"
	buddyservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/buddy-service"
	identityservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"
	"golang.org/x/net/context"

	"time"

	_ "gilab.com/pragmaticreviews/golang-gin-poc/docs"
	internalController "gilab.com/pragmaticreviews/golang-gin-poc/internal/controller"

	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/external-event-service"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/internal-event-service"
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
	identityRepo            = repository.NewIdentityRepository(db, firebase, storageClient)
	eventRepo               = repository.NewEventRepository(db)
	buddyRepo               = repository.NewBuddyRepository(db)
	identityService         = identityservice.NewIdentityService(identityRepo, firebase)
	newExternalEventService = externalEventService.NewEventService()
	newInternalEventService = internalEventService.NewEventService(
		*eventRepo,
		newExternalEventService,
		identityService,
	)
	buddyService          = buddyservice.NewBuddyService(db, buddyRepo, newInternalEventService)
	NewIdentityController = internalController.NewIdentityController(identityService)
	eventController       = internalController.NewEventController(newInternalEventService)
	buddyController       = internalController.NewBuddyController(buddyService)
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

	server.PATCH("/v1/identity/userinfo/interests", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		var patchUserInterestsRequest dto.PatchUserInterestsRequest
		if err := c.ShouldBindJSON(&patchUserInterestsRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
			return
		}
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		NewIdentityController.PatchUserInterests(c, parsedUID, patchUserInterestsRequest)
	})

	server.GET("/v1/identity/userinfo/interests", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		NewIdentityController.GetAllInterests(c)
	})

	server.GET("/v1/identity/profile/:id", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		profile, err := NewIdentityController.GetUserPublicProfileByID(c, id)
		if err != nil {
			return
		}
		c.JSON(200, profile)
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

		eventController.FindByKeywordOrLocation(c, req)

	})

	server.GET("/v1/events/:eventId", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		eventID := c.Param("eventId")
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		userID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		eventController.FindById(c, userID, eventID)
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
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		eventController.GetEventByUser(c, parsedUID)
	})

	server.GET("/v1/events/user/:userId", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
			return
		}
		parsedUserId, err := uuid.Parse(userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		eventController.GetEventByUserID(c, parsedUserId)
	})

	server.GET("/v1/buddy/requests", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
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
		res, err := buddyController.GetBuddyRequests(c, parsedUID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, res)
	})

	server.POST("/v1/buddy/requests", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		var req dto.CreateBuddyRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

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
		err = buddyController.CreateBuddyRequest(parsedUID, req)
		if err != nil {
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})

	server.POST("/v1/buddy/requests/:id/accept", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
		}
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		buddyRequestID := c.Param("id")
		buddyRequestUUID, err := uuid.Parse(buddyRequestID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = buddyController.AcceptBuddyRequest(parsedUID, buddyRequestUUID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})

	server.POST("/v1/buddy/requests/:id/reject", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
		}
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		buddyRequestID := c.Param("id")
		buddyRequestUUID, err := uuid.Parse(buddyRequestID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = buddyController.RejectBuddyRequest(parsedUID, buddyRequestUUID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})

	server.POST("/v1/buddy/requests/:id/block", tokenMiddleware(authClient, []enum.Role{enum.Admin, enum.User}), func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UID not found"})
		}
		parsedUID, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		buddyRequestID := c.Param("id")
		buddyRequestUUID, err := uuid.Parse(buddyRequestID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = buddyController.BlockBuddyRequest(parsedUID, buddyRequestUUID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
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
