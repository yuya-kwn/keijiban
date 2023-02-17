package main

import (
	"keijiban/controller"
	"keijiban/model"
)

func main() {
	router := controller.GetRouter()
	model.DbConnect()
	router.Run(":3030")
}
