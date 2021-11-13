package routes

import (
	Controller "lucky/app/controller"
	"lucky/app/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("api/whutlogin", Controller.WhutLogin)
	router.POST("api/login/whut/callback", Controller.Callback)
	router.POST("api/ccnulogin", Controller.CcnuLogin)

	api := router.Group("api/")

	api.Use(middleware.AuthMiddleware())
	{
		//api.POST("WhutLogin", Controller.WhutLogin)
		api.POST("/user/email/check", Controller.CheckUserEmail)
		api.POST("/user/email", Controller.BindEmail)
		api.GET("/user/info/lightman", Controller.GetLigtherInfo)

		wishes := api.Group("/wishes")
		{
			wishes.POST("/add", Controller.AddDesire)
			wishes.POST("/light", Controller.LightDesire)
			wishes.POST("/achieve", Controller.AchieveDesire)
			// wishes.GET("", Controller.GetAllDesire)
			wishes.GET("/user/post", Controller.GetUserPostDesire)
			wishes.GET("/user/light", Controller.GetUserLightDesire)
			wishes.GET("/details", Controller.GetWishDetails)
			wishes.GET("/categories", Controller.GetWishByCatagories)
			wishes.DELETE("", Controller.DeleteWish) // 删除愿望池中的愿望
			// wishes.POST("/10", Controller.Get10Wishes)
			wishes.POST("/giveup", Controller.CancelLightDesire)
		}

		// message := api.Group("/message")
		// {
		// 	message.POST("/leave", Controller.LeaveMessage) // 弃用
		// 	message.GET("/get", Controller.GetUserMessage)  // 弃用
		// }

		api.GET("")
	}

}
