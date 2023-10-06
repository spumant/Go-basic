package routers

import (
	"Go-basic/08-Gin/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
	}
}
