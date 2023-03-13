package dipa

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

func Insertlisttamu(name string, notelp string, email string, kota string) (InsertedID interface{}) {
	var listtamu listtamu
	listtamu.Name = name
	listtamu.Notelp = notelp
	listtamu.Email = email
	listtamu.Kota = kota

	return InsertOneDoc("dblisttamu", "listtamu", listtamu)
}

func InsertUndanganRapat(lokasi string, agenda string, peserta string) (InsertedID interface{}) {
	var UndanganRapat UndanganRapat
	UndanganRapat.Lokasi = lokasi
	UndanganRapat.Agenda = agenda
	UndanganRapat.Peserta = peserta

	return InsertOneDoc("dblisttamu", "UndanganRapat", UndanganRapat)
}

func Insertpesertarapat(nama string, instansi string, status string) (InsertedID interface{}) {
	var pesertarapat pesertarapat
	pesertarapat.Nama = nama
	pesertarapat.Instansi = instansi
	pesertarapat.Status = status

	return InsertOneDoc("dblisttamu", "pesertarapat", pesertarapat)
}

func Insertwakturapat(hal string, materi string) (InsertedID interface{}) {
	var wakturapat wakturapat
	wakturapat.Hal = hal
	wakturapat.Materi = materi

	return InsertOneDoc("dblisttamu", "wakturapat", wakturapat)
}

func Insertrapatmulai(pembicara string, durasi string) (InsertedID interface{}) {
	var rapatmulai rapatmulai
	rapatmulai.Pembicara = pembicara
	rapatmulai.Durasi = durasi

	return InsertOneDoc("dblisttamu", "rapatmulai", rapatmulai)
}