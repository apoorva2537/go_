package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func setupDB() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://ANKIT:ANKIT123@go-project.s1rwtd6.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Set the collection variable
	database := client.Database("test")
	collection = database.Collection("people")
}

// InsertPerson inserts a person into the MongoDB collection
func InsertPerson(person Person) interface{} {
	insertResult, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID
}

// FindPerson finds a person in the MongoDB collection
func FindPerson(filter bson.D) Person {
	var result Person
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// UpdatePerson updates a person in the MongoDB collection
func UpdatePerson(filter bson.D, update bson.D) int64 {
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult.ModifiedCount
}

// DeletePerson deletes a person from the MongoDB collection
func DeletePerson(filter bson.D) int64 {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}
