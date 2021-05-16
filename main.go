package main

import(
	"fmt"
	"net/http"

)

func main() {
	http.HandleFunc("/", http.NotFoundHandler().ServeHTTP

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}