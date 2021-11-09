package routes

import (
	Controller "lucky/app/controller"
	"lucky/app/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("api/whutLogin", Controller.WhutLogin)
	router.POST("api/login/whut/callback", Controller.Callback)
	router.POST("api/ccnulogin", Controller.CcnuLogin)

	api := router.Group("api/")

	api.Use(middleware.AuthMiddleware())
	{
		//api.POST("WhutLogin", Controller.WhutLogin)
		api.POST("/user/email", Controller.BindEmail)

		wishes := api.Group("/wishes")
		{
			wishes.POST("/add", Controller.AddDesire)
			wishes.POST("/light", Controller.LightDesire)
			wishes.POST("/achieve", Controller.AchieveDesire)
			wishes.GET("", Controller.GetAllDesire)
			wishes.GET("/user", Controller.GetUserDesire)
			wishes.GET("/user2", Controller.GetUserDesire2)
			wishes.GET("/details", Controller.GetWishDetails)
			wishes.GET("/categories", Controller.GetWishByCatagories)
			wishes.DELETE("", Controller.DeleteWish) // 删除愿望池中的愿望
			wishes.POST("/10", Controller.Get10Wishes)
			wishes.POST("/giveup", Controller.CancelLightDesire)
		}

		message := api.Group("/message")
		{
			message.POST("/leave", Controller.LeaveMessage)
			message.GET("/get", Controller.GetUserMessage)
		}

		api.GET("")
	}

}
