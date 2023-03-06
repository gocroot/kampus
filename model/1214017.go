package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataComplain struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Sistemcomp string             `bson:"sistemcomp,omitempty" json:"sistemcomp,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	Biodata    User               `bson:"user,omitempty" json:"user,omitempty"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Handphone string             `bson:"handphone,omitempty" json:"handphone,omitempty"`
}

type Helper struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Helpid    string             `bson:"helpid,omitempty" json:"helpid,omitempty"`
	Username  string             `bson:"username,omitempty" json:"username,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Handphone string             `bson:"handphone,omitempty" json:"handphone,omitempty"`
}
