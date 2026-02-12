package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/handlers"
	"github.com/walle692/D0018E/BackEnd/version2/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin" // requires "go install github.com/gin-gonic/gin@latest" to be ran
	"github.com/joho/godotenv"

	// go get github.com/gin-contrib/cors
	"github.com/gin-contrib/cors"
)

func main() {
	// initialize godotenv to be able to use env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	// ✅ CORS (put this BEFORE routes)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5001",    //For local tests
			"http://13.60.56.204:5001", // This should be the frontend IP and port when hosted
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // IMPORTANT for cookies/sessions
		MaxAge:           12 * time.Hour,
	}))

	// new mr chat code to fix cookie issue
	store := cookie.NewStore([]byte(global.Secret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false, // set true only when HTTPS
	})
	r.Use(sessions.Sessions("session", store))

	// setup cookie store this will automatically handle session cookie read/write
	//r.Use(sessions.Sessions("session", cookie.NewStore([]byte(global.Secret))))

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
		private.GET("/products", services.GetProducts)
	}
	admin := r.Group("/admin")
	admin.Use(handlers.AuthenticateSession, handlers.AuthenticateAdminSession)
	{
		admin.POST("/create-user", services.MakeUser)
	}

	return r
}
