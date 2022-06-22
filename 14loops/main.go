package main

import "fmt"

func main() {
	fmt.Println("\nWelcom to loops in golang!!!")

	//slice
	days := []string{"Sunday", "Monay", "Wensday", "Frday", "Saturday"}

//for loop type 1
	// for d:=0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

//for loop type 2
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

//for loop type 3
		for index, day := range days{
			fmt.Println(index, day)
		}

		roughValue := 1

		for roughValue < 10 {

			if roughValue == 2 {
				goto lco
			}

			if roughValue == 5 {
				//break
				roughValue++
				continue 	//continue will skip 5, hence we addeed 1 to it.
			}
			fmt.Println("Value is: ", roughValue)
			roughValue++
		}

		lco:
			fmt.Println("Jumping at niladridasofficial.42web.io")

}