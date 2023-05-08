package controllers

import (
	"errors"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
)

func InitFileInfoController(engine *gin.RouterGroup, fileInfo *features.MediaFile) {
	engine.GET(":id", func(c *gin.Context) {
		getMediaFileInfo(c, fileInfo)
	})
}

// @Summary		Get media file info
// @Description	Get media file info by Media ID
// @Tags			Media File
// @Param			id path int true "Media ID"
// @Produce		json
// @Success		200	{object} mediaFileResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media-file/{id} [get]
func getMediaFileInfo(c *gin.Context, mediaData *features.MediaFile) {
	id := c.Param("id")
	result, err := mediaData.GetMediaFileInfo(id)
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
