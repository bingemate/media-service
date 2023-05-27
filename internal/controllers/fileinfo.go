package controllers

import (
	"errors"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitFileInfoController(engine *gin.RouterGroup, fileInfo *features.MediaFile) {
	engine.GET("file-tmdb/:id", func(c *gin.Context) {
		getMediaFileInfo(c, fileInfo)
	})
}

// @Summary		Get media file info by its Media TMDB ID
// @Description	Get media file info by its Media TMDB ID
// @Tags			Media File
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaFileResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media-file/file-tmdb/{id} [get]
func getMediaFileInfo(c *gin.Context, mediaData *features.MediaFile) {
	mediaID, err := strconv.Atoi(c.Param("id"))
	result, err := mediaData.GetMediaFileInfo(mediaID)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		if errors.Is(err, features.ErrInvalidMediaType) {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMediaFileResponse(result))
}
