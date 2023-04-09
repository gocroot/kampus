package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama    string             `bson:"nama" json:"nama"`
	Status  string             `bson:"status" json:"status"`
	Jabatan string             `bson:"jabatan" json:"jabatan"`
	Gaji    string             `bson:"gaji" json:"gaji"`
}

type Honor struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama    string             `bson:"nama" json:"nama"`
	Status  string             `bson:"status" json:"status"`
	Jabatan string             `bson:"jabatan" json:"jabatan"`
	Gaji    string             `bson:"gaji" json:"gaji"`
}
type Job struct {
	Namajob string `bson:"namajob" json:"namajob"`
}
type Team struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama" json:"nama"`
	Deskripsi string             `bson:"deskripsi" json:"deskripsi"`
}
type Abouttt struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Isi_satu string             `bson:"isi_satu,omitempty" json:"isi_satu,omitempty"`
	Isi_dua  string             `bson:"isi_dua,omitempty" json:"isi_dua,omitempty"`
	Image    string             `bson:"image,omitempty" json:"image,omitempty"`
}
