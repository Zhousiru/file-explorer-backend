package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func list(c *gin.Context, path string) {
	f := new(fsm.Folder)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	if !util.IsDir(path) {
		c.JSON(400, Resp{
			Err: "it's a file",
		})
		return
	}

	subFile, err := f.GetSub(fsm.FLAG_FILE | fsm.FLAG_FOLDER)
	if err != nil {
		c.JSON(500, Resp{
			Err: err.Error(),
		})
		return
	}

	c.JSON(200, Resp{
		Payload: subFile,
	})
}
