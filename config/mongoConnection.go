package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName="goTest";
const colName="watchlist";
const connectionString="mongodb://localhost:27017";


var USER *mongo.Collection;
var Mongo *mongo.Database;

func MongoConnection(){
	clientOption :=options.Client().ApplyURI(connectionString)

	//connect to mongo db

	client,err :=mongo.Connect(context.TODO(), clientOption)

	if(err!=nil){
		log.Fatal(err)

	}

	fmt.Println("Mongo db connected")
	Mongo=client.Database(dbName)
	USER=client.Database(dbName).Collection("users");

	fmt.Println("Connection instance ready")

}
