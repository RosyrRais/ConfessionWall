package indexcontroller

import (
	"Wall/app/models"
	confessionservice "Wall/app/services/confessionService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var confessions *[]models.Confession
	//	var err gin.Error
	confessions, err := confessionservice.FindAll()
	if err != nil {
		utils.JsonResponseError(c, 100, "查询失败")
		return
	}

	utils.JsonResponseSuccess(c, confessions)
}
