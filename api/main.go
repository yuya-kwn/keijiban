package main

import (
	"keijiban/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(";3030")
}
