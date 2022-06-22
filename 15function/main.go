package main

import "fmt"

func main() {

	fmt.Println("\nWelcom to functions!!!")

	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	fmt.Println("ProAdder rult is: ", proadder(2, 3, 4, 5, 12131, 5))
}

func greeter() {
	fmt.Println("Namastey!")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proadder(vaalues ...int) int {
	total := 0

	for _, val := range vaalues {
		total += val
	}
	return total
}

//Annonomous Functions***
//Immidiately Invoked Functions***
