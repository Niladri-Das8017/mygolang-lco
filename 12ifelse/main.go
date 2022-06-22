package main

import "fmt"

func main() {
	fmt.Println("\nIf else in GoLang!!!")

	lognCount := 23

	var result string
	if lognCount < 10 {

		result = "Regulaa User"

	} else if lognCount > 10 {

		result = "New User"

	} else {
		result = "Result is exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Numbre is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is nott less than 10")
	}

	// if err != nil {
		
	// }

}
