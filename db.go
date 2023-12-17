package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func setup() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://apoorva1937:Apoorva123@goproject.zzhjzuy.mongodb.net/")
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
	collection = database.Collection("Users")
}
