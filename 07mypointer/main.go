package main

import "fmt"

func main() {

	fmt.Println("Welcome Mypointer!!!")
	//var ptr *int
	//fmt.Println("The initial value of ptr is: ", ptr)

	myNumber := 23 
	var ptr = &myNumber

	fmt.Println("Value of ptr: ", ptr)
	fmt.Println("Value of *ptr: ", *ptr)
	fmt.Println("Value of *ptr: ", *ptr+2)
}