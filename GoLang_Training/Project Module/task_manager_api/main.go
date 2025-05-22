package main

import (
	"fmt"
	"net/http"
	"task_manager_api/config"
	"task_manager_api/handler"
	"task_manager_api/repository"
	router "task_manager_api/routes"
	"task_manager_api/services"
	"time"
)

func main() {

	db := config.DBConnection()

	tmrep := repository.NewTaskManagerRepository(db)

	taskservice := services.NewTaskService(*tmrep)

	th := handler.NewTaskHandler(taskservice)

	routes := router.NewRouter(th)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(err)
	}

}
