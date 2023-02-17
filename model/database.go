package model

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Title string `form:"tilte"`
	Body  string `form:"body"`
}

var DB *gorm.DB

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
