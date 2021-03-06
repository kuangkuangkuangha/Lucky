package controller

import (
	"log"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 投递愿望
func AddDesire(c *gin.Context) {
	var desireJson model.Desire
	desireValidate := validate.DesireValidate
	desireModel := model.Desire{}
	userDesireModel := model.UserDesire{}
	var userDesireJson model.UserDesire

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}
	UserID := c.MustGet("user_id").(int)
	school := c.MustGet("school").(int) // GetInt?
	desireJson.School = school
	userDesireJson.UserID = UserID

	postWishCount := userDesireModel.GetUserWishCount(userDesireJson)

	if postWishCount >= 5 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "许愿次数已达上限", ""))
		return
	}

	desireJson.CreatAt = time.Now().In(common.ChinaTime)
	res := desireModel.AddDesire(desireJson)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	userDesireJson.DesireID = res.Data.(model.Desire).ID
	res2 := userDesireModel.AddUserDesire(userDesireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res2.Status, res2.Msg, gin.H{"wish_id": userDesireJson.DesireID}))

}

// 点亮愿望
func LightDesire(c *gin.Context) {
	var desireJson model.Desire
	var lightDesireJson model.LightInfo
	desireModel := model.Desire{}
	desireValidate := validate.DesireValidate
	lightDesireModel := model.LightInfo{}

	if err := c.ShouldBindJSON(&lightDesireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err))
		return
	}

	desireMap := helper.Struct2Map(lightDesireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "light"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err))
		return
	}

	desireJson.LightUser = c.MustGet("user_id").(int)
	userLightCount := desireModel.GetUserLightCount(desireJson)

	// check self desire
	userDesireModel := model.UserDesire{}
	var userDesireTemp model.UserDesire
	userDesireTemp.UserID = desireJson.LightUser
	userDesireTemp.DesireID = lightDesireJson.DesireID
	checkCode := userDesireModel.CheckUserDesire2(userDesireTemp)
	log.Println(checkCode)
	log.Println(userDesireTemp)
	if checkCode == common.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "不能点亮自己的愿望哦", ""))
		return
	}

	if desireModel.GetUserWishBugNotReCount(desireJson) >= 2 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "已经点亮了两个愿望了", ""))
		return
	}

	if userLightCount >= 7 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "已经点亮过七个愿望了", ""))
		return
	}

	desireInfo := desireModel.GetWishByID(lightDesireJson.DesireID).Data.(model.Desire)
	if desireInfo.State != 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "愿望已经被抢先实现了", ""))
		return
	}

	desireJson.ID = lightDesireJson.DesireID
	var userModel model.User
	desireJson.State = 1

	res := desireModel.LightDesire(desireJson)
	desireUserID := userDesireModel.GetUserIDbyWishID(lightDesireJson.DesireID)
	res = lightDesireModel.CreateLightInfo(lightDesireJson)
	userEmail := userModel.GetUserEmailByUserID(desireUserID)
	log.Println(userEmail)
	go func() {
		mailErr, err2 := helper.SendMail(userEmail, common.LightWish, "", "")
		log.Println(mailErr)
		if err2 != nil {
			log.Print(err2)
		}
	}()

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 获取用户的愿望池
func GetUserPostDesire(c *gin.Context) {
	var userDesireJson model.UserDesire
	desireModel := model.UserDesire{}
	userDesireValidate := validate.UserDesireValidate

	// if err := c.ShouldBindQuery(&userDesireJson); err != nil {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
	// 	return
	// }

	userDesireJson.UserID = c.MustGet("user_id").(int)

	userDesireMap := helper.Struct2Map(userDesireJson)

	if res, err := userDesireValidate.ValidateMap(userDesireMap, "getUser"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.GetUserAllDesire(userDesireJson)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 获取用户的愿望池
func GetUserLightDesire(c *gin.Context) {
	var userDesireJson model.UserDesire
	desireModel := model.UserDesire{}
	// userDesireValidate := validate.UserDesireValidate

	// if err := c.ShouldBindQuery(&userDesireJson); err != nil {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
	// 	return
	// }

	userDesireJson.UserID = c.MustGet("user_id").(int)

	// userDesireMap := helper.Struct2Map(userDesireJson)

	// if res, err := userDesireValidate.ValidateMap(userDesireMap, "getUser"); !res {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
	// 	return
	// }

	res := desireModel.GetUserAllDesire2(userDesireJson)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 通过点击一个愿望查看 该愿望 详情
func GetWishDetails(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望ID
	desireJson := model.Desire{}

	// 初始化一个 验证器 用来校验数据格式
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindQuery(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "byid"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.GetWishByID(desireJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

// 通过点击分类查看 同种愿望
func GetWishByCatagories(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望type
	var desireJson model.Desire

	// 初始化一个 验证器 用来校验数据格式
	// desireValidate := validate.DesireValidate

	if err := c.ShouldBindQuery(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	// desireMap := helper.Struct2Map(desireJson)

	// if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
	// 	return
	// }

	res := desireModel.GetWishByCategories(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

// 删除愿望
func DeleteWish(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望Type
	var desireJson model.Desire
	lightInfoModel := model.LightInfo{}
	// 初始化一个 验证器 用来校验数据格式
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindQuery(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	userID := c.MustGet("user_id").(int)
	userDesireModel := model.UserDesire{}
	var userDesireJson model.UserDesire
	userModel := model.User{}
	desireInfo := desireModel.GetWishByID(desireJson.ID).Data.(model.Desire)

	userDesireJson.DesireID = desireJson.ID
	userDesireJson.UserID = userID

	checkres := userDesireModel.CheckUserDesire2(userDesireJson)
	if checkres == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "这不是你的愿望欧", ""))
		return
	}

	if desireInfo.State == common.WishHaveLight {
		userEmail := userModel.GetUserEmailByUserID(desireInfo.LightUser)
		go func() {
			mailErr, err2 := helper.SendMail(userEmail, common.DeleteWish, desireInfo.Desire, "")
			log.Println(mailErr)
			if err2 != nil {
				log.Print(err2)
			}
		}()
	}
	_ = userDesireModel.DeleteUserDesire(userDesireJson)

	res := lightInfoModel.DeleteLighInfoByDesireID(desireJson.ID)
	res = desireModel.DeleteWish(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

// 测试返回token
func Token(c *gin.Context) {
	student_number := c.MustGet("student_number").(string)
	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, "studnt_number,", student_number))
}

// 取消
func CancelLightDesire(c *gin.Context) {
	desireModel := model.Desire{}
	desireValidate := validate.DesireValidate
	lightInfoModel := model.LightInfo{}
	cancelDesireJson := struct {
		Message string `json:"message"`
		ID      int    `json:"wish_id"`
	}{}

	if err := c.ShouldBindJSON(&cancelDesireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(cancelDesireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "cancel"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	// userDesireJson.DesireID = desireJson.ID
	// userDesireJson.UserID = desireJson.LightUser
	var desireJson model.Desire
	desireJson.ID = cancelDesireJson.ID
	// _ = userDesireModel.DeleteUserDesire(userDesireJson)
	requestUesrID := c.MustGet("user_id").(int)

	var userModel model.User
	var userDesireModel model.UserDesire
	desireInfo := desireModel.GetWishByID(desireJson.ID).Data.(model.Desire)

	if desireInfo.State == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "这个愿望还没被点亮", ""))
		return
	}

	if desireInfo.LightUser != requestUesrID {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "这个愿望不是你点亮的欧", ""))
		return
	}

	if desireInfo.State == common.WishHaveRealize {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "这个愿望已经被确认实现啦", ""))
		return
	}

	desireUserID := userDesireModel.GetUserIDbyWishID(desireJson.ID)
	userEmail := userModel.GetUserEmailByUserID(desireUserID)
	res := lightInfoModel.DeleteLighInfoByDesireID(desireJson.ID)
	res = desireModel.CancelAchieveDesire(desireJson)
	go func() {
		mailErr, err2 := helper.SendMail(userEmail, common.CancelLight, desireInfo.Desire, cancelDesireJson.Message)
		log.Println(mailErr)
		if err2 != nil {
			log.Print(err2)
		}
	}()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func AchieveDesire(c *gin.Context) {
	var desireJson model.Desire
	desireModel := model.Desire{}
	DesireModel := model.Desire{}
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.GetWishByID(desireJson.ID)
	desireInfo := res.Data.(model.Desire)
	requestUesrID := c.MustGet("user_id").(int)

	var DesireInfo model.Desire
	var userDesire model.UserDesire
	userID := userDesire.GetUserIDbyWishID(desireJson.ID)

	if desireInfo.State == 0 {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "这个愿望还没被点亮", ""))
		return
	}

	if userID == requestUesrID {
		desireJson.State = 2
		res := desireModel.AchieveDesire(desireJson)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	DesireInfo.ID = desireInfo.ID
	DesireInfo.LightUser = requestUesrID

	// 实现light_user为自己的愿望
	if DesireModel.CheckUserLishtDesire(DesireInfo) == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "这个不是你点亮的愿望欧", ""))
		return
	}

	var userModel model.User

	// 是谁的愿望就给谁发邮件，所以要在user_desire表中找到愿望的的主人的ID，然后再找的他的邮箱
	userEmail := userModel.GetUserEmailByUserID(userID)

	userName := userModel.GetUserNameByUserID(desireInfo.LightUser)

	go func() {
		mailErr, err2 := helper.SendMail(userEmail, common.HaveAchieve, userName, "")
		log.Println(mailErr)
		if err2 != nil {
			log.Print(err2)
		}
	}()
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "实现愿望成功，我们已经通知该同学啦", ""))
}

// 获取点亮人的信息
func GetLigtherInfo(c *gin.Context) {
	var lighterInfo model.LightInfo
	// var lighterModel model.LightInfo
	lighterModel := model.LightInfo{}

	if err := c.ShouldBindQuery(&lighterInfo); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	res := lighterModel.GetLightInfoByDesireID(lighterInfo)
	if res.Status == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}
