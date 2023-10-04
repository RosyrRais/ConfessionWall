package router

import (
	"Wall/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/wall"
	wall := r.Group(pre)
	{
		wall.POST("/login", userController.Login)
		wall.POST("/signin", userController.Signin)
	}
}
