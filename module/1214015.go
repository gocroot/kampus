package daffa

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func InsertUser(nm string, gnr string, eml string, nhp string) (InsertedID interface{}) {
	var user User
	user.Nama = nm
	user.Gender = gnr
	user.Email = eml
	user.NoHp = nhp

	return InsertOneDoc("dbuser", "user", user)
}

func InsertPembayaran(sts string, tbr string) (InsertedID interface{}) {
	var pembayaran Pembayaran
	pembayaran.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	pembayaran.Status = sts
	pembayaran.TotalBayar = tbr

	return InsertOneDoc("dbpembayaran", "pembayaran", pembayaran)
}

func InsertPendaftaran(nis string, nik string, nsa string) (InsertedID interface{}) {
	var pendaftaran Pendaftaran
	pendaftaran.Nis = nis
	pendaftaran.Nik = nik
	pendaftaran.NamaSiswa = nsa

	return InsertOneDoc("dbpendaftaran", "pendaftaran", pendaftaran)
}

func InsertPengumuman(hsi string, nli string, pgm string) (InsertedID interface{}) {
	var pengumuman Pengumuman
	pengumuman.HasilSeleksi = hsi
	pengumuman.Nilai = nli
	pengumuman.Program = pgm

	return InsertOneDoc("dbpengumuman", "pengumuman", pengumuman)
}

func InsertProgramKursus(nks string, jks string, pgr string) (InsertedID interface{}) {
	var programkursus ProgramKursus
	programkursus.NamaKursus = nks
	programkursus.JenjangKursus = jks
	programkursus.Pengajar = pgr

	return InsertOneDoc("dbprogramkursus", "programkursus", programkursus)
}