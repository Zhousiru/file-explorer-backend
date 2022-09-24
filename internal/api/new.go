package api

import (
	"os"
	"path"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
	errorCode "github.com/Zhousiru/file-explorer-backend/internal/error_code"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func newFolder(c *gin.Context, target string) {
	if isInvalidQuery(c, "name") {
		return
	}
	name := c.Query("name")

	fullPath := path.Join(config.Get(config.K_ROOT), target, name)

	if util.IsExist(fullPath) {
		c.JSON(400, Resp{
			Err: errorCode.FileAlreadyExist,
		})
		return
	}

	err := os.MkdirAll(fullPath, os.ModePerm)
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
