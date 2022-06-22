package main

import 	"fmt"

func main() {
	fmt.Println("Welcom to mymap!!!")
	languages := make(map[string]string)

	languages["JS"] = "JavaScipt"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("\nLists of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	
	delete(languages, "RB")
	
	fmt.Println("\nLists of all languages: ", languages)

	//loops are iteresting in golang
	for key, value := range languages{
		fmt.Printf("\nFor key %v, value is %v", key, value)
	}

	//coma ok in loop
	for _, value := range languages{
		fmt.Printf("\nValues are: %v", value)
	}
}
