package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func AuthenticateSession(c *gin.Context) {
	// get the session from the request context
	session := sessions.Default(c)

	// try to get the user from the session
	if user := session.Get(global.Userkey); user == nil {
		// no user in session, abort the request
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// user is authenticated, continue to the next handler
	c.Next()
}

func AuthenticateAdminSession(c *gin.Context) {
	sessions := sessions.Default(c)

	role := sessions.Get(global.Role)
	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func AuthenticateSellerSession(c *gin.Context) {
	sessions := sessions.Default(c)

	role := sessions.Get(global.Role)
	if role != "seller" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
