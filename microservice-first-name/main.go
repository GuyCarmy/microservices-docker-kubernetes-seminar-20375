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
	fmt.Println("Handles readiness request")
	fmt.Fprint(w, "im ready")
	return
}

func livenessHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handles liveness request")
	fmt.Fprint(w, "im alive")
	return
}

func firstNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handles firstName request")
	fmt.Fprint(w, FIRST_NAME)
	return
}
