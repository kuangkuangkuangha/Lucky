package middleware

import (
	"lucky/app/common"
	"lucky/app/controller"
	"lucky/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeError, "token不存在", nil))
		}

		student_number, err := helper.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeError, "权限不足", nil))
			c.Abort()
			return
		}

		UserID := controller.GetUserIdFromDB(student_number)
		UserSchool := controller.GetUserSchoolFromDB(student_number)

		c.Set("student_number", student_number)
		c.Set("user_id", UserID)
		c.Set("school", UserSchool)
		c.Next()
	}
}
