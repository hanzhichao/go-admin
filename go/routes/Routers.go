package routes

import (
	"github.com/gin-gonic/gin"
	"go-admin/go/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// 用户路由组
	userGroup := r.Group("api")
	{
		// 新建用户
		userGroup.POST("/users", controller.CreateUser)
		// 获取用户列表
		userGroup.GET("/users", controller.GetUserList)
		// 获取用户信息
		userGroup.GET("/users/:id", controller.GetUserById)
		// 更新用户信息
		userGroup.PUT("/users/:id", controller.UpdateUser)
		// 删除用户
		userGroup.DELETE("/users/:id", controller.DeleteUser)
	}
	return r

}