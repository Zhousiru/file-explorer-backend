package api

import (
	"path"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
	errorCode "github.com/Zhousiru/file-explorer-backend/internal/error_code"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context, target string) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, Resp{
			Err: "invalid parameters",
		})
		return
	}

	fullPath := path.Join(config.Get(config.K_ROOT), target, file.Filename)

	if util.IsExist(fullPath) {
		c.JSON(400, Resp{
			Err: errorCode.FileAlreadyExist,
		})
		return
	}

	err = c.SaveUploadedFile(file, fullPath)
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
