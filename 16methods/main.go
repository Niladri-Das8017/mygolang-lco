package main

import "fmt"

func main() {

	fmt.Println("\nStructs in Golang!!!")
	niladri := User{"Niladri", "niladr@go.dev", true, 23}
	fmt.Println(niladri)
	fmt.Println(niladri.Name)
	fmt.Printf("Nladri's details are: %+v\n", niladri)

	niladri.GetStatus()
	niladri.NewMail()
	fmt.Println("Orginal email: ", niladri.Email)	//we can see that original email is not manipulated. When we pass objects or strcts through such function or methods, it actually passes on a copy 
	//thus we need pointers to manipulate that.
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//meethods
func (u User) GetStatus() {
	fmt.Println("Is usr active? ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this uer is: ", u.Email)
}
