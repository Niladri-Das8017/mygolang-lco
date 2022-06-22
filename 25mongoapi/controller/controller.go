package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niladridas/mongoAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Niladri:liza1998@cluster0.wvsw2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority" //env variable
const dbName = "netflix"
const colName = "watchList"

//MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connct to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collectoin instance is ready")
}

// MongoDB helpers - folder

//insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err) //More controlle version of panic
	}

	fmt.Println("Inserted 1 movie in db wwith id: ", inserted.InsertedID)
}

//update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOnemovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete  count: ", deleteCount)
}

//delete all records from mongodb
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number  of  movies  deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//gt all record
func getAllMovies() []primitive.M {
	curor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	//array of movies
	var movies []primitive.M //since we are using "M" primitive package

	for curor.Next(context.Background()) {

		var movie bson.M
		//if cursor.Decode(&movie) worked,we have movie fieled up, else we will have an err.
		err := curor.Decode(&movie) //whenever we decode, we pass on a referrence like, "if you decode use my structure to decode that"
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer curor.Close(context.Background())
	return movies
}

// ACTUAL CONTROLLERS !!!

//get all record
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")

	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

//Create record
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	//w.Header().Set("Allow-Control-Allow-Methods", "POST") //only allow POST

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	//insert  1 movie
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

//update record
func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT") //only allow POSTs

	params := mux.Vars(r)

	//update 1 record
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//delete 1 record
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow POSTs

	params := mux.Vars(r)
	//calling handler
	deleteOnemovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

//delete all record
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcaton/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow POSTs

	//calling handeler
	count := deleteAllMovie()

	json.NewEncoder(w).Encode(count)
}
