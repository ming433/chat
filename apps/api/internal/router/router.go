package router

import (
	"chat/apps/api/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()
	r.GET("/ping", controller.GetUserInfo)
	r.GET("/user_list", controller.GetUserList)
	r.Run()
}
