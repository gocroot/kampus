package module

import (
	"context"
	"fmt"
	"os"

	"github.com/gocroot/kampus/model"
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

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataComp(sistem string, status string, bio model.User) (InsertedID interface{}) {
	var datacomp model.DataComplain
	datacomp.Sistemcomp = sistem
	datacomp.Status = status
	datacomp.Biodata = bio
	return InsertOneDoc("HelpdeskData", "data_complain", datacomp)
}

func InsertDataHelper(helper string, username string, nama string, email string, handphone string) (InsertedID interface{}) {
	help := new(model.Helper)
	help.Helpid = helper
	help.Username = username
	help.Nama = nama
	help.Email = email
	help.Handphone = handphone
	return InsertOneDoc("HelpdeskData", "helperdata", help)
}

func GetDataCompFromStatus(status string) (data model.DataComplain) {
	user := MongoConnect("HelpdeskData").Collection("data_complain")
	filter := bson.M{"status": status}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

//func GetDataAllbyStats(stats string) (data model.DataComplain) {
//	user := MongoConnect("HelpdeskData").Collection("data_complain")
//	filter := bson.M{"status": stats}
//	err, _ := user.Find(context.TODO(), filter)
//	if err != nil {
//		fmt.Println("GetALLData :", err)
//	}
//	return data
//}

func GetDataHelperFromPhone(phone string) (data model.Helper) {
	user := MongoConnect("HelpdeskData").Collection("helperdata")
	filter := bson.M{"handphone": phone}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

func DeleteDataHelper(phone string) (data model.Helper) {
	user := MongoConnect("HelpdeskData").Collection("helperdata")
	filter := bson.M{"handphone": phone}
	err, _ := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataHelper : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
