package riziq

import (
	"context"
	"fmt"
	"os"

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

func InsertPelanggan(nm string, alm string, tlp string, em string) (InsertedID interface{}) {
	var pelanggan Pelanggan
	pelanggan.Nama = nm
	pelanggan.Alamat = alm
	pelanggan.NoTelepon = tlp
	pelanggan.Email = em

	return InsertOneDoc("dbpelanggan", "pelanggan", pelanggan)
}

func InsertTagihan(tgh string, tth string, sp string) (InsertedID interface{}) {
	var tagihan Tagihan
	tagihan.TanggalTagihan = tgh
	tagihan.TotalTagihan = tth
	tagihan.StatusPembayaran = sp

	return InsertOneDoc("dbtagihan", "tagihan", tagihan)
}

func InsertPembayaran(tby string, jby string, mby string) (InsertedID interface{}) {
	var pembayaran Pembayaran
	pembayaran.TanggalPembayaran = tby
	pembayaran.JumlahPembayaran = jby
	pembayaran.MetodePembayaran = mby

	return InsertOneDoc("dbpembayaran", "pembayaran", pembayaran)
}

func InsertProduk(npr string, hpr string) (InsertedID interface{}) {
	var produk Produk
	produk.NamaProduk = npr
	produk.HargaProduk = hpr

	return InsertOneDoc("dbProduk", "produk", produk)
}

func InsertAbout(satu string, dua string) (InsertedID interface{}) {
	var about About
	about.IsiSatu = satu
	about.IsiDua = dua

	return InsertOneDoc("dbabout", "about", about)
}
