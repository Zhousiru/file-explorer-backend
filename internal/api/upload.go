package api

import "github.com/gin-gonic/gin"

func upload(c *gin.Context, path string) {
	if isInvalidQuery(c, "filename") {
		return
	}
	if isInvalidMethod(c, "POST") {
		return
	}
	// filename := c.Query("filename")
}
