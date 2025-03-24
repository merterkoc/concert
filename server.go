package main

import (
	"fmt"
	"net/http"

	_ "gilab.com/pragmaticreviews/golang-gin-poc/docs"
	controller "gilab.com/pragmaticreviews/golang-gin-poc/internal/delivery/http"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/service"

	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	utils "gilab.com/pragmaticreviews/golang-gin-poc/internal/helpers"
	"github.com/gin-gonic/gin"
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

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/events", func(c *gin.Context) {
		keyword := c.DefaultQuery("keyword", "")
		location := c.DefaultQuery("location", "")
		size, err := utils.ParseNullableInt(c.DefaultQuery("size", ""))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
			return
		}

		page, err := utils.ParseNullableInt(c.DefaultQuery("page", ""))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}

		// size ve page nullable (nil olabilir)
		if size != nil {
			fmt.Println("Size:", *size)
		} else {
			fmt.Println("Size is nil")
		}

		if page != nil {
			fmt.Println("Page:", *page)
		} else {
			fmt.Println("Page is nil")
		}

		events, err := eventController.FindByKeywordOrLocation(keyword, location, *size, *page)
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
