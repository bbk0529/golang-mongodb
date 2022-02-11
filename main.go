package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://myUserAdmin:abc123@localhost:27017"

type result2 struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"type,omitempty"`
	Rating int                `bson:"rating,omitempty"`
}

type metadata struct {
	SensorId int    `bson:"sensorId, omitempty"`
	Type     string `bson:"type, omitempty"`
}

type result struct {
	Timestamp time.Time          `bson: "timestamp, omitempty"`
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Metadata  metadata           `bson: "metadata, omiempty`
	Temp      int                `bson: "temp, omitempty"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	errTest(err)
	collection := client.Database("timeseries").Collection("weather")
	filter := bson.D{}

	results := readDatabase(ctx, collection, filter)
	for _, v := range results {
		fmt.Println(v)
		// fmt.Println(v.timestamp, v.Metadata.SensorId, v.Metadata.Type)
	}
}

func readDatabase(ctx context.Context, collection *mongo.Collection, filter bson.D) []result {
	var results []result
	cur, err := collection.Find(context.Background(), bson.D{})
	errTest(err)
	if err := cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	return results
}

func errTest(err error) {
	if err != nil {
		panic(err)
	}
}
