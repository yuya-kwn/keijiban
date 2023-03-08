package controller

import (
	"fmt"
	"keijiban/database"
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

func PostSignup(c *gin.Context) {
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := database.Signup(id, pw)
	if err != nil {
		c.Redirect(301, "/signup")
		return
	}
	c.HTML(http.StatusCreated, "index.html", gin.H{"user": user})
}
