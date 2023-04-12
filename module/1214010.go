package module

import (
	"context"
	"fmt"
	"time"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertDafdir(db *mongo.Database, keterangan string, kehadiran string) (InsertedID interface{}) {
	var dafdir model.Dafdir
	dafdir.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafdir.Keterangan = keterangan
	dafdir.Kehadiran = kehadiran
	return InsertOneDoc(db, "dafdir", dafdir)
}

func InsertNilai(db *mongo.Database, matpel string, kkm string, grade string) (InsertedID interface{}) {
	var nilai model.Nilaisis
	nilai.Matpel = matpel
	nilai.Kkm = kkm
	nilai.Grade = grade
	return InsertOneDoc(db, "nilai", nilai)
}

func InsertDafpel(db *mongo.Database, keterangan string, poin string) (InsertedID interface{}) {
	var dafpel model.Dafpel
	dafpel.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafpel.Keterangan = keterangan
	dafpel.Poin = poin
	return InsertOneDoc(db, "dafpel", dafpel)
}

func InsertPembayaran(db *mongo.Database, keterangan string, status string) (InsertedID interface{}) {
	var pembayaran model.Pembayaransis
	pembayaran.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	pembayaran.Keterangan = keterangan
	pembayaran.Status = status
	return InsertOneDoc(db, "pembayaran", pembayaran)
}

func GetDataDafdir(ket string, db *mongo.Database, col string) (data model.Dafdir) {
	user := db.Collection(col)
	filter := bson.M{"keterangan": ket}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDafdir :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetDataNilai(matpel string, db *mongo.Database, col string) (data model.Nilaisis) {
	user := db.Collection(col)
	filter := bson.M{"matpel": matpel}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataNilai :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetDataDafpel(ket string, db *mongo.Database, col string) (data model.Dafpel) {
	user := db.Collection(col)
	filter := bson.M{"keterangan": ket}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDafpel :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetDataPembayaransaw(stats string, db *mongo.Database, col string) (data model.Pembayaransis) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPembayaran  :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
