package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://votingapp:Ranjan11082@clustervote.ccp4sos.mongodb.net/banking_way"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init(){
	//CLIENT OPTION
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instanse
	fmt.Println("Collection is ready")
}

//INSERT 1 RECORD--->

func insertOneMovie(movie model.Netflix) error {
	if movie.ID.IsZero() {
		movie.ID = primitive.NewObjectID()
	}
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("duplicate key error: %v", err)
		}
		return err
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
	return nil
}


//update 1 record

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with the deleted count", deleteCount)
}

//delete all movies==>
func deleteAllMovies() int64{
	// filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("No of movies deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//GET ALL movie COLLECTIONS fromm DB-->
func getAllMovies()[]primitive.M{
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

//Actual  controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	err := insertOneMovie(movie)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(movie)
}


func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}																																																																			

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}