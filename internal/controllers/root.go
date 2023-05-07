package controllers

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type errorResponse struct {
	Error string `json:"error"`
}

func InitRouter(engine *gin.Engine, db *gorm.DB, env initializers.Env) {
	var mediaClient = tmdb.NewMediaClient(env.TMDBApiKey)
	var mediaServiceGroup = engine.Group("/media-service")
	InitMediaDataController(mediaServiceGroup.Group("/media"), mediaClient)
}
