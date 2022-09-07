package api

import (
	"fmt"
	"time"

	"github.com/Zhousiru/file-explorer-backend/internal/log"
	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	latency := time.Since(start)

	method := c.Request.Method
	statusCode := c.Writer.Status()

	if raw != "" {
		path = path + "?" + raw
	}

	log.Info(fmt.Sprintf("%s [%s] [%s %s] %d", log.GetColored("[API]", log.Green), latency, method, path, statusCode))

	c.Next()
}