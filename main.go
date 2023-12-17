package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	setup()
	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/createUser", CreatePerson).Methods("POST")
	router.HandleFunc("/getUsers", GetPeople).Methods("GET")
	router.HandleFunc("/getUser/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/updateUser/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", DeletePerson).Methods("DELETE")

	// Set up and start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}