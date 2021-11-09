package model

import (
	"lucky/app/common"
	"lucky/app/helper"

	"github.com/gin-gonic/gin"
)

type UserDesire struct {
	ID       int `json:"id" gorm:"id"`
	UserID   int `json:"user_id" gorm:"user_id" form:"user_id"`
	DesireID int `json:"wish_id" gorm:"desire_id" uri:"wish_id"`
}

func (UserDesire) TableName() string {
	return "user_desire"
}

func (model *UserDesire) AddUserDesire(data UserDesire) helper.ReturnType {

	err := db.Model(&UserDesire{}).Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加用户愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加用户愿望成功", Data: 0}

}

// 找出用户投递的愿望(分开成两部分了)
func (model *UserDesire) GetUserAllDesire(data UserDesire) helper.ReturnType {

	var desireIDs []int            // 先在user_desire中找出和用户有关的愿望的desire_id, 再到derire表中查出愿望详情
	var usered_desire []UserDesire // 在user_desire表中查找

	// var post_desire []Desire // 直接在user表中查找

	var Desires []Desire // 合并两部分查找的的愿望

	// 找出用户投递的愿望
	if err := db.Table("user_desire").Where("user_id = ?", data.UserID).Find(&usered_desire).Error; err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询用户desire_id失败", Data: err}
	}

	for _, value := range usered_desire {
		desireIDs = append(desireIDs, value.DesireID)
	}

	// 第一部分
	err := db.Table("desire").Where("id in (?)", desireIDs).Find(&Desires).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询用户愿望失败", Data: err}
	}

	// 找出用户点亮的愿望
	// 第二部分
	// if err := db.Table("desire").Where("light_user = ?", data.UserID).Find(&post_desire).Error; err != nil {
	// 	return helper.ReturnType{Status: common.CodeError, Msg: "查询用户失败", Data: err}
	// }

	// // 合并
	// Desires = append(Desires, post_desire...)

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取用户愿望成功", Data: gin.H{
		"wishes": Desires,
	}}

}

// 第二部分，找出用户点亮的愿望
func (model *UserDesire) GetUserAllDesire2(data UserDesire) helper.ReturnType {
	var post_desire []Desire

	if err := db.Table("desire").Where("light_user = ?", data.UserID).Find(&post_desire).Error; err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询用户失败", Data: err}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取用户愿望成功", Data: post_desire}
}

// 删除记录
func (model *UserDesire) DeleteUserDesire(data UserDesire) int {
	err := db.Where("user_id = ? AND desire_id = ?", data.UserID, data.DesireID).Delete(&data).Error
	if err != nil {
		return common.CodeError
	}
	return common.CodeSuccess
}

// 获取用户投递愿望的个数
func (model *UserDesire) GetUserWishCount(data UserDesire) int {
	count := 0
	err := db.Model(&UserDesire{}).Where("user_id = ?", data.UserID).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}

//
func (model *UserDesire) CheckUserDesire(data UserDesire) int {
	var userDesire UserDesire

	// 将user_id改成了ligh_user
	err := db.Where("user_id = ? AND id = ?", data.ID, data.DesireID).Find(&userDesire).Error
	if err != nil {
		return common.CodeError
	}
	return common.CodeSuccess
}

//
func (model *UserDesire) CheckUserDesire2(data UserDesire) int {
	var userDesire UserDesire

	// 将user_id改成了ligh_user
	err := db.Table("user_desire").Where("user_id = ? AND desire_id = ?", data.UserID, data.DesireID).Find(&userDesire).Error
	if err != nil {
		return common.CodeError
	}
	return common.CodeSuccess
}

// 获取用户id
func (model *UserDesire) GetUserIDbyWishID(data int) int {
	var userDesire UserDesire
	err := db.Model(&UserDesire{}).Where("desire_id = ?", data).Find(&userDesire).Error
	if err != nil {
		return -1
	}
	return userDesire.UserID
}
