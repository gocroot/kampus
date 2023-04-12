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

func InsertDashboardq(db *mongo.Database, Username string, Email string, Location string, Orders string, Lastorders string, Totalspent string, Sks string) (InsertedID interface{}) {
	var dashboardz model.Dashboardnye
	dashboardz.Username = Username
	dashboardz.Email = Email
	dashboardz.Location = Location
	dashboardz.Orders = Orders
	dashboardz.Lastorders = Lastorders
	dashboardz.Totalspent = Totalspent
	dashboardz.Sks = Sks
	return InsertOneDoc(db, "data_Dashboard", dashboardz)
}
func InsertMahasiswaq(db *mongo.Database, Nama_mhs string, email_mhs string, gambar_mhs string) (InsertedID interface{}) {
	var mahasiswaz model.Mahasiswanye
	mahasiswaz.Nama_mhs = Nama_mhs
	mahasiswaz.Email_mhs = email_mhs
	mahasiswaz.Gambar_mhs = gambar_mhs
	return InsertOneDoc(db, "data_Mahasiswa", mahasiswaz)
}
func InsertDosenq(db *mongo.Database, nama string, Jabatan string, Noted string, img_dosen string) (InsertedID interface{}) {
	var dosenz model.Dosenye
	dosenz.Nama = nama
	dosenz.Jabatan = Jabatan
	dosenz.Noted = Noted
	dosenz.Img_dosen = img_dosen
	return InsertOneDoc(db, "data_Dosen", dosenz)
}
func InsertAboutnye(db *mongo.Database, pertanyaan string, jawaban string) (InsertedID interface{}) {
	var aboutz model.Aboutnye
	aboutz.Pertanyaan = pertanyaan
	aboutz.Jawaban = jawaban
	return InsertOneDoc(db, "data_about", aboutz)
}
func InsertContacusq(db *mongo.Database, phone_number string, email string) (InsertedID interface{}) {
	var contacuszz model.Contacusnye
	contacuszz.Phone_number = phone_number
	contacuszz.Email = email
	return InsertOneDoc(db, "data_about", contacuszz)
}
func GetDataDashboardz(Username string, db *mongo.Database, col string) (data []model.Dashboardnye) {
	user := db.Collection(col)
	filter := bson.M{"Username": Username}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataDashboard: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataMahasiswaz(Nama_mhs string, db *mongo.Database, col string) (data []model.Mahasiswanye) {
	user := db.Collection(col)
	filter := bson.M{"Mahasiswa": Nama_mhs}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataMahasiswa: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataDosenz(Jabatan string, db *mongo.Database, col string) (data []model.Dosenye) {
	user := db.Collection(col)
	filter := bson.M{"Jabatan": Jabatan}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataDosen: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataAboutz(pertanyaan string, db *mongo.Database, col string) (data []model.Aboutnye) {
	user := db.Collection(col)
	filter := bson.M{"pertanyaan": pertanyaan}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataAbout: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataContacusz(phone_number string, db *mongo.Database, col string) (data []model.Contacusnye) {
	user := db.Collection(col)
	filter := bson.M{"Contacus": phone_number}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataContacus: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
