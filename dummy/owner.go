package dummy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Owner struct {
	OwnerName string
	Ig        string
	Image     string
}

func Data(c *gin.Context) {

	array2 := [...]Owner{
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

	c.Writer.Header().Set("Content-Type", "application/json")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"array": array2,
	})
}

func Dum() {
	fmt.Println("wo")
}
