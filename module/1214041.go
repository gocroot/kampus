package module

import (
    "context"
    "fmt"
    "os"
    

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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

func InsertProfile(pendidikan string, username string, bio string, checkin string, biodata Profile) (InsertID interface{}) {
    var listdata ListData
    listdata.Pendidikan = pendidikan
    listdata.Username = username
    listdata.Bio = bio
    listdata.Checkin = checkin
    listdata.Biodata = biodata
    return InsertOneDoc("adorable", "profil", listdata)
}


func CreateProfile(username string, fullName string, email string, password string) (insertedID interface{}) {
	profile := Profile{
		Username: username,
		FullName: fullName,
		Email:    email,
		Password: password,
	}
	return InsertOneDoc("mydb", "profiles", profile)
}

func GetProfileByUsername(username string) (profile Profile) {
	collection := MongoConnect("mydb").Collection("profiles")
	filter := bson.M{"username": username}
	err := collection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("GetProfileByUsername: %v\n", err)
	}
	return profile
}

func UpdateProfilePassword(username string, newPassword string) {
	collection := MongoConnect("mydb").Collection("profiles")
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"password": newPassword}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProfilePassword: %v\n", err)
	}
}

func DeleteProfile(username string) {
	collection := MongoConnect("mydb").Collection("profiles")
	filter := bson.M{"username": username}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Printf("DeleteProfile: %v\n", err)
	}
}

