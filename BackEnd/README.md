# Quick Start

Navigate into BackEnd/version2 run

```sh
go run .
```

Ensure the .evn DATEBASE_URL matches with correctly with the build enviorment

# Backend file structure

## main.go
In main.go API endpoints and methods are declared with path and function that is run

```go
r.POST("/login", handlers.Login)
```

## services/
In services/ the gin.Context is attached to functions and the session is attached boilerplate example where userID is passed as an argument in:

```go
func GetSellerProduct(c *gin.Context) {
	session := sessions.Default(c)
	userIDStr := session.Get(global.UserID)
	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	RETURNVALUE, err := UTILSPACKET.FUNCTON(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, RETURNVALUE)
}
```

## utils/

In utils/ functions that do touch the database with queries exits, functions here should optimaly be indepent on the gin framework. Functions can/are intergration tested in this directory under tests/ 






