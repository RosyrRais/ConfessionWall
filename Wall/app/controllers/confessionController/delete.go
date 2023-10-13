package confessioncontroller

import (
	confessionservice "Wall/app/services/confessionService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteData struct {
	//	Password2 string `json:"password2" binding:"required"`
	Id int `json:"id" binding:"required"`
	//	Text     string `json:"text" binding:"required"`
	//	Hidename int    `json:"hidename" binding:"required"`
	//	Open     int    `json:"open" binding:"required"`
	//Poster string `json:"poster" binding:"required"`
}

func DeleteConfession(c *gin.Context) {
	//获取参数

	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonResponseError(c, 100, "获取参数时出现错误")
		return
	}

	//是否存在

	err = confessionservice.CheckConfessionExistById(data.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonResponseError(c, 101, "帖子不存在")
		} else {
			utils.JsonResponseOtherError(c)
		}
		return
	}

	//删除成功

	err = confessionservice.DeleteConfessionById(data.Id)
	if err != nil {
		utils.JsonResponseError(c, 102, "删除时错误，请重试")
		return
	}

	utils.JsonResponseSuccess(c, nil)
}
