package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niladridas/mongoAPI/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4040", r))
	fmt.Println("Listening at port 4040...")
}
