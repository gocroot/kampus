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

func InsertDataUserdaps(db *mongo.Database, nama string, gender string, email string, no_hp string) (InsertedID interface{}) {
	var datauser model.Userdaps
	datauser.Nama = nama
	datauser.Gender = gender
	datauser.Email = email
	datauser.No_hp = no_hp

	return InsertOneDoc(db, "user", datauser)
}

func InsertDataPendaftaran(db *mongo.Database, nama_siswa string, nis string, nik string) (InsertedID interface{}) {
	var datapendaftaran model.Pendaftaran
	datapendaftaran.Nama_siswa = nama_siswa
	datapendaftaran.Nis = nis
	datapendaftaran.Nik = nik

	return InsertOneDoc(db, "pendaftaran", datapendaftaran)
}

func InsertDataPembayarandaps(db *mongo.Database, status string, total_bayar string) (InsertedID interface{}) {
	var datapembayaran model.Pembayarandaps
	datapembayaran.Status = status
	datapembayaran.Total_bayar = total_bayar

	return InsertOneDoc(db, "pembayaran", datapembayaran)
}

func InsertDataPengumuman(db *mongo.Database, hasil_seleksi string, nilai string, program string) (InsertedID interface{}) {
	var datapengumuman model.Pengumuman
	datapengumuman.Hasil_seleksi = hasil_seleksi
	datapengumuman.Nilai = nilai
	datapengumuman.Program = program

	return InsertOneDoc(db, "pengumuman", datapengumuman)
}

func InsertDataKursus(db *mongo.Database, nama_kursus string, jenjang_kursus string, pengajar string) (InsertedID interface{}) {
	var datakursus model.Kursus
	datakursus.Nama_kursus = nama_kursus
	datakursus.Jenjang_kursus = jenjang_kursus
	datakursus.Pengajar = pengajar

	return InsertOneDoc(db, "kursus", datakursus)
}

func GetDataUser(stats string, db *mongo.Database, col string) (data []model.Userdaps) {
	user := db.Collection(col)
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

func GetDataPendaftaran(stats string, db *mongo.Database, col string) (data []model.Pendaftaran) {
	user := db.Collection(col)
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

func GetDataPembayaran(stats string, db *mongo.Database, col string) (data []model.Pembayarandaps) {
	user := db.Collection(col)
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

func GetDataPengumuman(stats string, db *mongo.Database, col string) (data []model.Pengumuman) {
	user := db.Collection(col)
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

func GetDataKursus(stats string, db *mongo.Database, col string) (data []model.Kursus) {
	user := db.Collection(col)
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
