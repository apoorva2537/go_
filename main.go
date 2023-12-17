package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// Example usage
	setupDB()
	insertedID := InsertPerson(Person{"John Doe", 30, "New York"})
	fmt.Println("Inserted document with ID:", insertedID)

	// Find document
	filter := bson.D{{"name", "John Doe"}}
	foundPerson := FindPerson(filter)
	fmt.Printf("Found document: %+v\n", foundPerson)

	// Update document
	update := bson.D{{"$set", bson.D{{"age", 31},{"name", "Ankit Gupta"}}}}
	modifiedCount := UpdatePerson(filter, update)
	fmt.Printf("Modified %v document(s)\n", modifiedCount)

	// Delete document
	// deletedCount := DeletePerson(filter)
	// fmt.Printf("Deleted %v document(s)\n", deletedCount)
}
