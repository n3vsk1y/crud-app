package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n3vsk1y/crud-app/routes"
)

func main() {
	r := gin.Default()
	routes.RegRoutes(r)
	r.Run(":8080")
}
