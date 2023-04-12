package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Userdaps struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama   string             `bson:"nama" json:"nama"`
	Gender string             `bson:"gender" json:"gender"`
	Email  string             `bson:"email" json:"email"`
	No_hp  string             `bson:"nohp" json:"nohp"`
}

type Pendaftaran struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nis        string             `bson:"nis" json:"nis"`
	Nik        string             `bson:"nik" json:"nik"`
	Nama_siswa string             `bson:"nama_siswa" json:"nama_siswa"`
}

type Pembayarandaps struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Datetime    primitive.DateTime `bson:"tanggal" json:"tanggal"`
	Status      string             `bson:"status" json:"status"`
	Total_bayar string             `bson:"total_bayar,omitempty" json:"total_bayar,omitempty"`
}

type Pengumuman struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Hasil_seleksi string             `bson:"hasil_seleksi" json:"hasil_seleksi"`
	Nilai         string             `bson:"nilai" json:"nilai"`
	Program       string             `bson:"program" json:"program"`
}

type Kursus struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nama_kursus    string             `bson:"nama_kursus" json:"nama_kursus"`
	Jenjang_kursus string             `bson:"jenjang_kursus" json:"jenjang_kursus"`
	Pengajar       string             `bson:"pengajar" json:"pengajar"`
}
