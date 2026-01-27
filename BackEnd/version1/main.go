package main

import (
	"github.com/Walle692/D0018E/tree/main/BackEnd/version1/handlers"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin" // requires "go install github.com/gin-gonic/gin@latest" to be ran
)

func setupRouter() *gin.Engine {

	// this creates a router
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)

	return r
}

func main() {
	godotenv.Load()

	// this creates a router
	router := setupRouter()

	router.Run(":5000")

}
