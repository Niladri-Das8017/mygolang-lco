package controller

import (
	"net/http"
)

//HomePage
func ServeHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to API by Niladri Das</h1>"))
}
