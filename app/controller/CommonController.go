package controller

import (
	"lucky/app/common"
	"lucky/app/model"
)

func GetUserIdFromDB(student_number string) int {
	userModel := model.User{}
	res := userModel.GetUserByStudentNumber(student_number)
	if res.Status == common.CodeError {
		return common.CodeError
	}
	userID := res.Data.(model.User).ID

	return userID
}

func GetUserSchoolFromDB(student_number string) int {
	userModel := model.User{}
	res := userModel.GetUserByStudentNumber(student_number)
	userSchool := res.Data.(model.User).School // 类型断言
	return userSchool
}
