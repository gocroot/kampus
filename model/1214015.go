package daffa

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama         string             `bson:"nama" json:"nama"`
	Gender       string             `bson:"gender" json:"gender"`
	Email        string             `bson:"email" json:"email"`
	NoHp         string             `bson:"nohp" json:"nohp"`
}

type Pembayaran struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Datetime   primitive.DateTime `bson:"tanggal" json:"tanggal"`
	Status  	string            `bson:"status" json:"status"`
	TotalBayar  string            `bson:"total_bayar,omitempty" json:"total_bayar,omitempty"`
}

type Pendaftaran struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nis  		 string        `bson:"nis" json:"nis"`
	Nik  		 string        `bson:"nik" json:"nik"`
	NamaSiswa    string        `bson:"nama_siswa" json:"nama_siswa"`
}

type Pengumuman  struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	HasilSeleksi string           `bson:"hasil_seleksi" json:"hasil_seleksi"`
	Nilai        string           `bson:"nilai" json:"nilai"`
	Program      string           `bson:"program" json:"program"`
}

type ProgramKursus struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NamaKursus     string         `bson:"nama_kursus" json:"nama_kursus"`
	JenjangKursus  string         `bson:"jenjang_kursus" json:"jenjang_kursus"`
	Pengajar       string         `bson:"pengajar" json:"pengajar"`
}