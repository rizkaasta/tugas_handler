package handler

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//no 1
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

//no 2
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

 //no 3
 func Anggota(c *gin.Context) {
	noAnggota := c.Param("no")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Nomor Keanggotaan",
		"No Anggota" : noAnggota,
	})
 }

 func Buku(c *gin.Context) {
	 noBuku := c.Param("no") 

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Nomor Buku",
		"Nomor Buku" : noBuku,
	})
 }

 func Jurnal(c *gin.Context) {
	noJurnal := c.Param("no")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Nomor Jurnal",
		"Nomor Jurnal" : noJurnal,
	})
 }

 func Pengarang(c *gin.Context) {
	keyword := c.Param("pengarang")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Keyword Pengarang",
		"Keyword" : keyword,
	})
 }

 func Judul(c *gin.Context) {
	keyword := c.Param("judul")

	c.JSON(http.StatusOK, gin.H{
		"Message" : "Keyword Judul",
		"Keyword" : keyword,
	})
 }

 //POST
type MahasiswaInput struct {
	Nama string `json:"Nama_lengkap" binding:"required"`
	NIM string `json:"Nomor_induk" binding:"required"`
	Fakultas string `json:"Fakultas" binding:"required"`
	Prodi string `json:"Program_studi" binding:"required"`
	TahunMasuk string `json:"Tahun_masuk" binding:"required"`
}

 func DataMahasiswa(c *gin.Context) {
	var mahasiswaInput MahasiswaInput

	err := c.ShouldBindJSON(&mahasiswaInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Berhasil input data",
		"Time" : time.RFC850,
		"Nama Mahasiswa" : mahasiswaInput.Nama,
		"NIM" : mahasiswaInput.NIM,
		"Fakultas" : mahasiswaInput.Fakultas,
		"Prodi" : mahasiswaInput.Prodi,
		"Tahun Masuk" : mahasiswaInput.TahunMasuk,
	})
 }