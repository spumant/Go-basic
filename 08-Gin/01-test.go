package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default()
//
//	r.GET("/", func(context *gin.Context) {
//		context.String(200, "值：%v", "你好gin")
//	})
//
//	r.GET("/news", func(context *gin.Context) {
//		context.String(200, "我是新闻页面")
//	})
//
//	r.GET("/someJson", func(context *gin.Context) {
//		context.JSON(http.StatusOK, gin.H{"message": "hello world"})
//	})
//
//	r.GET("/moreJSON", func(c *gin.Context) {
//		// 方法二：使用结构体
//		var msg struct {
//			Name    string `json:"user"`
//			Message string
//			Age     int
//		}
//		msg.Name = "IT 营学院"
//		msg.Message = "Hello world!"
//		msg.Age = 18
//		c.JSON(http.StatusOK, msg)
//	})
//
//	r.Run(":8000")
//}
