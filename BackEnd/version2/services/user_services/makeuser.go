package user_services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walle692/D0018E/BackEnd/version2/utils/user"
)

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

	if err := user.CreateUser(req.Username, req.Password, req.Role); err != nil {
		// optionally detect duplicate username, etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
