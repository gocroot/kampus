package module

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

func InsertMahasiswa(nm string, npm string, smt string, kls string, prd string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nm
	mahasiswa.NPM = npm
	mahasiswa.Semester = smt
	mahasiswa.Kelas = kls
	mahasiswa.Prodi_kampus = prd

	return InsertOneDoc("dbmahasiswa", "mahasiswa", mahasiswa)
}

func InsertPresensi(khd string, ktr string) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Kehadiran = khd
	presensi.Keterangan = ktr

	return InsertOneDoc("dbpresensi", "presensi", presensi)
}

func InsertMataKuliah(kdmk string, nmmk string, sks string, jrs string) (InsertedID interface{}) {
	var matakuliah MataKuliah
	matakuliah.KodeMK = kdmk
	matakuliah.NamaMK = nmmk
	matakuliah.SKS = sks
	matakuliah.Jurusan = jrs

	return InsertOneDoc("dbmatakuliah", "matakuliah", matakuliah)
}

func InsertJadwalKuliah(matkul string, hari string, jmulai string, jselesai string, ruangan string) (InsertedID interface{}) {
	var jadwalkuliah JadwalKuliah
	jadwalkuliah.MataKuliah = matkul
	jadwalkuliah.Hari = hari
	jadwalkuliah.JamMulai = jmulai
	jadwalkuliah.JamSelesai = jselesai
	jadwalkuliah.Ruangan = ruangan

	return InsertOneDoc("dbjadwalkuliah", "jadwalkuliah", jadwalkuliah)
}

func InsertDosen(nidn string, nama string, matkul string) (InsertedID interface{}) {
	var dosen Dosen
	dosen.NIDN = nidn
	dosen.Nama = nama
	dosen.MataKuliah = matkul

	return InsertOneDoc("dbdosen", "dosen", dosen)
}
