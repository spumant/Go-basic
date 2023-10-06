package it

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, false)

	session := sessions.Default(c)
	// 配置session过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6小时
	})
	session.Set("username", "张三 111")
	session.Save()

	c.String(200, "首页")
}

func (con DefaultController) News(c *gin.Context) {
	//// cookie
	//username, _ := c.Cookie("username")
	// session
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "username=%v", username)
}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "张三", -1, "/", "localhost", false, false)
	c.String(200, "删除成功")
}
