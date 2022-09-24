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
	r.Use(Auth)
	r.GET("/*path", actionGet)
	r.POST("/*path", actionPost)

	return r.Run(addr)
}

func actionGet(c *gin.Context) {
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
	default:
		c.JSON(400, Resp{
			Err: "invalid action",
		})
		return
	}
}

func actionPost(c *gin.Context) {
	action := c.Query("action")
	path := c.Param("path")

	switch action {
	case "upload":
		upload(c, path)
	default:
		c.JSON(400, Resp{
			Err: "invalid action",
		})
		return
	}
}
