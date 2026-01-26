package handlers

import (
	"net/http"

	"github.com/Walle692/D0018E/tree/main/BackEnd/version1/services"

	"github.com/gin-gonic/gin"
)

// struct for the given post data
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	// create request variable
	var req LoginRequest

	// if the json is faulty throw error
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// use the authenticate user, pass username and password on to it
	token, err := services.AuthenticateUser(req.Username, req.Password)

	// error is thrown if the credentials are faulty
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// sets a cookie with token for the user
	c.SetCookie("token", token, 3600, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
