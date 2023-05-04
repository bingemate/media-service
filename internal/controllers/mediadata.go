package controllers

import (
	"github.com/bingemate/media-service/pkg"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitMediaDataController(engine *gin.RouterGroup, mediaClient pkg.MediaClient) {
	engine.GET("/movie/:id", func(c *gin.Context) {
		getMovie(c, mediaClient)
	})
}

func getMovie(c *gin.Context, mediaClient pkg.MediaClient) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaClient.GetMovie(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, result)
}
