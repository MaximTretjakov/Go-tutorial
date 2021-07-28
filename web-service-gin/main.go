package main

import (
	"github.com/MaximTretjakov/Go-tutorial/web-service-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.Run("localhost:8080")
}
