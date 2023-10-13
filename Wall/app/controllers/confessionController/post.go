package confessioncontroller

import (
	confessionservice "Wall/app/services/confessionService"
	"Wall/app/utils"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	//	Password2 string `json:"password2" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	Hidename int    `json:"hidename" binding:"required"`
	Open     int    `json:"open" binding:"required"`
	Poster   string `json:"poster" binding:"required"`
}

func PostConfession(c *gin.Context) {

	//获取参数

	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonResponseError(c, 100, "获取参数时出现错误")
		return
	}

	//发布成功

	err = confessionservice.InsertConfession(data.Title, data.Text, data.Hidename, data.Open, data.Poster)
	if err != nil {
		utils.JsonResponseError(c, 104, "发布失败，请检查网络并重试")
		return
	}

	utils.JsonResponseSuccess(c, nil)
}
