package bella

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

func InsertMataKuliah(namamatakuliah string, kodematakuliah string, dosen string, sks string) (InsertedID interface{}) {
	var matakuliah matakuliah
	matakuliah.NamaMatakuliah = namamatakuliah
	matakuliah.KodeMatakuliah = kodematakuliah
	matakuliah.Dosen = dosen
	matakuliah.SKS = sks

	return InsertOneDoc("DatabaseTugas3", "matakuliah", matakuliah)
}

func InsertJadwalKuliah(hari string, jammulai string, jamselesai string, ruang string) (InsertedID interface{}) {
	var jadwalkuliah jadwalkuliah
	jadwalkuliah.Hari = hari
	jadwalkuliah.JamMulai = jammulai
	jadwalkuliah.JamSelesai = jamselesai
	jadwalkuliah.Ruang = ruang

	return InsertOneDoc("DatabaseTugas3", "jadwalkuliah", jadwalkuliah)
}

func InsertKelas(ruang string, kapasitasmhs string) (InsertedID interface{}) {
	var kelas kelas
	kelas.Ruang = ruang
	kelas.KapasitasMhs = kapasitasmhs
	
	return InsertOneDoc("DatabaseTugas3", "kelas", kelas)
}

func InsertDosen(namadosen string, kodedosen string, matakuliah string) (InsertedID interface{}) {
	var dosen dosen
	dosen.NamaDosen = namadosen
	dosen.KodeDosen = kodedosen
	dosen.Matakuliah = matakuliah

	return InsertOneDoc("DatabaseTugas3", "dosen", dosen)
}

func InsertMahasiswa(namamhs string, kelas string, prodi string) (InsertedID interface{}) {
	var mahasiswa mahasiswa
	mahasiswa.NamaMhs = namamhs
	mahasiswa.Kelas = kelas
	mahasiswa.Prodi = prodi

	return InsertOneDoc("DatabaseTugas3", "mahasiswa", mahasiswa)
}