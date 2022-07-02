package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello hendler started")
	defer fmt.Println("server: hello hendler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)

	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8090", nil)
}
