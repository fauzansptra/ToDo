package main

import (
	"github.com/fauzansptra/go-rest-api/config"
    "github.com/fauzansptra/go-rest-api/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
