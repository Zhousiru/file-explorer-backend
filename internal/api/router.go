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
	r.GET("/*path", func(c *gin.Context) {
		action(c, "GET")
	})
	r.POST("/*path", func(c *gin.Context) {
		action(c, "POST")
	})

	return r.Run(addr)
}

func action(c *gin.Context, method string) {
	action := c.Query("action")
	target := c.Param("path")

	if isInvalidPath(c, target) {
		return
	}

	if method == "GET" {
		switch action {
		case "list":
			list(c, target)
		case "del":
			del(c, target)
		case "rename":
			rename(c, target)
		case "move":
			move(c, target)
		case "get":
			get(c, target)
		case "info":
			info(c, target)
		case "newFolder":
			newFolder(c, target)
		default:
			c.JSON(400, Resp{
				Err: "invalid action",
			})
			return
		}
	}

	if method == "POST" {
		switch action {
		case "upload":
			upload(c, target)
		default:
			c.JSON(400, Resp{
				Err: "invalid action",
			})
			return
		}
	}
}
