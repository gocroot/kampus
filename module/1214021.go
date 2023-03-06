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
func InsertMahasiswa(db string, mahasiswa Mahasiswa) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("mahasiswa").InsertOne(context.TODO(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertSyaratGrade(db string, syaratGrade SyaratGrade) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("syarat_grade").InsertOne(context.TODO(), syaratGrade)
	if err != nil {
		fmt.Printf("InsertSyaratGrade: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPrestasi(db string, prestasi Prestasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("prestasi").InsertOne(context.TODO(), prestasi)
	if err != nil {
		fmt.Printf("InsertPrestasi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertNilai(db string, nilai Nilai) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("nilai").InsertOne(context.TODO(), nilai)
	if err != nil {
		fmt.Printf("InsertNilai: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAlamat(db string, alamat Alamat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("alamat").InsertOne(context.TODO(), alamat)
	if err != nil {
		fmt.Printf("InsertAlamat: %v\n", err)
	}
	return insertResult.InsertedID
}
