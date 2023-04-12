package module

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// var MongoString string = os.Getenv("MONGOSTRING")

// func MongoConnect(dbname string) (db *mongo.Database) {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
// 	if err != nil {
// 		fmt.Printf("MongoConnect: %v\n", err)
// 	}
// 	return client.Database(dbname)
// }

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPelanggantagih(db *mongo.Database, nm string, alm string, tlp string, em string) (InsertedID interface{}) {
	var pelanggan model.Pelanggantagih
	pelanggan.Nama = nm
	pelanggan.Alamat = alm
	pelanggan.NoTelepon = tlp
	pelanggan.Email = em

	return InsertOneDoc(db, "pelanggan", pelanggan)
}

func InsertTagihan(db *mongo.Database, tgh string, tth string, sp string) (InsertedID interface{}) {
	var tagihan model.Tagihan
	tagihan.TanggalTagihan = tgh
	tagihan.TotalTagihan = tth
	tagihan.StatusPembayaran = sp

	return InsertOneDoc(db, "tagihan", tagihan)
}

func InsertDataPembayarantagih(db *mongo.Database, tby string, jby string, mby string) (InsertedID interface{}) {
	var pembayaran model.Pembayarantagih
	pembayaran.TanggalPembayaran = tby
	pembayaran.JumlahPembayaran = jby
	pembayaran.MetodePembayaran = mby

	return InsertOneDoc(db, "Pembayaran", pembayaran)
}

func InserDatatProduktagih(db *mongo.Database, npr string, hpr string) (InsertedID interface{}) {
	var produk model.Produktagih
	produk.NamaProduk = npr
	produk.HargaProduk = hpr

	return InsertOneDoc(db, "produk", produk)
}

func InsertDataAbouttagih(db *mongo.Database, satu string, dua string) (InsertedID interface{}) {
	var about model.Abouttagih
	about.IsiSatu = satu
	about.IsiDua = dua

	return InsertOneDoc(db, "about", about)
}
