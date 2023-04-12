package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertKaryawan(db *mongo.Database, nama string, status string, jabatan string, gaji string) (InsertedID interface{}) {
	var karyawan model.Karyawan
	karyawan.Nama = nama
	karyawan.Status = status
	karyawan.Jabatan = jabatan
	karyawan.Gaji = gaji

	return InsertOneDoc(db, "karyawan", karyawan)
}

func InsertDataHonor(db *mongo.Database, nama string, status string, jabatan string, gaji string) (InsertedID interface{}) {
	var datahonor model.Honor
	datahonor.Nama = nama
	datahonor.Status = status
	datahonor.Jabatan = jabatan
	datahonor.Gaji = gaji

	return InsertOneDoc(db, "honor", datahonor)
}
func InsertDataJob(db *mongo.Database, nama string) (InsertedID interface{}) {
	var datajob model.Job
	datajob.Namajob = nama

	return InsertOneDoc(db, "job", datajob)
}
func InsertDataTeam(db *mongo.Database, nama string, deskripsi string) (InsertedID interface{}) {
	var datateam model.Team
	datateam.Nama = nama
	datateam.Deskripsi = deskripsi

	return InsertOneDoc(db, "team", datateam)
}
func InsertDataAbouttt(db *mongo.Database, isi_satu string, isi_dua string, image string) (InsertedID interface{}) {
	var dataabout model.Abouttt
	dataabout.Isi_satu = isi_satu
	dataabout.Isi_dua = isi_dua
	dataabout.Image = image

	return InsertOneDoc(db, "about", dataabout)
}
func GetKaryawan(stats string, db *mongo.Database, col string) (data []model.Karyawan) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetKaryawan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataHonor(stats string, db *mongo.Database, col string) (data []model.Honor) {
	user := db.Collection(col)
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

func GetDataJob(stats string, db *mongo.Database, col string) (data []model.Job) {
	user := db.Collection(col)
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
func GetDataTeam(stats string, db *mongo.Database, col string) (data []model.Team) {
	user := db.Collection(col)
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
