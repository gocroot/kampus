package module

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertUser(db string, user UserSurat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertSurat(db string, surat Surat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("surat").InsertOne(context.TODO(), surat)
	if err != nil {
		fmt.Printf("InsertSurat: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKategori(db string, kategori Kategorisurat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kategori").InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Printf("InsertKategori: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertLokasi(db string, lokasi Lokasisurat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("lokasi").InsertOne(context.TODO(), lokasi)
	if err != nil {
		fmt.Printf("InsertLokasi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAbout(db string, about Aboutsurat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetUserData(telepon string) (data []UserSurat) {
	user := MongoConnect("suratdibai").Collection("users")
	filter := bson.M{"telepon": telepon}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetUserData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetNamaUser(nama string) (data []UserSurat) {
	user := MongoConnect("suratdibai").Collection("users")
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

func GetEmailData(isisurat string) (data []Surat) {
	user := MongoConnect("suratdibai").Collection("surat")
	filter := bson.M{"nosubject": isisurat}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetEmailData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertSuratData(db *mongo.Database, collect string, Isisurat string, Subject string) (InsertedID interface{}) {
	var srt Surat
	srt.Isisurat = Isisurat
	srt.Subject = Subject
	return InsertOneDoc(db, collect, srt)
}

func GetSurat(surat string) (data []Surat) {
	user := MongoConnect("suratdibai").Collection("surat")
	filter := bson.M{"subject": surat}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetSurat :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetoSurato(hiya string) (data []Surat) {
	user := MongoConnect("suratdibai").Collection("surat")
	filter := bson.M{"subject": hiya}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetSurat :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
