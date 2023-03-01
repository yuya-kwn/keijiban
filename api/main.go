package main

import (
	"fmt"
	"keijiban/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	engine := gin.Default()
	db := database.DbConnect()
	db.AutoMigrate(&database.Comment{})

	engine.LoadHTMLGlob("view/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	engine.POST("/comment", func(c *gin.Context) {
		comment := database.Comment{}
		c.BindJSON(&comment)
		if err := db.Create(&comment).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusCreated, comment)
	})

	engine.GET("/showcomment", func(c *gin.Context) {
		comment := []database.Comment{}
		data := db.Find(&comment)
		if data.Error != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v" ,comment)
		c.HTML(http.StatusOK, "show.html", gin.H{"data": comment})
	})

	engine.Run(":3030")
}
