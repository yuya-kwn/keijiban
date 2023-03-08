package main

import (
	"keijiban/controller"
	"keijiban/database"
)

func main() {
	database.DbConnect()
	router := controller.GetRouter()
	router.Run(":3030")
}
