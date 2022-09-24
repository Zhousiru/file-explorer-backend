package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/gin-gonic/gin"
)

func info(c *gin.Context, path string) {
	f := new(fsm.File)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	modTime, err := f.GetFormattedModTime()
	if err != nil {
		c.JSON(500, Resp{
			Err: err.Error(),
		})
		return
	}

	c.JSON(200, Resp{
		Payload: gin.H{
			"modTime": modTime,
		},
	})
}
