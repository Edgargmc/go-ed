package main

import (
	"fmt"
	"net/http"
)

func test() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(write http.ResponseWriter, request *http.Request)  {
	write.WriteHeader(200)
	fmt.Fprintf(write, "Welcome...")
}
