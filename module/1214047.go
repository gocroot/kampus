package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertUserData(Nama string, Email string, Telepon string, db *mongo.Database, col string) (InsertedID interface{}) {
	srt := new(model.UserSurat)
	srt.Nama = Nama
	srt.Email = Email
	srt.Telepon = Telepon
	return InsertOneDoc(db, col, srt)
}
func InsertSurat(Isisurat string, Subject string, db *mongo.Database, col string) (InsertedID interface{}) {
	srt := new(model.Surat)
	srt.Isisurat = Isisurat
	srt.Subject = Subject
	return InsertOneDoc(db, col, srt)
}
func InsertKategoriSurat(NamaKategori string, Surat string, db *mongo.Database, col string) (InsertedID interface{}) {
	srt := new(model.Kategorisurat)
	srt.NamaKategori = NamaKategori
	srt.Surat = srt.Surat
	return InsertOneDoc(db, col, srt)
}

func GetNamaUser(nama string, db *mongo.Database, col string) (data []model.UserSurat) {
	user := db.Collection(col)
	filter := bson.M{"nama": nama}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetNamaUser :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetUserData(telepon string, db *mongo.Database, col string) (data []model.UserSurat) {
	user := db.Collection(col)
	filter := bson.M{"telepon": telepon}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetTeleponUser :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetSurat(surat string, db *mongo.Database, col string) (data []model.UserSurat) {
	user := db.Collection(col)
	filter := bson.M{"subject": surat}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetSuratUser :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func InsertSuratChat(db *mongo.Database, collect string, Isisurat string, Subject string) (InsertedID interface{}) {
	var srt model.Surat
	srt.Isisurat = Isisurat
	srt.Subject = Subject
	return InsertOneDoc(db, collect, srt)
}
