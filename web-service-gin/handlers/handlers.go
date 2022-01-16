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

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
