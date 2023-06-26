package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertJalurPenerimaan(Jalurtanpates string, Jalurtes string, Beasiswa string, Status string, db *mongo.Database, col string) (InsertedID interface{}) {
	jp := new(model.JalurPenerimaan)
	jp.Jalurtanpates = Jalurtanpates
	jp.jalurtes = Jalurtes
	jp.Beasiswa = Beasiswa
	jp.Status = Status
	return InsertOneDoc(db, col, jp)
}
func InsertInformasi(Alur string, Catatan string, db *mongo.Database, col string) (InsertedID interface{}) {
	jp := new(model.Informasi)
	jp.Alur = Alur
	jp.Catatan = Catatan
	return InsertOneDoc(db, col, jp)
}
func InsertBiaya(Biayasemester string, db *mongo.Database, col string) (InsertedID interface{}) {
	jp := new(model.Biaya)
	jp.Biayasemester = Biayasemester
	return InsertOneDoc(db, col, jp)
}

func GetJalurPenerimaan(jalurtes string, db *mongo.Database, col string) (data []model.JalurPenerimaan) {
	user := db.Collection(col)
	filter := bson.M{"jalurtes": jalurtes}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetJalurPenerimaan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetInformasi(catatan string, db *mongo.Database, col string) (data []model.Informasi) {
	user := db.Collection(col)
	filter := bson.M{"catatan": catatan}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetInformasi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetBiaya(biayasemester string, db *mongo.Database, col string) (data []model.Biayasemester) {
	user := db.Collection(col)
	filter := bson.M{"biayasemester": biayasemester}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetBiaya :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func InsertBiayakampus(db *mongo.Database, collect string, biayasemester string,) (InsertedID interface{}) {
	var jp model.Biaya
	return InsertOneDoc(db, collect, jp)
	return InsertOneDoc(db, collect, jp)
}
