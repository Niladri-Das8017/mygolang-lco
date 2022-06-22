package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome Brother!!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	//coma Ok || err ok Syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thaks for rating: ", input)

	fmt.Printf("Type of this rating is: %T", input)
}
