package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mmarci96/go-url-shortener/handler"
	"github.com/mmarci96/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})
	r.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
