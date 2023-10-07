package main

import (
	"Go-basic/09-Gin+Gorm/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.AdminRoutersInit(r)
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
}
