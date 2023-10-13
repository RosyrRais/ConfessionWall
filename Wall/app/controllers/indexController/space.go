package indexcontroller

import (
	"Wall/app/models"
	confessionservice "Wall/app/services/confessionService"
	"Wall/app/services/userService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SpaceData struct {
	Name string `json:"name" binding:"required"`
}

func Space(c *gin.Context) {

	var data SpaceData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonResponseError(c, 100, "获取参数时出现错误")
		return
	}

	err = userService.CheckUserExistByName(data.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonResponseError(c, 101, "用户不存在")
		} else {
			utils.JsonResponseOtherError(c)
		}
		return
	}

	var confessions *models.Confession
	confessions, err = confessionservice.FindConfessionsByName(data.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonResponseError(c, 102, "帖子不存在")
		} else {
			utils.JsonResponseOtherError(c)
		}
		return
	}

	utils.JsonResponseSuccess(c, confessions)
}
