package handlers

import (
	"net/http"

	albums "github.com/MaximTretjakov/Go-tutorial/web-service-gin/data"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.Albums)
}
