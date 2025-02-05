package main

import (
	"go-todo-app/config"
	"go-todo-app/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
