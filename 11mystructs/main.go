package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nStructs in Golang!!!")
	niladri := User{"Niladri", "niladr@go.dev", true, 23}
	fmt.Println(niladri)
	fmt.Println(niladri.Name)
	fmt.Printf("Nladri's details are: %+v", niladri)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
