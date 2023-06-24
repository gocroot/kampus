package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPelanggantagih(Nama string, Alamat string, NoTelepon string, Email string, db *mongo.Database, col string) (InsertedID interface{}) {
	plg := new(model.Pelanggantagih)
	plg.Nama = Nama
	plg.Alamat = Alamat
	plg.NoTelepon = NoTelepon
	plg.Email = Email
	return InsertOneDoc(db, col, plg)
}

func GetUserNama(nama string, db *mongo.Database, col string) (data []model.Pelanggantagih) {
	user := db.Collection(col)
	filter := bson.M{"nama": nama}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetUserNama :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
