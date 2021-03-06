package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sky-2002/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// usually these should go in env variables
const connectionString = "---connection--url--of--your--mongoDB--database---"

const dbName = "netflix"
const colName = "watchlist"

// Important
var Collection *mongo.Collection

// connect with mongoDB

// init runs only at the first time
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success !!")

	Collection = client.Database(dbName).Collection(colName)

	// collection instance is ready
	fmt.Println("Collection reference is ready !")
}

// MongoDB helpers

// insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", inserted)
	fmt.Println("Inserted 1 movie in db with id :", inserted.InsertedID)
}

// update 1 movie
func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The result is", result)
	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete one record
func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}

	deleteCount, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted one movie successfully")
	fmt.Println(deleteCount.DeletedCount)
}

// delete all records
func deleteAllMovies() {
	// providing {} means delete all
	deleteCount, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies deleted")
	fmt.Println(deleteCount)
}

// get all records
func getAllMovies() []primitive.M {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie) // passing a reference
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form--urlencode")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

// controller to create a movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form--urlencode")
	w.Header().Set("Allow-Controls-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// marking the movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form--urlencode")
	w.Header().Set("Allow-Controls-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

// delete a movie
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form--urlencode")
	w.Header().Set("Allow-Controls-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Deleted one movie")
	json.NewEncoder(w).Encode(params)
}

// delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form--urlencode")
	w.Header().Set("Allow-Controls-Allow-Methods", "DELETE")

	deleteAllMovies()
	json.NewEncoder(w).Encode("Deleted all movies")
}
