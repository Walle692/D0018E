package handlers

import (

)

func AuthenticateSession(c *gin.Context) {
	// get the session from the request context
	session := sessions.Default(c)

	if user := session.Get(user)
}