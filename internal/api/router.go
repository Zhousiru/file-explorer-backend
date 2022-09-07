package api

import (
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func StartServer(addr string) error {
	r := gin.New()
	r.Use(Logger)
	r.Use(Cors)
	r.GET("/*path", action)

	return r.Run(addr)
}
