package iqbal

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertJalurPenerimaan(db string, jalurpenerimaan JalurPenerimaan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("jalurpenerimaan").InsertOne(context.TODO(), jalurpenerimaan)
	if err != nil {
		fmt.Printf("InsertJalurPenerimaan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertInformasi(db string, informasi Informasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("informasi").InsertOne(context.TODO(), informasi)
	if err != nil {
		fmt.Printf("InsertInformasi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertBiaya(db string, biaya Biaya) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("biaya").InsertOne(context.TODO(), biaya)
	if err != nil {
		fmt.Printf("InsertBiaya: %v\n", err)
	}
	return insertResult.InsertedID
}