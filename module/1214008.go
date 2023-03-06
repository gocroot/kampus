package module

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
func InsertKaryawan(db string, karyawan Karyawan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("karyawan").InsertOne(context.TODO(), karyawan)
	if err != nil {
		fmt.Printf("InsertKaryawan: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertHonor(db string, honor Honor) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("honor").InsertOne(context.TODO(), honor)
	if err != nil {
		fmt.Printf("InsertHonor: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertJob(db string, job Job) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("job").InsertOne(context.TODO(), job)
	if err != nil {
		fmt.Printf("InsertJob: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTeam(db string, team Team) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("team").InsertOne(context.TODO(), team)
	if err != nil {
		fmt.Printf("InsertTeam: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertAbout(db string, about About) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}
