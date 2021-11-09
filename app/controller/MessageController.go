package controller

import (
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 给愿望留言
func LeaveMessage(c *gin.Context) {
	var messageJson model.Message
	messageModel := model.Message{}

	if err := c.ShouldBindJSON(&messageJson); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
	}

	messageMap := helper.Struct2Map(messageJson)

	MessageValidator := validate.MessageValidate

	if boo, err := MessageValidator.ValidateMap(messageMap, "leave"); !boo {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	if res := messageModel.LeaveMessage(messageJson); res.Status == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}
}

// 获取用户的留言
func GetUserMessage(c *gin.Context) {
	messageJson := struct {
		DesireID int `json:"desire_id"`
	}{}

	if err := c.ShouldBindJSON(&messageJson); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	messageMap := helper.Struct2Map(messageJson)
	messageModel := model.Message{}
	MessageValidator := validate.MessageValidate

	if boo, err := MessageValidator.ValidateMap(messageMap, "get"); !boo {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := messageModel.GetMessageByID(messageJson.DesireID)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}
