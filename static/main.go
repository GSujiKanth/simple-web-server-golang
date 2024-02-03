package main

import {
	"fmt", 
	"log", 
	"net/http"
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.handle("/hello", helloHandler)
}