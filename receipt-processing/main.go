package main

import (
	"receipt-processing/routes"
)

func main(){
	router := routes.SetupRouter()
	router.Run(":8080")
}