package config

import (
	"auth_fiber/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName="goTest";
const colName="watchlist";
const connectionString="mongodb://localhost:27017";


var collection *mongo.Collection;

func MongoConnection(){
	
	clientOption :=options.Client().ApplyURI(connectionString)

	//connect to mongo db

	client,err :=mongo.Connect(context.TODO(), clientOption)

	if(err!=nil){
		log.Fatal(err)

	}

	fmt.Println("Mongo db connected")

	collection=client.Database(dbName).Collection(colName);

	fmt.Println("Connection instance ready")

}


//inser record

func insertOneMovie(movie models.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if(err !=nil){
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id" , inserted.InsertedID);
}

//update 1 record

func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId);

	filter := bson.M{
		"_id":id,
	}
	update :=bson.M{"$set":bson.M{
		"watched":true,
	}}

	result,err := collection.UpdateOne(context.Background(),filter,update);

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count",result.ModifiedCount)

}