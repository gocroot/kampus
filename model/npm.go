package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dashboard struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username   string             `bson:"Username" json:"Username"`
	Email      string             `bson:"Email" json:"Email"`
	Location   string             `bson:"Location" json:"Location"`
	Orders     string             `bson:"Orders" json:"Orders"`
	Lastorders string             `bson:"Lastorders" json:"Lastorders"`
	Totalspent string             `bson:"Totalspent" json:"Totalspent"`
	Sks        string             `bson:"Sks" json:"Sks"`
}
type Mahasiswa struct {
	Nama_mhs   string `bson:"Nama_mhs" json:"Nama_mhs"`
	Email_mhs  string `bson:"email_mhs" json:"email_mhs"`
	Gambar_mhs string `bson:"gambar_mhs" json:"gambar_mhs"`
}
type Dosen struct {
	Nama      string `bson:"nama" json:"nama"`
	Jabatan   string `bson:"Jabatan" json:"Jabatan"`
	Noted     string `bson:"Noted" json:"Noted"`
	Img_dosen string `bson:"img_dosen" json:"img_dosen"`
}
type About struct {
	Pertanyaan string `bson:"pertanyaan" json:"pertanyaan"`
	Jawaban    string `bson:"jawaban" json:"jawaban"`
}
type Contacus struct {
	Phone_number string `bson:"phone_number" json:"phone_number"`
	Email        string `bson:"email" json:"email"`
}
