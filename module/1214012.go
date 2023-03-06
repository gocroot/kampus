package module

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dhs)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertNilai(nama_matkul string, kode_matkul int, nama_dosen string, sks dhs) (InsertedID interface{}) {
	var Nilai Nilai
	nilai.ID = ID
	nilai.Nama_matkul = nama_matkul
	nilai.Kode_matkul = kode_matkul
	nilai.Sks = sks
	nilai.Grade = grade
	return InsertOneDoc("adorable", "nilai", nilai)
}

func GetDhsFromKode_matkul(kode_matkul string) dhs {
	dhs := MongoConnect("adorable").Collection("nilai")
	filter := bson.M{"kode_matkul": kode_matkul}
	err := dhs.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getdhsFromKode_matkul: %v\n", err)
	}
	return dhs
}
