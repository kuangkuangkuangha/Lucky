package model

import (
	"log"
	"lucky/app/common"
	"lucky/app/helper"
	"time"
)

type Desire struct {
	ID            int       `json:"wish_id" gorm:"id" uri:"wish_id" form:"wish_id"`
	Desire        string    `json:"wish" gorm:"desire"`
	LightAt       time.Time `json:"light_at" gorm:"light_at"`
	CreatAt       time.Time `json:"creat_at" gorm:"creat_at"`
	WishmanName   string    `json:"wishman_name" gorm:"wishman_name"`
	WishmanQQ     string    `json:"wishman_qq" gorm:"wishman_qq"`
	WishmanWechat string    `json:"wishman_wechat" gorm:"wishman_wechat"`
	WishmanTel    string    `json:"wishman_tel" gorm:"wishman_tel"`
	LightUser     int       `json:"light_user" gorm:"light_user"`
	State         int       `json:"state" gorm:"state"`
	Type          int       `json:"type" gorm:"type" form:"categories"`
	School        int       `json:"school"`
}

// school 说明
// 1 武理
// 0 华师

// state 说明
// 0 未被点亮
// 1 以点亮未实现
// 2 已经被实现

func (Desire) TableName() string {
	return "desire"
}

func (model *Desire) AddDesire(data Desire) helper.ReturnType {

	err := db.Model(&Desire{}).Omit("creat_at", "light_at").Create(&data).Error // 没有.Error会报错

	if err != nil {
		log.Print("43", err)
		return helper.ReturnType{Status: common.CodeError, Msg: "添加愿望失败", Data: err}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加愿望成功", Data: data}
}

// 点亮愿望
func (model *Desire) LightDesire(data Desire) helper.ReturnType {

	var desire Desire

	err := db.Model(&Desire{}).Where("id = ?", data.ID).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败", Data: err}
	}

	if desire.State == 0 {
		desire.State = data.State
		desire.LightAt = time.Now()
		desire.LightUser = data.LightUser
		err := db.Model(&Desire{}).Update(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败,数据库错误", Data: err}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望成功", Data: desire}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望失败，该愿望已经被别人抢先点亮了", Data: ""}
}

func (model *Desire) CheckUserLishtDesire(data Desire) int {
	var desire Desire

	// 将user_id改成了ligh_user
	err := db.Table("desire").Where("light_user = ? AND id = ?", data.LightUser, data.ID).Find(&desire).Error
	if err != nil {
		return common.CodeError
	}
	return common.CodeSuccess
}

// 实现的是点亮人是自己的愿望，而不是实现自己的愿望
func (model *Desire) AchieveDesire(data Desire) helper.ReturnType {

	var desire Desire

	err := db.Model(&Desire{}).Where("id = ?", data.ID).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败", Data: err.Error()}
	}

	if desire.State == 1 {
		desire.State = 2
		//desire.LightAt = time.Now()
		desire.LightUser = data.LightUser
		err := db.Model(&Desire{}).Update(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败,数据库错误", Data: err.Error()}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望成功", Data: desire}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望失败，该愿望已经被别人抢先实现了", Data: ""}
}

func (model *UserDesire) AchieveDesire(data Desire) helper.ReturnType {

	var desire Desire

	err := db.Model(&Desire{}).Where("id = ?", data.ID).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败", Data: err.Error()}
	}

	if desire.State == 1 {
		desire.State = 2
		//desire.LightAt = time.Now()
		desire.LightUser = data.LightUser
		err := db.Model(&Desire{}).Update(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败,数据库错误", Data: err.Error()}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望成功", Data: desire}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望失败，该愿望已经被别人抢先实现了", Data: ""}
}

// 左右翻动查看单个愿望
func (model *Desire) GetWishByID(desireID int) helper.ReturnType {

	var desire Desire

	err := db.Model(&Desire{}).Where("id = ?", desireID).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查看愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查看愿望成功", Data: desire}

}

// 按分类查看愿望
func (model *Desire) GetWishByCategories(data Desire) helper.ReturnType {

	var desire []*Desire

	err := db.Model(&Desire{}).Where("type = ? AND state = ?", data.Type, 0).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查看愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查看愿望成功", Data: desire}

}

// 删除愿望
func (model *Desire) DeleteWish(data Desire) helper.ReturnType {

	err := db.Delete(&Desire{}, data.ID).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除愿望成功", Data: ""}

}

// 一次性返回10个愿望
func (model *Desire) Get10Wishes() helper.ReturnType {
	var res []*Desire

	err := db.Model(&Desire{}).Limit(10).Find(&res).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询愿望成功", Data: res}
}

// 获取所有愿望
func (model *Desire) GetAllWishes() helper.ReturnType {
	var res []Desire
	err := db.Model(&Desire{}).Where("state = ?", 0).Find(&res).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询愿望成功", Data: res}
}

// 取消实现愿望
func (model *Desire) CancelAchieveDesire(data Desire) helper.ReturnType {
	err := db.
		Model(&Desire{}).Where("id = ?", data.ID).
		Updates(map[string]interface{}{"light_user": -1, "state": 0}).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "取消点亮愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "取消点亮愿望成功", Data: nil}

}

// 获取用户点亮愿望的个数
func (model *Desire) GetUserLightCount(data Desire) int {
	count := 0
	err := db.Model(&Desire{}).Where("light_user = ?", data.LightUser).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}

// 获取用户点亮但未实现的愿望个数
func (model *Desire) GetUserWishBugNotReCount(data Desire) int {
	count := 0
	err := db.Model(&Desire{}).Where("light_user = ? AND state = ?", data.LightUser, common.WishHaveLight).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}
