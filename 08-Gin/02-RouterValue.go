package main

//import (
//	"encoding/xml"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type UserInfo struct {
//	Username string `json:"username" form:"username"`
//	Password string `json:"password" form:"password"`
//}
//
//type Article struct {
//	Title   string `json:"title" xml:"title"`
//	Content string `json:"content" xml:"content"`
//}
//
//func main() {
//	r := gin.Default()
//
//	r.GET("/", func(context *gin.Context) {
//		username := context.Query("username")
//		age := context.Query("age")
//		page := context.DefaultQuery("page", "1")
//		context.JSON(http.StatusOK, gin.H{
//			"username ": username,
//			"age":       age,
//			"page":      page,
//		})
//	})
//
//	// 请求传值
//	r.GET("/article", func(c *gin.Context) {
//		id := c.DefaultQuery("id", "1")
//
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "新闻详情",
//			"id":  id,
//		})
//	})
//
//	// 获取表单post数据
//	r.POST("/doAddUser", func(c *gin.Context) {
//		username := c.PostForm("username")
//		password := c.PostForm("password")
//		age := c.DefaultPostForm("age", "20")
//		c.JSON(http.StatusOK, gin.H{
//			"username": username,
//			"password": password,
//			"age":      age,
//		})
//	})
//
//	// 获取GET POST 传递的数据绑定到结构体
//	r.GET("/getUser", func(c *gin.Context) {
//		user := &UserInfo{}
//		if err := c.ShouldBind(&user); err == nil {
//			c.JSON(http.StatusOK, user)
//		} else {
//			c.JSON(http.StatusOK, gin.H{
//				"err": err.Error(),
//			})
//		}
//	})
//
//	r.POST("/doAddUser2", func(c *gin.Context) {
//		user := &UserInfo{}
//		if err := c.ShouldBind(&user); err == nil {
//			c.JSON(http.StatusOK, user)
//		} else {
//			c.JSON(http.StatusOK, gin.H{
//				"err": err.Error(),
//			})
//		}
//	})
//
//	// 获取Post Xml数据
//	r.POST("/xml", func(c *gin.Context) {
//		article := &Article{}
//		data, _ := c.GetRawData()
//		if err := xml.Unmarshal(data, &article); err == nil {
//			c.JSON(http.StatusOK, article)
//		} else {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"err": err.Error(),
//			})
//		}
//	})
//}
