package main

import "fmt"

func main() {
	var username string = "Niladri"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T\n\n", isLoggedIn)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Variable is of type : %T\n\n", smallint)

	var smallFloat float64 = 255.626611451311685151665
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T\n\n", smallFloat)

	//default values and some aliases

	var someVar int
	fmt.Println(someVar)
	fmt.Printf("Variable is of type : %T\n\n", someVar)

	//implicit type
	var website = "niladridasofficial.42web.io"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n\n", website)

	//no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
