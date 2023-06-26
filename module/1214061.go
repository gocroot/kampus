package module

import (
	"context"
	"fmt"
	"os"
	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUserTagihan(Nama string, Email string, Telepon string, db *mongo.Database, col string) (InsertedID interface{}) {
	srt := new(model.UserTagihan)
	srt.Nama = Nama
	srt.Email = Email
	srt.Telepon = Telepon
	return InsertOneDoc(db, col, srt)
}
func InsertTagihan(Isisurat string, status string, db *mongo.Database, col string) (InsertedID interface{}) {
	srt := new(model.Tagihan)
	srt.Isitagihan = isitagihan
	srt.status = status
	return InsertOneDoc(db, col, srt)
}

func InsertDataTagihan(db string, datatagihan DataTagihan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("datatagihan").InsertOne(context.TODO(),datatagihan)
	if err != nil {
		fmt.Printf("InsertDataTagihan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataSudah(db string, datasudah DataSudah) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("datasudah").InsertOne(context.TODO(), datasudah)
	if err != nil {
		fmt.Printf("InsertDataSudah: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataBelum(db string, databelum DataBelum) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("databelum").InsertOne(context.TODO(), databelum)
	if err != nil {
		fmt.Printf("InsertDataBelum: %v\n", err)
	}
	return insertResult.InsertedID
}
}
func GetUserTagihan(telepon string, db *mongo.Database, col string) (data []model.UserTagihan) {
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
func GetTagihan(surat string, db *mongo.Database, col string) (data []model.Tagihan) {
	user := db.Collection(col)
	filter := bson.M{"status": tagihan}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetTagihanUser :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
