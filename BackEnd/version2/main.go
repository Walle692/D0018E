package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/handlers"
	"github.com/walle692/D0018E/BackEnd/version2/services/basket_services"
	"github.com/walle692/D0018E/BackEnd/version2/services/order_services"
	"github.com/walle692/D0018E/BackEnd/version2/services/products_services"
	"github.com/walle692/D0018E/BackEnd/version2/services/user_services"

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
		private.GET("/me", user_services.Me)
		private.GET("/status", user_services.Status)
		private.GET("/products", products_services.ListAll)
		private.GET("/products/:id", products_services.GetByProductID)
		private.POST("/basket/add", basket_services.AddToBasket)
		private.DELETE("/basket/delete", basket_services.DeleteFromBasket)
		private.GET("/basket", basket_services.GetBasket)
		private.GET("/orders", order_services.GetUserOrders)
		private.POST("/checkout", basket_services.CheckOut)
		private.POST("/search/products", products_services.SearchProduct)
	}
	seller := r.Group("/seller")
	seller.Use(handlers.AuthenticateSession, handlers.AuthenticateSellerSession)
	{
		seller.POST("/create-product", products_services.CreateProduct)
		seller.GET("/products", products_services.GetSellerProduct)
		seller.DELETE("/products/delist", products_services.Delist)
		seller.POST("/products/update", products_services.Update)
		seller.POST("/orders", order_services.SearchOrdersLimited)
	}

	admin := r.Group("/admin")
	admin.Use(handlers.AuthenticateSession, handlers.AuthenticateAdminSession)
	{
		admin.POST("/create-user", user_services.MakeUser)
		admin.POST("/orders", order_services.SearchOrders)
	}

	return r
}
