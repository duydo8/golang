package main

import (
	"gonew/entity"
	"gonew/middleware"
	"gonew/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	// VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery(), middleware.Logger())
	r.GET("/video", func(ctx *gin.Context) {
		ctx.JSON(200, videoService.FindAll())
	})
	r.POST("/video/post", func(ctx *gin.Context) {
		var video entity.Video
		ctx.BindJSON(&video)
		ctx.JSON(200, videoService.Save(video))
	})
	r.Run(":8080")
}
