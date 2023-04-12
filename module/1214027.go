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

func InsertMataKuliah(db *mongo.Database, namamatakuliah string, kodematakuliah string, dosen string, sks string) (InsertedID interface{}) {
	var matakuliah Matakuliahbel
	matakuliah.NamaMatakuliah = namamatakuliah
	matakuliah.KodeMatakuliah = kodematakuliah
	matakuliah.Dosen = dosen
	matakuliah.SKS = sks

	return InsertOneDoc(db, "matakuliah", matakuliah)
}

func InsertJadwalKuliah(db *mongo.Database, hari string, jammulai string, jamselesai string, ruang string) (InsertedID interface{}) {
	var jadwalkuliah Jadwalkuliahbel
	jadwalkuliah.Hari = hari
	jadwalkuliah.JamMulai = jammulai
	jadwalkuliah.JamSelesai = jamselesai
	jadwalkuliah.Ruang = ruang

	return InsertOneDoc(db, "jadwalkuliah", jadwalkuliah)
}

func InsertKelas(db *mongo.Database, ruang string, kapasitasmhs string) (InsertedID interface{}) {
	var kelas Kelasbel
	kelas.Ruang = ruang
	kelas.KapasitasMhs = kapasitasmhs
	
	return InsertOneDoc(db, "kelas", kelas)
}

func InsertDosen(db *mongo.Database, namadosen string, kodedosen string, matakuliah string) (InsertedID interface{}) {
	var dosen Dosenbel
	dosen.NamaDosen = namadosen
	dosen.KodeDosen = kodedosen
	dosen.Matakuliah = matakuliah

	return InsertOneDoc(db, "dosen", dosen)
}

func InsertMahasiswa(db *mongo.Database, namamhs string, kelas string, prodi string) (InsertedID interface{}) {
	var mahasiswa Mahasiswabel
	mahasiswa.NamaMhs = namamhs
	mahasiswa.Kelas = kelas
	mahasiswa.Prodi = prodi

	return InsertOneDoc(db, "mahasiswa", mahasiswa)
}

func GetDataMatakuliah(kodematakuliah string, db *mongo.Database, col string) (data model.Matakuliahbel) {
	user := db.Collection(col)
	filter := bson.M{"kodemtkuliah": kodematakuliah}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatamatakuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataJadwalkuliah(hari string, db *mongo.Database, col string) (data model.Jadwalkuliahbel) {
	user := db.Collection(col)
	filter := bson.M{"hari": hari}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatajadwalkuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataKelas(ruang string, db *mongo.Database, col string) (data model.Kelasbel) {
	user := db.Collection(col)
	filter := bson.M{"ruang": ruang}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatakelas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataDosen(namadosen string, db *mongo.Database, col string) (data model.Dosenbel) {
	user := db.Collection(col)
	filter := bson.M{"namadosen": namadosen}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatadosen :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataMahasiswa(namamhs string, db *mongo.Database, col string) (data model.Mahasiswabel) {
	user := db.Collection(col)
	filter := bson.M{"namamhs": namamhs}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatamahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}