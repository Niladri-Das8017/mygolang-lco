package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("\nSwitch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1    //if it gets 0 iit will bee ncrmented by 1, hence, "+1"

	fmt.Println("Valu of dice is ", diceNumber)



	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 & you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		//Go neeeds explicitly menion fallthrough 
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot & roll the dice again")

	//Not needed here though....
	default:
		fmt.Println("what was that!!!")
	}

}
