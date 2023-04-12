package module

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(db string, user Userdaps) (insertedID interface{}) {
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
func InsertPembayaran(db string, pembayaran Pembayarandaps) (insertedID interface{}) {
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

func GetDataUser(stats string) (data []User) {
	user := MongoConnect("proyek-2").Collection("user")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataUser :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPendaftaran(stats string) (data []Pendaftaran) {
	user := MongoConnect("proyek-2").Collection("pendaftaran")
	filter := bson.M{"nis": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPendaftaran :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPembayaran(stats string) (data []Pembayaran) {
	user := MongoConnect("proyek-2").Collection("pembayaran")
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPembayaran :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPengumuman(stats string) (data []Pengumuman) {
	user := MongoConnect("proyek-2").Collection("pengumuman")
	filter := bson.M{"program": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPengumuman :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataKursus(stats string) (data []Kursus) {
	user := MongoConnect("proyek-2").Collection("kursus")
	filter := bson.M{"pengajar": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataKursus :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
