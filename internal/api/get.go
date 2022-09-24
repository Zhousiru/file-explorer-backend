package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func get(c *gin.Context, path string) {
	if util.IsDir(path) {
		c.JSON(400, Resp{
			Err: "it's a folder",
		})
		return
	}

	f := new(fsm.File)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	fullPath := f.GetFullPath()
	c.FileAttachment(fullPath, f.GetFilename())
}
