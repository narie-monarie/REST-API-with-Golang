package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Rest")
}

func main() {
	r := gin.Default()
	r.GET("/hello", hello)

	r.Run(":8080")
	fmt.Println("The Server is running....")
}
