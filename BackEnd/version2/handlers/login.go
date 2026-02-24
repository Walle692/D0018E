package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils/user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// login is a handler that parses a form and checks for specific data.
func Login(c *gin.Context) {
	session := sessions.Default(c)

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	}

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)

	// validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// check for username and password match, usually from a database
	dbpassword, role, userID, err := user.GetPassword(username)
	// if there is error or passwords do not match failure
	if err != nil || dbpassword != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	//set global role

	// save the username in the session
	session.Set(global.Userkey, username)
	session.Set(global.Role, role)
	// Store int as an string for edge cases
	session.Set(global.UserID, strconv.Itoa(userID))
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}
