package controllers

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/initializers"
	"github.com/bingemate/media-service/internal/features"
	"github.com/bingemate/media-service/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type errorResponse struct {
	Error string `json:"error"`
}

func InitRouter(engine *gin.Engine, db *gorm.DB, env initializers.Env) {
	var mediaServiceGroup = engine.Group("/media-service")
	var mediaClient = tmdb.NewMediaClient(env.TMDBApiKey)
	var mediaRepository = repository.NewMediaRepository(db)
	var mediaData = features.NewMediaData(env.MovieTargetFolder, env.TvTargetFolder, mediaClient, mediaRepository)
	InitMediaDataController(mediaServiceGroup.Group("/media"), mediaData)
}
