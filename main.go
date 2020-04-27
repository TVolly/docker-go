package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello all")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Index Page")
	})

	http.ListenAndServe(":8080", nil)
}
