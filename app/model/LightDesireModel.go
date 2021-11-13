package model

import (
	"lucky/app/common"
	"lucky/app/helper"
)

type LightInfo struct {
	ID          int    `json:"id" gorm:"id"`
	DesireID    int    `json:"wish_id" gorm:"desire_id" form:"wish_id" uri:"wish_id"`
	LightName   string `json:"light_name,omitempty" gorm:"light_name"`
	LightQQ     string `json:"light_qq,omitempty" gorm:"light_qq"`
	LightWechat string `json:"light_wechat,omitempty" gorm:"light_wechat"`
	LightTel    string `json:"light_tel,omitempty" gorm:"light_tel"`
}

func (model *LightInfo) CreateLightInfo(lightInfo LightInfo) helper.ReturnType {
	err := db.Create(&lightInfo).Error
	if err != nil {
		return helper.ReturnRes(common.CodeError, "点亮愿望失败", "")
	}
	return helper.ReturnRes(common.CodeSuccess, "点亮愿望成功", "")
}

func (model *LightInfo) GetLightInfoByDesireID(lightInfo LightInfo) helper.ReturnType {
	var res LightInfo
	err := db.Where("desire_id = ?", lightInfo.DesireID).First(&res).Error
	if err != nil {
		return helper.ReturnRes(common.CodeError, "查询点亮信息失败", "")
	}
	return helper.ReturnRes(common.CodeSuccess, "查询点亮信息成功", res)
}

func (model *LightInfo) DeleteLighInfoByDesireID(desireID int) helper.ReturnType {
	err := db.Where("desire_id = ?", desireID).Delete(&LightInfo{}).Error
	if err != nil {
		return helper.ReturnRes(common.CodeError, "删除信息失败", "")
	}
	return helper.ReturnRes(common.CodeSuccess, "删除信息成功", "")
}
