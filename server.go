package main

import (
	"database/sql"
	"time"

	_ "gilab.com/pragmaticreviews/golang-gin-poc/docs"
	controller "gilab.com/pragmaticreviews/golang-gin-poc/internal/delivery/http"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/service"
	"github.com/gin-contrib/cors"

	//boot "gilab.com/pragmaticreviews/golang-gin-poc/boot"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	eventService service.EventService = service.NewEventService(
		envService.GetEnvServiceInstance(),
	)
	eventController controller.EventController = controller.New(eventService)
)

func main() {
	// Start the server
	server := gin.Default()

	DbStart()

	// CORS Middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Tüm kaynaklara izin ver (güvenlik için kısıtlayabilirsiniz)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/events/:id", func(c *gin.Context) {
		id := c.Param("id")
		event, err := eventController.FindById(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, event)
	})
	server.GET("/events", func(c *gin.Context) {
		var req dto.GetEventRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		events, err := eventController.FindByKeywordOrLocation(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, events)
	})
	env := envService.GetEnvServiceInstance()
	err := server.Run(":" + env.Env.AppPort)
	if err != nil {
		return
	}
}

func DbStart() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gigbuddy?parseTime=true")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
