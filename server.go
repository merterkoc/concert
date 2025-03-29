package main

import (
	"errors"
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	identityservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/identity-service"

	"time"

	_ "gilab.com/pragmaticreviews/golang-gin-poc/docs"
	externalController "gilab.com/pragmaticreviews/golang-gin-poc/external/controller"
	internalController "gilab.com/pragmaticreviews/golang-gin-poc/internal/controller"

	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/event-service"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	eventservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/event-service"
	"github.com/gin-contrib/cors"
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
	identityService         = identityservice.NewIdentityService(repository.NewIdentityRepository(db, firebase))
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

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/events/:id", func(c *gin.Context) {
		id := c.Param("id")
		event, err := NewExternalEventController.FindById(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, event)
	})
	server.POST("/identity/verify", func(c *gin.Context) {
		var req dto.VerifyTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := NewIdentityController.VerifyToken(c, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, token)
	})
	server.POST("/identity/create", func(c *gin.Context) {
		var req dto.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, err := NewIdentityController.CreateUser(c, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, user)
	})
	server.GET("/events", func(c *gin.Context) {
		var req eventDTO.GetEventRequest

		if err := c.ShouldBindQuery(&req); err != nil {
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
	server.POST("/events/:id/:eventId/join", func(c *gin.Context) {
		id := c.Param("id")
		eventID := c.Param("eventId")
		uid, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Unexpected UUID:", err)
			return
		}

		err = eventController.JoinEvent(uid, eventID)
		if err != nil {
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "joined"})
	})
	server.POST("/events/:id/:eventId/leave", func(c *gin.Context) {
		id := c.Param("id")
		eventID := c.Param("eventId")
		uid, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Unexpected UUID:", err)
			return
		}

		leaveErr := eventController.LeaveEvent(uid, eventID)
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
	server.GET("/events/:id/user", func(c *gin.Context) {
		id := c.Param("id")
		uid, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Invalid UUID:", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		event, err := eventController.GetEventByUser(uid)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, event)
	})
	env := envService.GetEnvServiceInstance()
	err := server.Run(":" + env.Env.AppPort)
	if err != nil {
		return
	}
}
