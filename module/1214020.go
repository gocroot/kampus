package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertMahasiswa(Nama string, NPM string, Semester string, Kelas string, Prodi_Kampus string, db *mongo.Database, col string) (InsertedID interface{}) {
	mhs := new(model.Mahasiswa)
	mhs.Nama = Nama
	mhs.NPM = NPM
	mhs.Semester = Semester
	mhs.Kelas = Kelas
	mhs.Prodi_Kampus = Prodi_Kampus
	return InsertOneDoc(db, col, mhs)
}

func InsertPresensi(Datetime string, Kehadiran string, Keterangan string, db *mongo.Database, col string) (InsertedID interface{}) {
	prs := new(model.Presensi)
	prs.Datetime = Datetime
	prs.Kehadiran = Kehadiran
	prs.Keterangan = Keterangan
	return InsertOneDoc(db, col, prs)
}

func InsertMataKuliah(KodeMK string, NamaMK string, SKS string, Jurusan string, db *mongo.Database, col string) (InsertedID interface{}) {
	mk := new(model.MataKuliah)
	mk.KodeMK = KodeMK
	mk.NamaMK = NamaMK
	mk.SKS = SKS
	mk.Jurusan = Jurusan
	return InsertOneDoc(db, col, mk)
}

func InsertJadwalKuliah(MataKuliah string, Hari string, JamMulai string, JamSelesai string, Ruangan string, db *mongo.Database, col string) (InsertedID interface{}) {
	jk := new(model.JadwalKuliah)
	jk.MataKuliah = MataKuliah
	jk.Hari = Hari
	jk.JamMulai = JamMulai
	jk.JamSelesai = JamSelesai
	jk.Ruangan = Ruangan
	return InsertOneDoc(db, col, jk)
}

func InsertDosen(NIDN string, Nama string, MataKuliah string, db *mongo.Database, col string) (InsertedID interface{}) {
	ds := new(model.Dosen)
	ds.NIDN = NIDN
	ds.Nama = Nama
	ds.MataKuliah = MataKuliah
	return InsertOneDoc(db, col, ds)
}

func GetMahasiswa(Nama string, db *mongo.Database, col string) (data []model.Mahasiswaa) {
	user := db.Collection(col)
	filter := bson.M{"nama": nama}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetNPMMahasiswa(NPM string, db *mongo.Database, col string) (data []model.Mahasiswa) {
	user := db.Collection(col)
	filter := bson.M{"npm": npm}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetPresensi(kehadiran string, db *mongo.Database, col string) (data []model.Presensi) {
	user := db.Collection(col)
	filter := bson.M{"kehadiran": kehadiran}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return
}
