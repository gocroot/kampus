package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Profil struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Gambar    string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
	Bio       string             `bson:"bio,omitempty" json:"bio,omitempty"`
	Alamat    string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Pekerjaan string             `bson:"pekerjaan,omitempty" json:"pekerjaan,omitempty"`
}

type PengalamanKerja struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	NamaPerus   string             `bson:"nama_perusahaan,omitempty" json:"nama_perusahaan,omitempty"`
	Posisi      string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	TanggalMulai string            `bson:"tanggal_mulai,omitempty" json:"tanggal_mulai,omitempty"`
	TanggalSelesai string          `bson:"tanggal_selesai,omitempty" json:"tanggal_selesai,omitempty"`
}

type Pendidikan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	NamaInstitusi string            `bson:"nama_institusi,omitempty" json:"nama_institusi,omitempty"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	TanggalMulai string            `bson:"tanggal_mulai,omitempty" json:"tanggal_mulai,omitempty"`
	TanggalSelesai string          `bson:"tanggal_selesai,omitempty" json:"tanggal_selesai,omitempty"`
}

type Keahlian struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	NamaKeahlian string           `bson:"nama_keahlian,omitempty" json:"nama_keahlian,omitempty"`
	Tingkat    int               `bson:"tingkat,omitempty" json:"tingkat,omitempty"`
}
//upload ulang