package iqbal

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JalurPenerimaan struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Jalurtanpates 	string             `bson:"jalurtanpates" json:"jalurtanpates"`
	Jalurtes    	string             `bson:"jalurtes" json:"jalurtes"`
	Beasiswa        string             `bson:"Beasiswa" json:"Beasiswa"`
	Status          string             `bson:"status" json:"status"`
}
type Informasi struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Alur 		string             `bson:"alur" json:"alur"`
	Catatan     string             `bson:"catatan" json:"catatan"`
}

type Biaya struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biayasemester string             `bson:"biayasemester" json:"biayasemester"`
}