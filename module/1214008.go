package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataKaryawan(db, nama string, status string, jabatan string, gaji string) (InsertedID interface{}) {
	var datakaryawan model.Karyawan
	datakaryawan.Nama = nama
	datakaryawan.Status = status
	datakaryawan.Jabatan = jabatan
	datakaryawan.Gaji = gaji

	return InsertOneDoc(db, "karyawan", datakaryawan)
}
func InsertKaryawan(db string, karyawan model.Karyawan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("karyawan").InsertOne(context.TODO(), karyawan)
	if err != nil {
		fmt.Printf("InsertKaryawan: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertHonor(db string, honor model.Honor) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("honor").InsertOne(context.TODO(), honor)
	if err != nil {
		fmt.Printf("InsertHonor: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertJob(db string, job model.Job) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("job").InsertOne(context.TODO(), job)
	if err != nil {
		fmt.Printf("InsertJob: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTeam(db string, team model.Team) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("team").InsertOne(context.TODO(), team)
	if err != nil {
		fmt.Printf("InsertTeam: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertAbout(db string, about model.Abouttt) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}
func GetDataKaryawan(stats string) (data []model.Karyawan) {
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
