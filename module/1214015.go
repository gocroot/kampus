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
func InsertUser(db string, user User) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("user").InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPendaftaran(db string, pendaftaran Pendaftaran) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("pendaftaran").InsertOne(context.TODO(), pendaftaran)
	if err != nil {
		fmt.Printf("InsertPendaftaran: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPembayaran(db string, pembayaran Pembayaran) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("pembayaran").InsertOne(context.TODO(), pembayaran)
	if err != nil {
		fmt.Printf("InsertPembayaran: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPengumuman(db string, pengumuman Pengumuman) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("pengumuman").InsertOne(context.TODO(), pengumuman)
	if err != nil {
		fmt.Printf("InsertPengumuman: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertKursus(db string, kursus Kursus) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kursus").InsertOne(context.TODO(), kursus)
	if err != nil {
		fmt.Printf("InsertKursus: %v\n", err)
	}
	return insertResult.InsertedID
}
