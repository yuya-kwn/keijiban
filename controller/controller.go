package controller

import (
	"fmt"
	"keijiban/database"
	"keijiban/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTop(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func PostComment(c *gin.Context) {
	db := database.DbConnect()
	db.AutoMigrate(&database.Comment{})
	comment := database.Comment{}
	c.BindJSON(&comment)
	if err := db.Create(&comment).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusCreated, comment)
}

func GetComment(c *gin.Context) {
	db := database.DbConnect()
	db.AutoMigrate(&database.Comment{})
	comment := []database.Comment{}
		data := db.Find(&comment)
		if data.Error != nil {
			fmt.Println("error")
		}
		fmt.Printf("%v", comment)
		c.HTML(http.StatusOK, "show.html", gin.H{"data": comment})
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := database.DB.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
