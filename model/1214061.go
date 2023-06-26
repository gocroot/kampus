package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTagihan struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama    string             `bson:"nama" json:"nama"`
	Email   string             `bson:"email" json:"email"`
	Telepon string             `bson:"telepon" json:"telepon"`
}
type Tagihan struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Isitagihan string             `bson:"isitagihan" json:"isitagihan"`
	subject  string             `bson:"subject" json:"subject"`
}

type DataTagihan struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama_Mahasiswa string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Program_studi  string             `bson:"program_studi" json:"program_studi"`
	Jumlah  string             `bson:"jumlah" json:"jumlah"`
	Status  string             `bson:"status" json:"status"`
}
type DataSudah struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama_Mahasiswa string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Status   string          `bson:"status" json:"status"`
}

type DataBelum struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama_Mahasiswa string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Status   string          `bson:"status" json:"status"`

}