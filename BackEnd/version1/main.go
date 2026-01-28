package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Walle692/D0018E/BackEnd/version1/handlers"
	"github.com/Walle692/D0018E/BackEnd/version1/utils"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin" // requires "go install github.com/gin-gonic/gin@latest" to be ran
)

func setupRouter(pg *utils.Postgres) *gin.Engine {

	// this creates a router
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler(pg))

	return r
}

func main() {
	godotenv.Load()

	pool, err := utils.NewPG(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("error setting up pool")
		return
	}
	defer pool.Close()

	// this creates a router
	router := setupRouter(pool)

	router.Run(":5000")

}
