package main

import (
	"fmt"
	"go_new/helper"
)

func tryRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * tryRecursive(value-1)
	}
}

type People struct {
	Name   string
	Age    int
	Status string
}
type Owner struct {
	ownerName string
	ig        string
	image     string
}

func (orang People) sayHello(name string) {
	fmt.Println("Selamat pagi", name, "salam dari ", orang.Name)
}

func main() {
	fmt.Println("heloo")
	var array = [...]string{
		"satu",
		"dua",
		"tiga",
		"empat",
	}
	slice := array[1:2]

	array2 := []Owner{
		{
			ownerName: "Dominikus Andika Kurniawan",
			ig:        "dominikusandika",
			image:     "",
		},
		{
			ownerName: "Dominikus Andika Kurniawan",
			ig:        "dominikusandika",
			image:     "",
		},
		{
			ownerName: "Dominikus Andika Kurniawan",
			ig:        "dominikusandika",
			image:     "",
		},
		{
			ownerName: "Dominikus Andika Kurniawan",
			ig:        "dominikusandika",
			image:     "",
		},
	}

	camp := map[int]string{
		1: "kamu",
		2: "aku",
	}
	fmt.Println(slice)
	fmt.Println(camp[1])

	fmt.Println(tryRecursive(2))
	andika := People{
		Name:   "Dominikus",
		Age:    20,
		Status: "Single",
	}
	fmt.Println(array2)

	andika.sayHello("ivan")
	sa := SaySorry("devona")
	fmt.Println(sa)
	helper.SayPagi("Andika")
	helper.Api()
}
