package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mahasiswa struct
type Mahasiswa struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap,omitempty" json:"nama_lengkap,omitempty"`
	NPM         string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NilaiTugas  float32            `bson:"nilai_tugas,omitempty" json:"nilai_tugas,omitempty"`
	NilaiUTS    float32            `bson:"nilai_uts,omitempty" json:"nilai_uts,omitempty"`
	NilaiUAS    float32            `bson:"nilai_uas,omitempty" json:"nilai_uas,omitempty"`
	Grade       string             `bson:"grade,omitempty" json:"grade,omitempty"`
}

// SyaratGrade struct
type SyaratGrade struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nilai  float32            `bson:"nilai,omitempty" json:"nilai,omitempty"`
	Syarat string             `bson:"syarat,omitempty" json:"syarat,omitempty"`
	Grade  string             `bson:"grade,omitempty" json:"grade,omitempty"`
}

// Prestasi struct
type Prestasi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap,omitempty" json:"nama_lengkap,omitempty"`
	NPM         string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Prestasi    string             `bson:"prestasi,omitempty" json:"prestasi,omitempty"`
}

// Nilai struct
type Nilai struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MahasiswaID primitive.ObjectID `bson:"mahasiswa_id,omitempty" json:"mahasiswa_id,omitempty"`
	NPM         string             `bson:"npm,omitempty" json:"npm,omitempty"`
	AlamatID    primitive.ObjectID `bson:"alamat_id,omitempty" json:"alamat_id,omitempty"`
	NilaiTugas  float32            `bson:"nilai_tugas,omitempty" json:"nilai_tugas,omitempty"`
	NilaiUTS    float32            `bson:"nilai_uts,omitempty" json:"nilai_uts,omitempty"`
	NilaiUAS    float32            `bson:"nilai_uas,omitempty" json:"nilai_uas,omitempty"`
}

// Alamat struct
type Alamat struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MahasiswaID primitive.ObjectID `bson:"mahasiswa_id,omitempty" json:"mahasiswa_id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap,omitempty" json:"nama_lengkap,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}
