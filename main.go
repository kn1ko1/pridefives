package main

import (
	"fmt"
	"net/http"
	//"text/template"
)

func main() {
	//http.HandleFunc("/", formHandler)
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP status 500 - Internal Server Errors")
	}
}
