package confessioncontroller

import (
	confessionservice "Wall/app/services/confessionService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateData struct {
	Id       int    `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	Hidename int    `json:"hidename" binding:"required"`
	Open     int    `json:"open" binding:"required"`
	Poster   string `json:"poster" binding:"required"`
}

func UpdateConfession(c *gin.Context) {
	//获取参数

	var data UpdateData
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

	//修改成功

	err = confessionservice.UpdateConfessionById(data.Id, data.Title, data.Text, data.Hidename, data.Open, data.Poster)
	if err != nil {
		utils.JsonResponseError(c, 102, "修改时错误，请重试")
		return
	}

	utils.JsonResponseSuccess(c, nil)
}
