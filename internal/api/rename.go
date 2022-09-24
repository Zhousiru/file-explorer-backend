package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/gin-gonic/gin"
)

func rename(c *gin.Context, path string) {
	if isInvalidQuery(c, "new") {
		return
	}
	newName := c.Query("new")

	f := new(fsm.File)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	err = f.Rename(newName)
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
