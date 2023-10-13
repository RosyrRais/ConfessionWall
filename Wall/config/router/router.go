package router

import (
	confessioncontroller "Wall/app/controllers/confessionController"
	indexcontroller "Wall/app/controllers/indexController"
	"Wall/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/wall"
	wall := r.Group(pre)
	{
		wall.POST("/login", userController.Login)
		wall.POST("/signin", userController.Signin)
		wall.POST("/index", indexcontroller.Index)
		wall.POST("/post", confessioncontroller.PostConfession)
		wall.POST("/delete", confessioncontroller.DeleteConfession)
		wall.POST("/update", confessioncontroller.UpdateConfession)
	}
}
