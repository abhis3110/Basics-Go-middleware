package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Inside the handler")
	w.Write([]byte("OK - executing handler"))
}

func main() {
	// to create a middleware, we need a route and handler 
	http.Handle("/", handler)
	http.ListenAndServe(":8000", nil)
}