package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// by mistake if not the path is hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// hello is a get method, we won't anythinng to post in hello function
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// ideal condition
	fmt.Fprintf(w, "!hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // we are telling golang to checkout the static directory. it will automatically check the index.html
	http.Handle("/", fileServer)                        // handle the root route
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("staring server at port 8080\n") // print when server connects the port
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
