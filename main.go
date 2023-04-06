package main

import (
	"github.com/InYuusha/api"
	"github.com/gin-gonic/gin"
)

func main() {

	port := "5000"
	app := gin.Default()
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
