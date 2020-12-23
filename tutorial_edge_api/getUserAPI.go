package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// struct for version
type Versions struct {
	Version string `json : "version"`
}

var v Versions

//homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Endpoint hit: homepage")
}

//get_version : returns -> json of sudo version
func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get_version")
	json.NewEncoder(w).Encode(v)
}

//request handler
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage).Methods("GET")
	myRouter.HandleFunc("/get_version", getVersion).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	v.Version = "1.15"
	handleRequest()
}
