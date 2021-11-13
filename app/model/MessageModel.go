package model

import (
	"lucky/app/common"
	"lucky/app/helper"
)

type Message struct {
	ID       int    `json:"message_id" gorm:"id"`
	Desireid int    `json:"desire_id" gorm:"column:desire_id"`
	Message  string `json:"message" gorm:"message"`
}

func (model *Message) LeaveMessage(data Message) helper.ReturnType {
	if err := db.Model(&Message{}).Create(&data).Error; err != nil {
		return helper.ReturnRes(common.CodeError, "留言失败", err.Error())
	} else {
		return helper.ReturnRes(common.CodeSuccess, "留言成功", data)
	}
}

func (model *Message) GetMessageByID(id int) helper.ReturnType {
	message := Message{}
	err := db.
		Model(&Message{}).
		Where("desire_id = ?", id).
		First(&message).
		Error

	if err != nil {
		return helper.ReturnRes(common.CodeError, "获取留言失败", err.Error())
	}
	return helper.ReturnRes(common.CodeSuccess, "获取留言成功", message)
}
