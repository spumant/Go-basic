package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (con BaseController) success(c *gin.Context) {
	c.String(200, "成功")
}

func (con BaseController) error(c *gin.Context) {
	c.String(400, "失败")
}
