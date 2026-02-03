package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
	"github.com/walle692/D0018E/BackEnd/version2/global"

)

// logout is a handler that makes the user logout
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user 	:= session.Get(global.Userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(global.Userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
