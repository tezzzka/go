package db

import (
	// "encoding/json"
	// "reflect"
	
	"log"
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func establishment() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("numbers")
	return collection
}
func JsonStructToMongoDB() {
	collection:=establishment()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.D{{"Id", 0}, {"Name", "sur"},{"Description", "Lorem ipsum"},{"Price", 0.1}})
	if err!=nil {
		panic(err)
	}
	// fmt.Println(res)
}

func GetAllRecords() string {
	ctx, cancel:= context.WithTimeout(context.Background(), 30*time.Second)
	collection:=establishment()
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { fmt.Println(err) }
	defer cur.Close(ctx)

	res:=""

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil { fmt.Println(err) }
		// do something with result....
		a:=fmt.Sprint(result)
		res+=a
		// fmt.Println(reflect.TypeOf(result))
		// fmt.Println(reflect.TypeOf(a))
		if err != nil { fmt.Println(err) }
	 }
	
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(res)
	return res
	// fmt.Println(reflect.TypeOf(cur))
	
}

func GetItem(id int) string {
	// fmt.Println(id)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
    	log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
    	log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("testing").Collection("numbers")

	type Record struct {
		Id int32
		Name string
		Description string
		Price float64
	}

	filter := bson.D{{"Id", id}}
	
	var result Record

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Found a single document: %+v\n", result)

	return fmt.Sprint(result)

}

func DelById(id int) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
    	log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
    	log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to MongoDB!")

	collection := client.Database("testing").Collection("numbers")

	type Record struct {
		Id int32
		Name string
		Description string
		Price float64
	}

	filter := bson.D{{"Id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

// func (DeLItemById string) {

// }

// func PostItem() {

// }

