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


// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		showError(w, "404 Page not found.", 404)
// 		return
// 	}

// 	switch r.Method {
// 	case "GET":
// 		http.ServeFile(w, r, "index.html")
// 	case "POST":
// 		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Print("HTTP status 500 - Internal Server Errors", err)
// 			return
// 		}
// 	}
// }

// func showError(w http.ResponseWriter, message string, statusCode int) {
// 	t, err := template.ParseFiles("template/errors.html")
// 	if err == nil {
// 		w.WriteHeader(statusCode)
// 		t.Execute(w, message)
// 		return
// 	}
// }
