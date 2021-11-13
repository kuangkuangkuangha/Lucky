package model

import (
	"fmt"
	"log"
	"lucky/app/common"
	"lucky/app/helper"

	"strconv"
)

type User struct {
	ID           int    `json:"id" gorm:"id"`
	IdcardNumber string `json:"idcard_number" gorm:"column:student_number"`
	Password     string `json:"password" gorm:"password"`
	School       int    `json:"school" gorm:"school"`
	Wechat       string `json:"wechat" gorm:"wechat"`
	Name         string `json:"name" gorm:"name"`
	Gender       int    `json:"gender" gorm:"gender"`
	Tel          string `json:"tel" gorm:"tel"`
	Email        string `json:"email" gorm:"email"`
	Major        string `json:"major" gorm:"major"`
}

func (User) TableName() string {
	return "user"
}

func (model *User) LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.Where("student_number = ? AND password = ?", data.IdcardNumber, data.Password).First(&user).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "登录验证成功", Data: user}
	}
}

func (model *User) BindEmail(data User) helper.ReturnType {

	err := db.Model(&User{}).Where("student_number = ?", data.IdcardNumber).Update(&data).Error

	if err != nil {
		return helper.ReturnRes(common.CodeError, "绑定邮箱失败", err.Error())
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: data.Email}
	}

}

func (model *User) CreateUser(user User) helper.ReturnType {
	err := db.Create(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建用户失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建用户成功", Data: user}
}

func (model *User) GetUserByRealName(realName string) helper.ReturnType {
	user := User{}
	err := db.
		Model(&User{}).
		Where("name = ?", realName).
		First(&user).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: user}
}

func (model *User) UpdateUserInfo(user User) helper.ReturnType {
	err := db.
		Model(&user).
		Select("email").
		Update(map[string]interface{}{"email": user.Email}).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "绑定邮箱失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: nil}

}

//  CCNU Login登陆验证
func (model *User) CcnuLogin(data User) helper.ReturnType {
	user := User{}

	info, err := GetUserInfoFormOne(data.IdcardNumber, data.Password)
	if err != nil {
		return helper.ReturnRes(common.CodeError, "登录失败", nil)
	}

	// res如果是个err，说明是第一次登入
	// res如果是和user结构，说明lucky数据库中存在用户，返回直接登入
	if res := db.Where("student_number = ?", data.IdcardNumber).First(&user); res.Error != nil {
		// _, _ := GetUserInfoFormOne(data.IdcardNumber, data.Password)
		// if err2 != nil {
		// 	return helper.ReturnRes(common.CodeError, "用户名或密码错误", res)
		// }

		xb, err3 := strconv.Atoi(info.User.Xb)
		if err3 != nil {
			fmt.Println("sex changed error")
		}
		// 存储用户信息
		data.Gender = xb
		data.School = 0
		data.Password = "" // 不存用户的密码
		data.Major = info.User.DeptName
		data.Name = info.User.Name

		if err := db.Model(&User{}).Create(&data).Error; err != nil {
			return helper.ReturnRes(common.CodeError, "添加用户失败", err)
		}
	}

	return helper.ReturnRes(common.CodeSuccess, "登陆成功", nil)
}

func (model *User) GetUserByStudentNumber(student_number string) helper.ReturnType {
	user := User{}

	err := db.
		Model(&User{}).
		Where("student_number = ?", student_number).
		First(&user).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: user}
}

// 根据用户id 获取邮箱
func (model *User) GetUserEmailByUserID(UserID int) string {
	var user User
	err := db.Model(&User{}).Where("id = ?", UserID).First(&user).Error
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return user.Email
}

// 根据用户id 获取姓名
func (model *User) GetUserNameByUserID(UserID int) string {
	var user User
	err := db.Model(&User{}).Where("id = ?", UserID).Find(&user).Error
	if err != nil {
		return ""
	}
	return user.Name
}

// 根据ID 获取用户信息
func (model *User) GetInfoByUserID(userId int) helper.ReturnType {
	var user User
	err := db.Model(&User{}).Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询user失败", Data: nil}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询user成功", Data: user}
}
