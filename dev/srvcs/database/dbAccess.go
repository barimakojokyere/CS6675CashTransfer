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

func InsertIntoDB(dbName string, collectionName string, input interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(dbName)
	collectionObj := database.Collection(collectionName)
	result, err := collectionObj.InsertOne(ctx, input)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)

	return nil
}

func UpdateInDB(dbName string, collectionName string, id string, input interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(dbName)
	collectionObj := database.Collection(collectionName)
	_, err = collectionObj.ReplaceOne(
		ctx,
		bson.M{"_id": id},
		input,
	)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveFromDB(dbName string, collectionName string, id string) (output bson.M, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		return output, err
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(dbName)
	collectionObj := database.Collection(collectionName)
	err = collectionObj.FindOne(ctx, bson.M{"_id": id}).Decode(&output)
	if err != nil {
		return output, err
	}

	return output, nil
}

func RemoveFromDB(dbName string, collectionName string, id string) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.DBURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	fmt.Println("Successfully connected to DB")

	database := client.Database(dbName)
	collectionObj := database.Collection(collectionName)
	result, err := collectionObj.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Println(result.DeletedCount)

	return nil
}
