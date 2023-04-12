package module

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertDashboard(db string, dashboard Dashboardnye) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dashboard").InsertOne(context.TODO(), dashboard)
	if err != nil {
		fmt.Printf("InsertDashboard: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertMahasiswa(db string, mahasiswa Mahasiswanye) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("mahasiswa").InsertOne(context.TODO(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDosen(db string, dosen Dosenye) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dosen").InsertOne(context.TODO(), dosen)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertAbout(db string, about Aboutnye) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertContacus(db string, contacus Contacusnye) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("contacus").InsertOne(context.TODO(), contacus)
	if err != nil {
		fmt.Printf("InsertContacus: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDataDosen(stats string) (data []Dosenye) {
	user := MongoConnect("tablerps").Collection("dosen")
	filter := bson.M{"Jabatan": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDosen :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataAbout(stats string) (data []Aboutnye) {
	user := MongoConnect("tablerps").Collection("about")
	filter := bson.M{"pertanyaan": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAbout :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataContacus(stats string) (data []Contacusnye) {
	user := MongoConnect("tablerps").Collection("contacus")
	filter := bson.M{"phone_number": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataContacus :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataDashboard(stats string) (data []Dashboardnye) {
	user := MongoConnect("tablerps").Collection("dashboard")
	filter := bson.M{"Location": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDashboard :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataMahasiswa(stats string) (data []Mahasiswanye) {
	user := MongoConnect("tablerps").Collection("mahasiswa")
	filter := bson.M{"email_mhs": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
