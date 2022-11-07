package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type People struct {
	Name   string
	Age    int
	Status string
}

type Community struct {
	communityName string
	anggota       []People
}

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
func data(c *gin.Context) {
	// some1 := People{
	// 	Name:   "Lily Janvieka",
	// 	Age:    20,
	// 	Status: "Singgle",
	// }
	// some2 := People{
	// 	Name:   "Stephani Langi",
	// 	Age:    19,
	// 	Status: "Singgle",
	// }
	// array := map[string]string{
	// 	"nama":  "Dominikus Andika Kurniawan",
	// 	"age":   "20",
	// 	"NIM":   "2440056972",
	// 	"Major": "Sistem Informasi dan Manajemen",
	// }
	// array1 := map[int]People{
	// 	1: some1,
	// 	2: some2,
	// }
	// array2 := [...]People{
	// 	some1,
	// 	some2,
	// }
	himsisfo := Community{
		communityName: "Himsisfo",
		anggota: []People{
			People{Name: "Lily Janvieka",
				Age:    20,
				Status: "Singgle"},
			People{Name: "Stephani Langi",
				Age:    19,
				Status: "Singgle"},
		},
	}
	c.JSON(http.StatusOK, himsisfo)
}

func Api() {
	r := gin.Default()
	r.GET("/massage", massage)
	r.GET("/kegiatan/:nama/:aksi", action)
	r.GET("/data", data)
	r.Run()
}
