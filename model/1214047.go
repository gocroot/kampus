package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSurat struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama    string             `bson:"nama" json:"nama"`
	Email   string             `bson:"email" json:"email"`
	Telepon string             `bson:"telepon" json:"telepon"`
}
type Surat struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Isisurat string             `bson:"isisurat" json:"isisurat"`
	Subject  string             `bson:"subject" json:"subject"`
}
type Kategorisurat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NamaKategori string             `bson:"nama_kategori" json:"nama_kategori"`
	Surat        []Surat            `bson:"surat" json:"surat"`
}
type Lokasisurat struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Region string             `bson:"region,omitempty" json:"region,omitempty"`
}
type Aboutsurat struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Isi_satu string             `bson:"isi_satu,omitempty" json:"isi_satu,omitempty"`
	Isi_dua  string             `bson:"isi_dua,omitempty" json:"isi_dua,omitempty"`
	Image    string             `bson:"image,omitempty" json:"image,omitempty"`
}

type Returnsurat struct {
	Status     string `bson:"status,omitempty" json:"status,omitempty"`
	Message    string `bson:"message,omitempty" json:"message,omitempty"`
	InsertedID string `bson:"insertedid,omitempty" json:"insertedid,omitempty"`
}
