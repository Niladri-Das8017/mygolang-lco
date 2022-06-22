package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("\nWelcome to mymaths!!!")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	//fmt.Println("The sum is: ", mynumberOne + int(mynumberTwo))  ****It will parsee the float into int, so imagine it's 4.5 cr then 0.5 cr will just washed up & we don't want that.****

	//random number
	//***There is nothing random in the computer

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) //rand.int returns an integer, that's why we don't have to convert it before any increment.

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNum)

}
