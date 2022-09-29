package controllers

import "github.com/gin-gonic/gin"


func GetAll(c *gin.Context){
	p := "Batata"

	c.JSON(200, p)
}