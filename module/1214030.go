package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DataAkreditas(Perguruan_Tinggi string, Program_studi string, Peringkat string, Status string, db *mongo.Database, col string) (InsertedID interface{}) {
	DataLam := new(model.DataAkreditas)
	DataLam.Perguruan_Tinggi = Perguruan_Tinggi
	DataLam.Program_studi = Program_studi
	DataLam.Peringkat = Peringkat
	return InsertOneDoc(db, col, DataLam)
}
func DataProgramStudi(Program_studi string, program string, db *mongo.Database, col string) (InsertedID interface{}) {
	DataLam := new(model.DataProgramStudi)
	DataLam.Program_studi = Program_studi
	DataLam.Program = program
	return InsertOneDoc(db, col, DataLam)
}
func Profile(Isi_satu string, db *mongo.Database, col string) (InsertedID interface{}) {
	DataLam := new(model.Profile)
	DataLam.Isi_satu = Isi_satu
	return InsertOneDoc(db, col, DataLam)
}
func GetDataPeringkat(peringkat string, db *mongo.Database, col string) (data []model.DataAkreditas) {
	user := db.Collection(col)
	filter := bson.M{"peringkat": peringkat}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPeringkat :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataStatus(status string, db *mongo.Database, col string) (data []model.DataAkreditas) {
	user := db.Collection(col)
	filter := bson.M{"status": status}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAkreditas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
