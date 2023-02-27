package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Comment struct {
	ID    uint	`json:"id"`
	Body  string `json:"body"`
}

var DB *gorm.DB
var err error

func DbConnect() *gorm.DB {
	dsn := "user:pass@tcp(mysql:3306)/entity?charset=utf8mb4"

	count := 0
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		for {
			if err == nil {
				break
			}
			fmt.Println("データベース再接続...")
			time.Sleep(time.Second)
			count++
			if count > 60 {
				fmt.Println("データベース接続失敗")
				log.Fatal(err)
			}
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		}
	}
	fmt.Println("データベース接続成功")

	return db
}

func main() {
	engine := gin.Default()
	db := DbConnect()
	db.AutoMigrate(&Comment{})

	engine.LoadHTMLGlob("view/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	engine.POST("/comment", func(c *gin.Context) {
		comment := Comment{}
		c.BindJSON(&comment)
		if err := db.Create(&comment).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusCreated, comment)
	})

	engine.GET("/showcomment", func(c *gin.Context) {
		comment := Comment{}
		result := db.Find(&comment)
		if result.Error != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, comment)
	})

	engine.Run(":3030")
}
