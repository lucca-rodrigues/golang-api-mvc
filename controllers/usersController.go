package controllers

import (
	"fmt"
	"strconv"

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

func Create(c *gin.Context){
	db := database.GetDatabase()
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func Update(c *gin.Context){
	id := c.Params.ByName("id")
	db := database.GetDatabase()

	var user models.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
	}
	c.BindJSON(&user)

	db.Save(&user)
	c.JSON(200, user)
}

func Delete(c *gin.Context){
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.User{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}