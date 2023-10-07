package admin

import (
	"Go-basic/09-Gin+Gorm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//// 查询数据库
	//userList := []models.User{}
	//models.DB.Find(&userList)

	// 查询age大于19的用户
	userList := []models.User{}
	models.DB.Where("age>19").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

func (con UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "it",
		Age:      22,
		AddTime:  100000111,
		Email:    "222@qq.com",
	}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"result": "增加数据成功",
	})
}

func (con UserController) Edit(c *gin.Context) {
	//// 保存所有字段
	//// 查询
	//user := models.User{Id: 3}
	//models.DB.Find(&user)
	//// 更新
	//user.Username = "赵六"
	//user.Email = "444@gmail.com"
	//models.DB.Save(&user)

	//// 更新指定字段
	//user := models.User{}
	//models.DB.Model(&user).Where("id = ?", 3).Update("username", "赵六")

	user := models.User{}
	models.DB.Where("id = ?", 3).Find(&user)
	user.Username = "赵六"
	user.Email = "444@gmail.com"
	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"result": "修改用户成功",
	})
}

func (con UserController) Delete(c *gin.Context) {
	//user := models.User{Id: 3}
	//models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("username = ?", "张三").Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"result": "删除用户成功",
	})
}
