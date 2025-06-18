package main

import (
	"fmt"
	"net/http"
	"os"
)

const REVISION = "4.0"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello, world! (Rev #" + REVISION + ")"
	}

	_, err := fmt.Fprintln(w, response)
	if err != nil {
		return
	}
	fmt.Println("Servicing an impatient beginner's request from Openshift via GO.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
