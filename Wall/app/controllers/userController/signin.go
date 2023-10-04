package userController

import (
	"Wall/app/services/userService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SigninUser struct {
	//Name     string `json:"name" binding:"required"`
	//Password string `json:"password" binding:"required"`
	//ID       uint   `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Sex       string `json:"sex" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

func Signin(c *gin.Context) {
	//获取参数

	var data SigninUser
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonResponseError(c, 100, "获取参数时出现错误")
		return
	}

	//是否已存在

	err = userService.CheckUserExistByName(data.Name)
	if err == nil {
		utils.JsonResponseError(c, 101, "用户名已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonResponseOtherError(c)
		return
	}

	err = userService.CheckUserExistByEmail(data.Email)
	if err == nil {
		utils.JsonResponseError(c, 102, "邮箱已被注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonResponseOtherError(c)
		return
	}

	//密码是否一致

	flag := userService.ComparePassword(data.Password, data.Password2)
	if !flag {
		utils.JsonResponseError(c, 103, "密码不一致")
		return
	}

	//注册成功

	err = userService.InsertUser(data.Name, data.Email, data.Sex, data.Age, data.Password)
	if err != nil {
		utils.JsonResponseError(c, 104, "注册失败，请检查网络并重试")
		return
	}

	utils.JsonResponseSuccess(c, nil)
}
