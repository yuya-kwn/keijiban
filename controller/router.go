package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.GET("/", GetTop)
	r.POST("/comment", PostComment)
	r.GET("/showcomment", GetComment)
	r.GET("/signup", GetSignup)
	r.POST("/signup", PostSignup)
	return r
}