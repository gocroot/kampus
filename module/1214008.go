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
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataKaryawan(db, nama string, status string, jabatan string, gaji string) (InsertedID interface{}) {
	var datakaryawan Karyawan
	datakaryawan.Nama = nama
	datakaryawan.Status = status
	datakaryawan.Jabatan = jabatan
	datakaryawan.Gaji = gaji

	return InsertOneDoc(db, "karyawan", datakaryawan)
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
func InsertAbout(db string, about Abouttt) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}
func GetDataKaryawan(stats string) (data []Karyawan) {
	user := MongoConnect("penggajian").Collection("karyawan")
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataKaryawan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataHonor(stats string) (data []Honor) {
	user := MongoConnect("penggajian").Collection("honor")
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataHonor :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataJob(stats string) (data []Job) {
	user := MongoConnect("penggajian").Collection("job")
	filter := bson.M{"namajob": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataJob :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataTeam(stats string) (data []Team) {
	user := MongoConnect("penggajian").Collection("team")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataTeam :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
