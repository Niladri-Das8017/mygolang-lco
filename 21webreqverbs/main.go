package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("\nWelcome to web verb vdeo - lco!!!")

	// PerformGetRequest() //Functoin Call Get
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // it will close the response after doing job

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Conttent length is: ", response.ContentLength)

	//This is a mandetory part we need to print content, don't matter whch way we print it.
	content, _ := ioutil.ReadAll(response.Body)

	//way 1 to print content

	fmt.Println(string(content))

	//way 2 to print content using "strings" package
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content) //### Solved: Here's the initalization ***

	fmt.Println("ByteCount is: ", byteCount)

	fmt.Println("Response: ", responseString.String())

}

func PerformPostJsonRequest() {
	fmt.Println("post")
	const myURL = "http://localhost:8000/post"

	//fakee json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang",
			"price" : 0,
			"platform" : "youtube"
		}
	`)

	//send post request
	response, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myURL = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "niladri")
	data.Add("lastname", "das")
	data.Add("email", "niladri@go.dev")

	//sending reequest
	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
