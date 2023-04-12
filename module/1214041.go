package module
import (
    "context"
    "fmt"
    

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
   
)

type Profile struct {
	Username string `bson:"username"`
	FullName string `bson:"full_name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type ListData struct{
	Pendidikan  string    `bson:"pendidikan"`
	Bio         string    `bson:"bio"`
	Username    string    `bson:"username"`
	Checkin     string    `bson:"checkin"`
	Biodata     Profile   `bson:"biodata"`
}


func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProfile(pendidikan string, username string, bio string, checkin string, biodata Profile, db *mongo.Database) (InsertID interface{}) {
    var listdata ListData
    listdata.Pendidikan = pendidikan
    listdata.Username = username
    listdata.Bio = bio
    listdata.Checkin = checkin
    listdata.Biodata = biodata
    return InsertOneDoc(db, "profil", listdata)
}


func CreateProfile(username string, fullName string, email string, password string, db *mongo.Database, input string) (insertedID interface{}) {
	profile := Profile{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: password,
	}
	return InsertOneDoc(db, input, profile)
}

func GetProfileByUsername(username string, db *mongo.Database, input string) (profile Profile) {
	collection := db.Collection(input)
	filter := bson.M{"username": username}
	err := collection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("GetProfileByUsername: %v\n", err)
	}
	return profile
}

func GetDataProfFromStatus(status string, db *mongo.Database, input string) (profile Profile) {
	collection := db.Collection(input)
	filter := bson.M{"status": status}
	err := collection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("GetProfileByUsername: %v\n", err)
	}
	return profile
}

func UpdateProfilePassword(username string, newPassword string, input string, db *mongo.Database) {
	collection := db.Collection(input)
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"password": newPassword}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProfilePassword: %v\n", err)
	}
}

func DeleteProfile(username string, db *mongo.Database, input string) {
    collection := db.Collection(input)
    filter := bson.M{"username": username}
    _, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        fmt.Printf("DeleteProfile: %v\n", err)
    }
    fmt.Println("Berhasil menghapus data Profile.")
}

//upload
//
