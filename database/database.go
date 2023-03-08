package database

import (
	"fmt"
	"keijiban/crypto"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Comment struct {
	ID   uint   `json:"id"`
	Body string `json:"body"`
}

type User struct {
	gorm.Model
	UserId   string `json:"userid"`
	Password string `json:"password"`
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

func Signup(userId, password string) (*User, error) {
	db := DbConnect()
	user := User{}
	db.Where("user_id = ?", userId).First(&user)
	if user.ID != 0 {
		fmt.Println("同一名のUserIdが既に登録されています。")
		return nil, err
	}

	encryptPw, err := crypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}
	user = User{UserId: userId, Password: encryptPw}
	db.Create(&user)
	return &user, nil
}
