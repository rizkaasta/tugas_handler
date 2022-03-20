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
		"Message":   "item buku",
		"Item Name": judul,
		"Time":      time,
		"Price":     harga,
	})
}

func IdProduct(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"Message":      "item name",
		"Product name": name,
		"Price":        price,
	})
}

func SignUp(c *gin.Context) {
	nama := c.Query("nama")
	time := time.RFC850

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Selamat bergabung bersama kami",
		"Nama" : nama,
		"Wakru Daftar" : time,
	})
}

func SignIn(c *gin.Context) {
	nama := c.Query("nama")
	time := time.RFC850

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Selamat datang kembali",
		"Nama" : nama,
		"Waktu login" : time,
	})
}
 func Booking(c *gin.Context) {
	buku := c.Query("judul")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Buku berhasil dipinjam",
		"Judul Buku" : buku,
	})
 }

 func Return(c *gin.Context) {
	buku := c.Query("judul")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Buku berhasil dikembalikan",
		"Judul Buku" : buku,
	})
 }

 func SignOut(c *gin.Context) {
	nama := c.Query("nama")
	time := time.RFC850

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Terimakasih dan samapai berjumpa kembali",
		"Nama" : nama,
		"Waktu logout" : time,
	})
 }

