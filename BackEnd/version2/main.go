package main

import (
	"context"
	"log"
	"os"

	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/handlers"
	"github.com/walle692/D0018E/BackEnd/version2/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin" // requires "go install github.com/gin-gonic/gin@latest" to be ran
)

func main() {
	// initialize pgpool
	ctx := context.Background()

	if err := global.InitPG(ctx, os.Getenv("DATABASE_URL")); err != nil {
		log.Fatalf("failed to init postgres: %v", err)
	}

	// initialize engine
	e := engine()

	if err := e.Run("0.0.0.0:5000"); err != nil {
		log.Fatal("unable to start")
	}
}

func engine() *gin.Engine {
	// new gin engine
	r := gin.New()

	// enable gin logging
	r.Use(gin.Logger())

	// setup cookie store this will automatically handle session cookie read/write
	r.Use(sessions.Sessions("session", cookie.NewStore([]byte(global.Secret))))

	// routes for logging in and out
	r.POST("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)

	// private route group, protected by authentication
	// all routes in this group require a valid session
	private := r.Group("/private")
	// enable the authenticate session for this group
	private.Use(handlers.AuthenticateSession)
	{
		private.GET("/me", services.Me)
		private.GET("/status", services.Status)
	}

	return r
}
