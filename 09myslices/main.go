package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome slices!!!!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist : %T\n", fruitList)

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	//By default it's an array with a  layer of abtruction. Hence, slices.
	highScores := make([]int, 4) //made a slice of type int with size 4

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	//************ ^
	//In this case we can see we used make() to create default memory allocation, just how much we neded.
	// But when we needed more memory we used append() methode & reallocate memory for the segment. It saved lot of memory and oped out memory wastage.
	//************

	//Sort package

	fmt.Println("\n")
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println("\n")
	//Remove a value from slic baseed on index *******

	var courses = []string{"ReactJs", "JavaScript", "Swift", "Python", "Ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
