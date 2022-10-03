package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default();
	router.GET("/", helloWorld)
	
	router.Run("localhost:8080")
}

func helloWorld (ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Hello World")
}