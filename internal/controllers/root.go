package controllers

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/initializers"
	"github.com/bingemate/media-service/internal/features"
	"github.com/bingemate/media-service/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(engine *gin.Engine, db *gorm.DB, env initializers.Env) {
	var mediaServiceGroup = engine.Group("/media-service")
	var mediaClient = tmdb.NewMediaClient(env.TMDBApiKey)
	var mediaRepository = repository.NewMediaRepository(db)
	var mediaData = features.NewMediaData(mediaClient, mediaRepository)
	var mediaFile = features.NewMediaFile(env.MovieTargetFolder, env.TvTargetFolder, mediaRepository)
	var mediaDiscover = features.NewMediaDiscovery(mediaClient, mediaRepository)
	var mediaAssetData = features.NewMediaAssetsData(mediaClient)
	var mediaCalendar = features.NewCalendarService(mediaClient, mediaRepository)
	InitMediaDataController(mediaServiceGroup.Group("/media"), mediaData)
	InitFileInfoController(mediaServiceGroup.Group("/media-file"), mediaFile)
	InitDiscoverController(mediaServiceGroup.Group("/discover"), mediaDiscover)
	InitCalendarController(mediaServiceGroup.Group("/calendar"), mediaCalendar)
	InitMediaAssetsController(mediaServiceGroup.Group("/assets"), mediaAssetData)
	InitPingController(mediaServiceGroup.Group("/ping"))
}
