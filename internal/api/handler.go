package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/log"
	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	log.Info(c.Param("path"))
}
