package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST API - GET REQUEST",
	})
}

func postHello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST API - POST REQUEST",
	})
}

func putHello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST API - PUT REQUEST",
	})
}

func deleteHello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST API - DELETE  REQUEST",
	})
}

func main() {
	r := gin.Default()

	//GET POST PUT DELETE
	r.GET("/", getHello)
	r.POST("/", postHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)

	//grouping
	r1 := r.Group("/api"){
		r1.GET("/hello", getHello)
		r1.POST("/hello", postHello)
		r1.PUT("/hello", putHello)
		r1.DELETE("/hello", deleteHello)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
