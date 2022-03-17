package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Product(c *gin.Context) {
	judul := c.Query("judul")
	harga := c.Query("harga")
	time := time.RFC850

	c.JSON(http.StatusOK, gin.H{
		"Message" : "item buku",
		"Item Name" : judul,
		"Time" : time,
		"Price" : harga,
	})
}

func IdProduct( c*gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"id" : id,
		"Message" : "item name",
		"Product name" : name,
		"Price" : price,
	})
}