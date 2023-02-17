package controller

import(
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", ShowAllModel)
	r.GET("/show/:id", ShowOneModel)
	r.GET("/create", ShowCreataModel)
	r.POST("/create", CreateModel)
	r.GET("/edit/:id",ShowEditModel)
	r.POST("/edit", EditModel)
	r.GET("/delete/:id", ShowDeleteModel)
	r.POST("/delete", DeleteModel)
	
	return r
}