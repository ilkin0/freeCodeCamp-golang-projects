package main

import (
	f "fmt"
	l "log"
	net "net/http"
)

func main() {

	fs := net.FileServer(net.Dir("./static"))

	net.Handle("/", fs)
	net.HandleFunc("/hello", helloHandler)
	net.HandleFunc("/form", formHandler)

	f.Println("Starting server at port 8080")

	if err := net.ListenAndServe(":8080", nil); err != nil {
		l.Fatal(err)
	}
}

func formHandler(writer net.ResponseWriter, request *net.Request) {
	if request.URL.Path != "/form" {
		net.Error(writer, "404 Not Found", net.StatusNotFound)
		return
	}

	if request.Method != net.MethodPost {
		net.Error(writer, f.Sprintf("%s method is not supported", request.Method), net.StatusMethodNotAllowed)
		return
	}

	if err := request.ParseForm(); err != nil {
		f.Fprintf(writer, "Parseform() err: %v", err)
		return
	}

	f.Fprintf(writer, "POST request successful!\n")
	
	name := request.FormValue("name")
	address := request.FormValue("address")
	
	f.Fprintf(writer, "Name is: %s\n", name)
	f.Fprintf(writer, "Address is: %s\n", address)
}

func helloHandler(writer net.ResponseWriter, request *net.Request) {
	if request.URL.Path != "/hello" {
		net.Error(writer, "404 Not Found", net.StatusNotFound)
		return
	}

	if request.Method != net.MethodGet {
		net.Error(writer, f.Sprintf("%s method is not supported", request.Method), net.StatusMethodNotAllowed)
		return
	}

	f.Fprintf(writer, "hello!")
}
