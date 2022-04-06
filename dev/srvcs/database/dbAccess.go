package database

import (
	"cashtransfer/dev/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertIntoDB(collectionName string, input interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		panic("Error connecting to Database")
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(utils.DBNAME)
	collectionObj := database.Collection(collectionName)
	result, err := collectionObj.InsertOne(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func UpdateInDB(collectionName string, id string, input interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		panic("Error connecting to Database")
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(utils.DBNAME)
	collectionObj := database.Collection(collectionName)
	result, err := collectionObj.ReplaceOne(
		ctx,
		bson.M{"_id": id},
		input,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ModifiedCount)
}

func RetrieveFromDB(collectionName string, id string) (output bson.M) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		panic("Error connecting to Database")
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(utils.DBNAME)
	collectionObj := database.Collection(collectionName)
	err = collectionObj.FindOne(ctx, bson.M{"_id": id}).Decode(&output)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	return output
}

func RemoveFromDB(collectionName string, id string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		panic("Error connecting to Database")
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(utils.DBNAME)
	collectionObj := database.Collection(collectionName)
	result, err := collectionObj.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.DeletedCount)
}
