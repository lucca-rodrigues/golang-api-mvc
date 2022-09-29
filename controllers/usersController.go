package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luccarodrigues/golang-api-mvc/database"
	"github.com/luccarodrigues/golang-api-mvc/models"
)


func GetAll(c *gin.Context){
	db := database.GetDatabase()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}