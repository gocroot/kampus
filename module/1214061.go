package nizar

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
func InsertDataTagihan(db string, datatagihan DataTagihan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("datatagihan").InsertOne(context.TODO(),datatagihan)
	if err != nil {
		fmt.Printf("InsertDataTagihan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataSudah(db string, datasudah DataSudah) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("datasudah").InsertOne(context.TODO(), datasudah)
	if err != nil {
		fmt.Printf("InsertDataSudah: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataBelum(db string, databelum DataBelum) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("databelum").InsertOne(context.TODO(), databelum)
	if err != nil {
		fmt.Printf("InsertDataBelum: %v\n", err)
	}
	return insertResult.InsertedID
}
