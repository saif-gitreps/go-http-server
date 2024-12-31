package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err,)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello-world", helloWorldHandler)

	fmt.Printf("Server is running on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}