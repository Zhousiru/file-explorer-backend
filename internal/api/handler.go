package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func listSub(c *gin.Context) {
	path := c.Param("path")

	if !util.IsDir(path) {
		c.JSON(400, Resp{
			Payload: nil,
			Err:     "it's a file",
		})
		return
	}

	f := new(fsm.Folder)
	f.SetPath(path)

	subFile, err := f.GetSub(fsm.FLAG_FILE | fsm.FLAG_FOLDER)
	if err != nil {
		c.JSON(500, Resp{
			Payload: nil,
			Err:     err.Error(),
		})
		return
	}

	c.JSON(200, Resp{
		Payload: subFile,
		Err:     "",
	})
}
