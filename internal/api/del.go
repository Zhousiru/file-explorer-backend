package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/gin-gonic/gin"
)

func del(c *gin.Context, path string) {
	f := new(fsm.File)

	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	err = f.Delete()
	if err != nil {
		c.JSON(500, Resp{
			Err: err.Error(),
		})
		return
	}

	c.JSON(200, Resp{
		Msg: "ok",
	})
}
