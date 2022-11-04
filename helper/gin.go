package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func massage(c *gin.Context) {
	test := SayPagi("Andika")
	c.JSON(http.StatusOK, gin.H{
		"message": test,
	})
}

func action(c *gin.Context) {
	name := c.Param("nama")
	aksi := c.Param("aksi")
	combin := name + " sedang " + aksi
	c.JSON(http.StatusOK, combin)
}

func Api() {
	r := gin.Default()
	r.GET("/massage", massage)
	r.GET("/kegiatan/:nama/:aksi", action)
	r.Run()
}
