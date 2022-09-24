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
	r.Any("/*path", action)

	return r.Run(addr)
}

func action(c *gin.Context) {
	action := c.Query("action")
	path := c.Param("path")

	switch action {
	case "list":
		list(c, path)
	case "del":
		del(c, path)
	case "rename":
		rename(c, path)
	case "move":
		move(c, path)
	case "get":
		get(c, path)
	case "info":
		info(c, path)
	case "upload":
		upload(c, path)
	default:
		c.JSON(400, Resp{
			Err: "invalid action",
		})
		return
	}
}
