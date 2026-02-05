package services

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/global"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
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

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
	Role     string `json:"role"     binding:"required,oneof=buyer seller admin"`
}

func MakeUser(c *gin.Context) {
	// Read input from request body
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CreateUser(req.Username, req.Password, req.Role); err != nil {
		// optionally detect duplicate username, etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
