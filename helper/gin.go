package helper

import (
	"go_new/dummy"
	"net/http"

	"github.com/gin-gonic/gin"
)

type People struct {
	Name   string
	Age    int
	Status string
}
type URI struct {
	Details string
}

type Owner struct {
	OwnerName string
	Ig        string
	Image     string
}

// type Community struct {
// 	communityName string
// 	anggota       []People
// }

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
	som3 := Owner{
		OwnerName: "Dominikus Andika Kurniawan",
		Ig:        "dominikusandika",
		Image:     "",
	}
	// details := URI{
	// 	Details: "heloo",
	// }
	array2 := [...]Owner{
		som3,
		{
			OwnerName: "Dominikus Andika Kurniawan",
			Ig:        "dominikusandika",
			Image:     "",
		},
		{
			OwnerName: "Dominikus Andika Kurniawan",
			Ig:        "dominikusandika",
			Image:     "",
		},
		{
			OwnerName: "Dominikus Andika Kurniawan",
			Ig:        "dominikusandika",
			Image:     "",
		},
		{
			OwnerName: "Dominikus Andika Kurniawan",
			Ig:        "dominikusandika",
			Image:     "",
		},
	}
	// himsisfo := Community{
	// 	communityName: "Himsisfo",
	// 	anggota: []People{
	// 		{
	// 			Name:   "Lily Janvieka",
	// 			Age:    20,
	// 			Status: "Singgle",
	// 		},
	// 		{
	// 			Name:   "Lily Janvieka",
	// 			Age:    20,
	// 			Status: "Singgle",
	// 		},
	// 	},
	// }
	// fmt.Println(array2)
	dummy.Dum()
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"data":  som3,
		"array": array2,
	})
}

func Api() {
	r := gin.Default()
	r.GET("/massage", massage)
	r.GET("/kegiatan/:nama/:aksi", action)
	r.GET("/data", dummy.Data)
	r.Run()
}
