package userController

import (
	"Wall/app/models"
	"Wall/app/services/userService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//获取参数

	var data LoginUser
	err := c.ShouldBindJSON(&data)
	if err != nil {
		// c.JSON(500, gin.H{
		// 	"code": 100,
		// 	"msg":  "获取参数时出现错误",
		// 	"data": nil,
		// })
		utils.JsonResponseError(c, 100, "获取参数时出现错误")
		return
	}

	//是否存在

	err = userService.CheckUserExistByName(data.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// c.JSON(500, gin.H{
			// 	"code": 101,
			// 	"msg":  "用户不存在",
			// 	"data": nil,
			// })
			utils.JsonResponseError(c, 101, "用户不存在")
		} else {
			// c.JSON(500, gin.H{
			// 	"code": 102,
			// 	"msg":  "其他错误",
			// 	"data": nil,
			// })
			utils.JsonResponseOtherError(c)
		}
		return
	}

	//判断正误

	var user *models.User
	user, err = userService.GetUser(data.Name)
	if err != nil {
		// c.JSON(500, gin.H{
		// 	"code": 102,
		// 	"msg":  "其他错误",
		// 	"data": nil,
		// })
		utils.JsonResponseOtherError(c)
		return
	}

	flag := userService.ComparePassword(data.Password, user.Password)
	if !flag {
		// c.JSON(500, gin.H{
		// 	"code": 103,
		// 	"msg":  "密码错误",
		// 	"data": nil,
		// })
		utils.JsonResponseError(c, 102, "密码错误")
		return
	}

	//登录成功

	// c.JSON(200, gin.H{
	// 	"code": 200,
	// 	"msg":  "OK",
	// 	"data": user,
	// })
	utils.JsonResponseSuccess(c, user)
}
