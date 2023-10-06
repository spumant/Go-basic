package routers

import (
	"Go-basic/08-Gin/controllers/it"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/default")
	{
		defaultRouters.GET("/", it.DefaultController{}.Index)
		defaultRouters.GET("/news", it.DefaultController{}.News)
		defaultRouters.GET("/delete", it.DefaultController{}.DeleteCookie)
	}
}
