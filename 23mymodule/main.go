package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("\nHello mod!!! ")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/print", print)

	log.Fatal(http.ListenAndServe(":4020", r)) //after execution check port 4000
}

func greeter() {
	fmt.Println("hey there mod users")
}

// when somebody sending us a request and if we want to use parameters, URLs and all of that, that all is inside the "r"; but if we want to send some response back,
// that is going to be done through response-writer "w"
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on Niladri's PC</h1>"))
	return
}
func print(w http.ResponseWriter, r *http.Request) {
	datbyte, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	content := string(datbyte)
	fmt.Println(content)
	json.NewEncoder(w).Encode(content)
	return
}
