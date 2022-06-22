package main

import (
	"fmt"
	"net/http"
	"sync"
)

var siganls = []string{"test"}

//Waight group variable
var wg sync.WaitGroup //pointer
//Mutex
var mut sync.Mutex //pointer

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(siganls)
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {

		fmt.Println("OOPS in endpoint")

	} else {

		mut.Lock()//Lock
		siganls = append(siganls, endpoint)
		mut.Unlock()//Unlock

		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)


	}

}
