package helper

import "fmt"

func SayPagi(name string) string {
	fmt.Println("Selamat pagi ", name)
	pesan := "Selamat Pagi " + name
	return pesan
}
