package services

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

// returns the username
func Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(global.Userkey)
	role := session.Get(global.Role)
	c.JSON(http.StatusOK, gin.H{"user": user, "role": role})
}

// status will tell the user whether it is logged in or not.
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
