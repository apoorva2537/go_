package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Person represents a person entity
type Person struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Age  int    `json:"age,omitempty" bson:"age,omitempty"`
	City string `json:"city,omitempty" bson:"city,omitempty"`
}

// CreatePerson creates a new person in the MongoDB collection
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.InsertedID)
}

// GetPeople returns all people from the MongoDB collection
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}

// GetPerson returns a specific person from the MongoDB collection by ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var person Person
	err := collection.FindOne(context.TODO(), bson.M{"_id": personID}).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

// UpdatePerson updates a person in the MongoDB collection by ID
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	var updatedPerson Person
	_ = json.NewDecoder(r.Body).Decode(&updatedPerson)

	result, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": personID}, updatedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.ModifiedCount)
}

// DeletePerson deletes a person from the MongoDB collection by ID
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID := params["id"]

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": personID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.DeletedCount)
}
