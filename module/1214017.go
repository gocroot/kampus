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

func InsertDataComp(db *mongo.Database, sistem string, status string, bio model.Userhd) (InsertedID interface{}) {
	var datacomp model.DataComplainhd
	datacomp.Sistemcomp = sistem
	datacomp.Status = status
	datacomp.Biodata = bio
	return InsertOneDoc(db, "data_complain", datacomp)
}

func InsertDataHelper(helper string, username string, nama string, email string, handphone string, db *mongo.Database, col string) (InsertedID interface{}) {
	help := new(model.Helperhd)
	help.Helpid = helper
	help.Username = username
	help.Nama = nama
	help.Email = email
	help.Handphone = handphone
	return InsertOneDoc(db, col, help)
}

func GetDataCompFromStatus(status string, db *mongo.Database, col string) (data model.DataComplainhd) {
	user := db.Collection(col)
	filter := bson.M{"status": status}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

func GetDataAllbyStats(stats string, db *mongo.Database, col string) (data []model.DataComplainhd) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataHelperFromPhone(phone string, db *mongo.Database, col string) (data model.Helperhd) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

func DeleteDataHelper(phone string, db *mongo.Database, col string) (data model.Helperhd) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err, _ := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataHelper : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}

func GetDataCompFromHandphone(phone string, db *mongo.Database, col string) (data model.DataComplainhd) {
	user := db.Collection(col)
	filter := bson.M{"user.handphone": phone}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getKaryawanFromNama: %v\n", err)
	}
	return data
}

func GetDataJumlah(tahun string, db *mongo.Database, col string) (data []model.JumlahComplainhd) {
	user := db.Collection(col)
	filter := bson.M{"tahun": tahun}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataJumlahComplain: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertJumlahComplain(db *mongo.Database, collect string, bulan string, tahun string, jumlah string) (InsertedID interface{}) {
	var jumcomp model.JumlahComplainhd
	jumcomp.Tahun = tahun
	jumcomp.Bulan = bulan
	jumcomp.Jumlah = jumlah
	return InsertOneDoc(db, collect, jumcomp)
}
