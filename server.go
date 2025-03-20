package main

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/controller"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	eventService    service.EventService       = service.NewEventService()
	eventController controller.EventController = controller.New(eventService)
)

func main() {
	// Start the server
	server := gin.Default()

	server.GET("/events", func(c *gin.Context) {
		c.JSON(200, eventController.FindByKeyword("keyword"))

	})

	err := server.Run(":8081")
	if err != nil {
		return
	}
}
