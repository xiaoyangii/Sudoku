package handler

import (
	"strconv"
	pack "sudu/handler/type"
	"sudu/service"

	"github.com/gin-gonic/gin"
)

func GetSudoku(c *gin.Context) {
	resp := pack.NewSudoResponse()
	diff := c.Query("diffcult")
	if diff == "" {
		resp.SetSudoResponse(40001, "empty diff", nil)

		c.JSON(40001, resp)
		return
	}
	diffcult, err := strconv.ParseInt(diff, 10, 64)
	if err != nil {
		resp.SetSudoResponse(40002, "parse diff error", nil)
		c.JSON(40002, resp)
		return
	}
	id := c.Query("id")
	if id == "" {
		resp.SetSudoResponse(40003, "empty id", nil)
		c.JSON(40003, resp)
		return
	}
	matrix, err := service.GetSudoku(c, int(diffcult), id)
	if err != nil {
		resp.SetSudoResponse(40005, err.Error(), nil)
		c.JSON(200, resp)
	}
	resp.SetSudoResponse(200, "successful", matrix)
	c.JSON(200, resp)

}

func SolveSuduku(c *gin.Context) {
	resp := pack.NewSudoResponse()
	id := c.Query("id")
	if id == "" {
		resp.SetSudoResponse(40003, "empty id", nil)
		c.JSON(40003, resp)
		return
	}
	solve, err := service.SolveSuduku(c, id)
	if err != nil {
		resp.SetSudoResponse(40006, "empty id", nil)
		c.JSON(40006, resp)
		return
	}
	resp.SetSudoResponse(200, "successful", string(solve))
	c.JSON(200, resp)
}
