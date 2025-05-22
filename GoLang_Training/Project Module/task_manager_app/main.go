package main

import (
	"task_manager_api/db"
	"task_manager_api/router"
)

func main() {
	config.Connect()

	r := router.SetupRouter()
	r.Run(":8080") // run on localhost:8080
}
