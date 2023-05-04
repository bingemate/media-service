package controllers

import (
	"github.com/bingemate/media-service/initializers"
	"github.com/bingemate/media-service/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type errorResponse struct {
	Error string `json:"error"`
}

func InitRouter(engine *gin.Engine, db *gorm.DB, env initializers.Env) {
	var mediaClient = pkg.NewMediaClient(env.TMDBApiKey)
	InitMediaDataController(engine.Group("/media"), mediaClient)
}
