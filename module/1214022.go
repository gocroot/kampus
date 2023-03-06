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
func InsertDashboard(db string, dashboard Dashboard) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dashboard").InsertOne(context.TODO(), dashboard)
	if err != nil {
		fmt.Printf("InsertDashboard: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertMahasiswa(db string, mahasiswa Mahasiswa) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("mahasiswa").InsertOne(context.TODO(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDosen(db string, dosen Dosen) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dosen").InsertOne(context.TODO(), dosen)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
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
func InsertContacus(db string, contacus Contacus) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("contacus").InsertOne(context.TODO(), contacus)
	if err != nil {
		fmt.Printf("InsertContacus: %v\n", err)
	}
	return insertResult.InsertedID
}
