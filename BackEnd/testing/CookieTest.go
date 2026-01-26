package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // requires "go install github.com/gin-gonic/gin@latest" to be ran
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get cookie
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}
		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	// this creates a router
	router := gin.Default()

	// defines a route for the router when one visits this they get a cookie
	router.GET("/login", func(c *gin.Context) {
		// set cookie {"label": "ok"} with max age 30 sec
		c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
		c.String(200, "login success!")
	})

	// when trying to visit home if there is no cookie then the perosn gets a error message
	router.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "your home page!"})
	})

	router.Run(":5000")

}
