package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {

	dsn := "user:pass@tcp(db:3306)/keijiban?charset=utf8mb4"
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
