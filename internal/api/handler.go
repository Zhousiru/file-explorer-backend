package api

import (
	"github.com/Zhousiru/file-explorer-backend/internal/fsm"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
	"github.com/gin-gonic/gin"
)

func action(c *gin.Context) {
	action := c.Query("action")
	path := c.Param("path")

	switch action {
	case "list":
		list(c, path)
	case "del":
		del(c, path)
	case "rename":
		rename(c, path)
	case "move":
		move(c, path)
	default:
		c.JSON(400, Resp{
			Err: "invalid action",
		})
		return
	}
}

func list(c *gin.Context, path string) {
	if !util.IsDir(path) {
		c.JSON(400, Resp{
			Err: "it's a file",
		})
		return
	}

	f := new(fsm.Folder)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
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

func del(c *gin.Context, path string) {
	f := new(fsm.File)

	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	err = f.Delete()
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

func rename(c *gin.Context, path string) {
	if isValidQuery(c, "new") {
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

func move(c *gin.Context, path string) {
	if isValidQuery(c, "new") {
		return
	}
	newPath := c.Query("new")

	f := new(fsm.File)
	err := f.SetPath(path)
	if err != nil {
		c.JSON(400, Resp{
			Err: err.Error(),
		})
		return
	}

	err = f.Move(newPath)
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
