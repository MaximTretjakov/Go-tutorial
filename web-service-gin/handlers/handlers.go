package handlers

import (
	"net/http"

	albums "github.com/MaximTretjakov/Go-tutorial/web-service-gin/data"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.Albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum albums.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums.Albums = append(albums.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
