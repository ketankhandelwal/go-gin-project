package main

import (
	"gin-project-api/controllers"
	"gin-project-api/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService        = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	// server := gin.Default()

	server := gin.New()
	server.Use(gin.Recovery())

	server.GET("/getVideos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/addVideos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")

}
