package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// insert data ke koleksi Mahasiswa
func insertMahasiswa(db *mongo.Database, mhs model.Mahasiswawoylah) (insertedID interface{}, err error) {
	result, err := db.Collection("Mahasiswa").InsertOne(context.Background(), mhs)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Mahasiswa: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Jurusan
func insertJurusan(db *mongo.Database, jur model.Jurusanwoy) (insertedID interface{}, err error) {
	result, err := db.Collection("Jurusan").InsertOne(context.Background(), jur)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Jurusan: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Nilai
func insertNilai(db *mongo.Database, n model.Nilaiwoy) (insertedID interface{}, err error) {
	result, err := db.Collection("Nilai").InsertOne(context.Background(), n)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Nilai: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Alamat
func insertAlamat(db *mongo.Database, alamat model.Alamatwoy) (insertedID interface{}, err error) {
	result, err := db.Collection("Alamat").InsertOne(context.Background(), alamat)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Alamat: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi NPM
func insertNPM(db *mongo.Database, npm model.NPMwoylah) (insertedID interface{}, err error) {
	result, err := db.Collection("NPM").InsertOne(context.Background(), npm)
	if err != nil {
		return nil, fmt.Errorf("failed to insert NPM: %v", err)
	}
	return result.InsertedID, nil
}
