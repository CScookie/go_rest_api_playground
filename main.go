package main

import (
	"github.com/gin-gonic/gin"

	"github.com/go_rest_api_playground/database"
	"github.com/go_rest_api_playground/routes"

	_ "github.com/go_rest_api_playground/docs"
)

func main() {

	r := gin.Default()

	database.ConnectDatabase()

	routes.Routes(r)

	r.Run(":8080") // listen and serve on "localhost:8080"
}
