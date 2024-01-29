package main

import (
	"io"
	"log"
	"os"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/middleware"
	"github.com/Pugpaprika21/go-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	db.Migrate()

	gin.DisableConsoleColor()
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())

	routes.Setup(router)

	router.Run(":8000")
}
