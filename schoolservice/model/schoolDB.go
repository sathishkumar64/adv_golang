package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//School type is expose services.
type School struct {
	SchoolId   string  `json:"school_id"`
	SchoolName string  `json:"school_name"`
	EduMode    string  `json:"edu_mode"`
	Address    Address `json:"address"`
	//rating     float32
}

//Address type is expose services.
type Address struct {
	Address string `json:"address"`
	State   string `json:"state"`
	City    string `json:"city"`
}

var (
	DB *mongo.Client
	Schooldb *mongo.Collection
	MongoCtx  context.Context
)

//DB_connect
func DbConnect() {
	fmt.Println("Connecting to MongoDB...")
	MongoCtx = context.Background()
	DB, err := mongo.Connect(MongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping(MongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	Schooldb = DB.Database("schoolservice").Collection("school")
	//fmt.Println("Closing MongoDB connection")
	//DB.Disconnect(MongoCtx)
	//fmt.Println("Done.")
}
