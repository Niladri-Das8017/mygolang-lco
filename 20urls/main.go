package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jbdaj151af"

func main() {
	fmt.Println("\nWelcome to handling urls!!!")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) //it's not a proerty, it's a method
	fmt.Println(result.RawQuery)

	// Yes, we can store entire RawQueery in a variable, but, there is a better way to do it.

	qparams := result.Query()

	fmt.Printf("\nThe type of query params are: %T\n", qparams) //it will result url.Values, a fancy name for key value pairs. So we can play with it using dcstionary as it contains keys and values

	fmt.Println(qparams["coursename"]) //used the key to fetch values.

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//***Creating url from chunks***

	partsOfUrl := &url.URL{ //Remember, we will be always passing on referance in this case, not the copy of data.

		//There is special parameters we have to define, we cannot name them anything.

		Scheme:  "https", // we don't need too put colon or double-slash
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	//Now we can coonstruct an url out of it
	anotherURL := partsOfUrl.String() //convertting it into string

	fmt.Println(anotherURL)

}
