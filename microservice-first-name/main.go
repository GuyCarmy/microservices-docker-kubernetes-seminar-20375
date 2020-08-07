package main

import (
	"fmt"
	"net/http"
)

const FIRST_NAME = "Guy"

func main() {
	http.HandleFunc("/_ready", readinessHandler)
	http.HandleFunc("/_alive", livenessHandler)
	http.HandleFunc("/get_first_name", firstNameHandler)
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
