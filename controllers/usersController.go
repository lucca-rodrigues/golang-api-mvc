package userscontroller

import "github.com/gin-gonic/gin"


func get(c *gin.Context){
	p := "Batata"

	c.JSON(200, p)
}