package api

import (
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

func isInvalidMethod(c *gin.Context, m string) bool {
	if c.Request.Method != m {
		c.JSON(400, Resp{
			Err: "invalid method",
		})
		return true
	}
	return false
}
