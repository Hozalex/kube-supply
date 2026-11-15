package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler() http.Handler {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	return r
}
