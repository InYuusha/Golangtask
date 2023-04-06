package main

import (
	"log"
	"os"

	"github.com/InYuusha/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app := gin.Default()
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
