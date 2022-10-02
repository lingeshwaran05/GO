package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull")
<<<<<<< HEAD
	Name := r.FormValue("Name")
	Address := r.FormValue("Address")
	fmt.Fprintf(w, "\nName= %s \n", Name)
	fmt.Fprintf(w, "Address= %s", Address)
=======
	NAME := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s \n", NAME)
	fmt.Fprintf(w, "Address=%s\n", address)
>>>>>>> c66bc6ade2c27a304adb7af955e7c0e6e338cc23
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
<<<<<<< HEAD
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
=======
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
>>>>>>> c66bc6ade2c27a304adb7af955e7c0e6e338cc23
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
