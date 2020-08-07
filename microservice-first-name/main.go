package main

import (
	"fmt"
	"net/http"
	"os"
)

const FIRST_NAME = "Guy"

func main() {
	http.HandleFunc("/_ready", readinessHandler)
	http.HandleFunc("/_alive", livenessHandler)
	http.HandleFunc("/get_first_name", firstNameHandler)
	http.HandleFunc("/get_full_name", fullNameHandler)
	fmt.Println("http server listening on port: 8090")
	http.ListenAndServe(":8090", nil)
}

func readinessHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handled readiness request")
	fmt.Fprintln(w, "im ready")
	return
}

func livenessHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handled liveness request")
	fmt.Fprintln(w, "im alive")
	return
}

func firstNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handled firstName request")
	fmt.Fprintln(w, FIRST_NAME)
	return
}

func fullNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handled fullName request")
	last_name, err := http.Get(os.Getenv("MICROSERVICE_LAST_NAME"))
	if err != nil {
		fmt.Println("Existed full name handler with error ", err.Error())
		http.Error(w, "full name handler failed", 500)
		return
	}
	fmt.Fprintln(w, "My full name is ", FIRST_NAME, last_name)
	return
}
