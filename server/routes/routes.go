package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luccarodrigues/golang-api-mvc/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.GetAll)
			users.POST("/", controllers.Create)
		}
	}
	return router
}