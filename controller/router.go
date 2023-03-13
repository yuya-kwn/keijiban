package controller

import (
	"keijiban/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.GET("/", GetTop)
	r.POST("/comment", PostComment)
	r.GET("/showcomment", GetComment)
	r.GET("/signup", GetSignup)
	api := r.Group("/api")
	{
		api.POST("/token", GenerateToken)
		api.POST("/signup", RegisterUser)
	}
	secured := api.Group("/secured").Use(middleware.Auth())
	{
		secured.GET("/ping", Ping)
	}
	return r
}
