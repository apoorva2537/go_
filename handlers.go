package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Response structure for JSON responses
type Response struct {
	Status int `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
// Person represents a person entity
type Person struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Age  int    `json:"age,omitempty" bson:"age,omitempty"`
	Branch string `json:"Branch,omitempty" bson:"Branch,omitempty"`
	City string `json:"city,omitempty" bson:"city,omitempty"`
}
// ...

// CreatePerson creates a new person in the MongoDB collection
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Println("Error inserting person:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to insert person")
		return
	}

	response := Response{
		Status:201,
		Message: "User created successfully",
		Data:    result.InsertedID,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// GetPeople returns all people from the MongoDB collection
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error finding people:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve people")
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			log.Println("Error decoding person:", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve people")
			return
		}
		people = append(people, person)
	}

	response := Response{
		Status:200,
		Message: "Fetched user successfully",
		Data:    people,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetPerson returns a specific person from the MongoDB collection by ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var person Person
	err := collection.FindOne(context.TODO(), bson.M{"_id": personID}).Decode(&person)
	if err != nil {
		log.Println("Error finding person:", err)
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	response := Response{
		Status:200,
		Message: "Retrieved user successfully",
		Data:    person,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// UpdatePerson updates a person in the MongoDB collection by ID
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var updatedPerson Person
	_ = json.NewDecoder(r.Body).Decode(&updatedPerson)
	filter := bson.M{"_id": personID}
	update := bson.M{"$set": updatedPerson}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result Person
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&result)
	if err != nil {
		log.Println("Error updating person:", err)
		respondWithError(w, http.StatusInternalServerError, "Id Not Found")
		return
	}

	response := Response{
		Status:200,
		Message: fmt.Sprintf("User updated with ID %s", personID),
		Data:    result,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// DeletePerson deletes a person from the MongoDB collection by ID
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]
	var deletedPerson Person
	err := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": personID}).Decode(&deletedPerson)
	if err != nil {
		log.Println("Error deleting person:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete User Or User Already Deleted")
		return
	}

	response := Response{
		Status:200,
		Message: fmt.Sprintf("Deleted person with ID %s", personID),
		Data:deletedPerson,
	
	}

	respondWithJSON(w, http.StatusOK, response)
}

// respondWithError sends a JSON response with an error message
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	
	response := Response{
		Status:400,
		Message: message,

	}

	respondWithJSON(w, statusCode, response)
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
