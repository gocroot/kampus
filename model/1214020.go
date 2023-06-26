package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama         string             `bson:"nama" json:"nama"`
	NPM          string             `bson:"npm" json:"npm"`
	Semester     string             `bson:"semester" json:"semester"`
	Kelas        string             `bson:"kelas" json:"kelas"`
	Prodi_kampus string             `bson:"prodi_kampus" json:"prodi_kampus"`
}

type Presensi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Datetime   string             `bson:"tanggal" json:"tanggal"`
	Kehadiran  string             `bson:"kehadiran" json:"kehadiran"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
}

type MataKuliah struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	KodeMK  string             `bson:"kode_mk" json:"kode_mk"`
	NamaMK  string             `bson:"nama_mk" json:"nama_mk"`
	SKS     string             `bson:"sks" json:"sks"`
	Jurusan string             `bson:"jurusan" json:"jurusan"`
}

type JadwalKuliah struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MataKuliah string             `bson:"mata_kuliah" json:"mata_kuliah"`
	Hari       string             `bson:"hari" json:"hari"`
	JamMulai   string             `bson:"jam_mulai" json:"jam_mulai"`
	JamSelesai string             `bson:"jam_selesai" json:"jam_selesai"`
	Ruangan    string             `bson:"ruangan" json:"ruangan"`
}

type Dosen struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NIDN       string             `bson:"nidn" json:"nidn"`
	Nama       string             `bson:"nama" json:"nama"`
	MataKuliah string             `bson:"mata_kuliah" json:"mata_kuliah"`
}
