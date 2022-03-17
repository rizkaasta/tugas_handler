package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Product(c *gin.Context) {
	judul := c.Query("judul")
	harga := c.Query("harga")
	time := time.Now()

	c.JSON(http.StatusOK, gin.H{
		"Message" : "item buku",
		"Item Name" : judul,
		"Time" : time,
		"Price" : harga,
	})
}