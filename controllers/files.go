package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tschechniker/goCloud/models"
)

// @Summary Find files
// @Description get files
// @Tags files
// @Accept  json
// @Produce  json
// @Success 200 {array} models.File
// @Router /v1/files [get]
func FindFiles(c *gin.Context) {
	var files []models.File
	models.DB.Find(&files)

	c.JSON(http.StatusOK, gin.H{"data": files})
}
