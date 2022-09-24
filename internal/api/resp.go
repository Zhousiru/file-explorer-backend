package api

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Payload interface{} `json:"payload"`
	Err     string      `json:"err"`
	Msg     string      `json:"msg"`
}

func isInvalidQuery(c *gin.Context, q ...string) bool {
	for _, el := range q {
		if c.Query(el) == "" {
			c.JSON(400, Resp{
				Err: "invalid parameters",
			})
			return true
		}
	}
	return false
}

func isInvalidPath(c *gin.Context, p string) bool {
	fullPath := path.Join(config.Get(config.K_ROOT), p)

	rel, err := filepath.Rel(config.Get(config.K_ROOT), fullPath)

	if err != nil || strings.HasPrefix(rel, "..") {
		c.JSON(400, Resp{
			Err: "invalid path",
		})
		return true
	}
	return false
}
