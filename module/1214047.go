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
	insertResult, err := MongoConnect(db).Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertSurat(db string, surat Surat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("surat").InsertOne(context.TODO(), surat)
	if err != nil {
		fmt.Printf("InsertSurat: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKategori(db string, kategori Kategori) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kategori").InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Printf("InsertKategori: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertLokasi(db string, lokasi Lokasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("lokasi").InsertOne(context.TODO(), lokasi)
	if err != nil {
		fmt.Printf("InsertLokasi: %v\n", err)
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
