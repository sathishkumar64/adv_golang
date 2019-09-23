package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Handler Started")
	defer log.Printf("Handler Ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
