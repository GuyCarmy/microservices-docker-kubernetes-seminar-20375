package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("Handles readiness request")
	fmt.Fprintln(w, "im ready")
	return
}

func livenessHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handles liveness request")
	fmt.Fprintln(w, "im alive")
	return
}

func firstNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handles firstName request")
	fmt.Fprintln(w, FIRST_NAME)
	return
}

func fullNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handles fullName request")
	get_last_name_url := os.Getenv("MICROSERVICE_LAST_NAME_URL") + "/get_last_name"
	response, err := http.Get(get_last_name_url)
	if err != nil {
		fmt.Println("Existed full name handler with error ", err.Error())
		http.Error(w, "full name handler failed", 500)
		return
	}
	last_name, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Existed full name handler with error ", err.Error())
		http.Error(w, "full name handler failed", 500)
		return
	}
	fmt.Fprintln(w, "My full name is", FIRST_NAME, string(last_name))
	return
}
