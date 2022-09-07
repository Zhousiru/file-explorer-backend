package api

import (
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Payload interface{} `json:"payload"`
	Err     string      `json:"err"`
	Msg     string      `json:"msg"`
}

func isValidQuery(c *gin.Context, q ...string) bool {
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
