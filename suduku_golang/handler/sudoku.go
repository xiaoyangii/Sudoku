package handler

import (
	"strconv"
	pack "sudu/handler/type"
	"sudu/service"

	"github.com/gin-gonic/gin"
)

func GetSudoku(c *gin.Context) {

	diff := c.Query("diffcult")
	resp := pack.NewSudoResponse()
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
	matrix := service.GetSudoku(int(diffcult))
	resp.SetSudoResponse(200, "successful", matrix)
	c.JSON(200, resp)

}
