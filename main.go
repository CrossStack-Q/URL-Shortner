package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AuraReaper/go-url-shortner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load() ; err != nil {
		fmt.Println(err)
	}

	app := gin.Default()

	setUpRouters(app)

	port := os.Getenv("APP_PORT");
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Run(":"+port))
}

func setUpRouters(app *gin.Engine) {

	app.POST("/api/v1" , routes.ShortenURL)
	app.GET("/api/v1/:shortID" , routes.GetByShortID)
	app.POST("/api/v1/:shortID" , routes.EditURL)
	app.DELETE("/api/v1/:shortID" , routes.DeleteURL)
	app.POST("/api/v1/addTag" , routes.AddTag)
}