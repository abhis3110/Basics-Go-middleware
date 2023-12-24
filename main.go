package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// middleware takes http.Handler as input and return http.Handler (as handler function)
func loggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
		// log request details before calling the next handler
		fmt.Println("Before handler is executed")
		fmt.Printf("%s %s %s\n", r. Method, r.RequestURI, time.Now())

		w.Write([]byte("Adding response via middleware\n"))
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)

		// log response details after the next handler handler has finished
		fmt.Println("After handler is executed")
		fmt.Printf("%s %s %s\n", r. Method, r.RequestURI, time.Now())
	})

}


func main() {
	// To create a middleware, we need a route and handler 
	// wrap a simple handler with the logging middleware
	handler:=http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside the handler")
	    w.Write([]byte("OK - executing handler"))
	})

	http.Handle("/", loggingMiddleware(handler))
	http.ListenAndServe(":8000", nil)
} 