package main

import (
	"net/http"
	"strconv"

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

func getProductById(ctx *gin.Context) {
	id := c.Params("id") //string
	idn, _ := strconv.ParseInt(id, 10, 0)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"id":   idn,
		"name": "Product A",
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
	var r1 = r.Group("/api")
	{
		r1.GET("/hello", getHello)
		r1.POST("/hello", postHello)
		r1.PUT("/hello", putHello)
		r1.DELETE("/hello", deleteHello)
	}

	//handling path Params
	r.GET("/product/:id", getProductById)
	r.GET("/profile/:username", showProfile)
	r.GET("/compute/:num1/add/:num2", compute)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
