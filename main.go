package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request){
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(response, "POST successful")
	name := request.FormValue("name")
	fmt.Fprintf(response, "Hello %s\n", name )
}

func helloHandler(response http.ResponseWriter, request *http.Request){
if request.URL.Path != "/hello"{
	http.Error(response, "404 not found", http.StatusNotFound)
	return
}
	if request.Method != "GET" {
		http.Error(response, "Unsuported method", http.StatusNotFound)
	return
	}

	fmt.Fprintf(response,"Hello this is a route to say hello")
}

func main()	{
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}