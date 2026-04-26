package db

import (
	"context"
	"film-web-service/model"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const uri = "mongodb://admin:admin@localhost:27017"

type Mongo struct {
	Client *mongo.Client
}

func Init() (*Mongo, error) {
	_, _ = context.WithTimeout(context.Background(), 30*time.Second)
	client, error := mongo.Connect(options.Client().ApplyURI(uri))
	if error != nil {
		return nil, error
	}
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connect to mongo server \n")

	return &Mongo{client}, nil
}

func (mongo *Mongo) Close() {
	err := mongo.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (mongo *Mongo) FindFilmId(objectId string) ([]model.Film, error) {
	var films []model.Film

	objId, err := bson.ObjectIDFromHex(objectId)
	if err != nil {
		return nil, err
	}
	cur, err := mongo.Client.Database("mflix").Collection("movies").Find(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		err := cur.All(context.Background(), &films)
		if err != nil {
			return nil, err
		}
	}

	return films, nil
}

func (mongo *Mongo) PostFilm(data model.Film) error {
	_, err := mongo.Client.Database("mflix").Collection("movies").InsertOne(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *Mongo) PutFilmId(objectId string, data model.Film) error {
	objId, err := bson.ObjectIDFromHex(objectId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": data}

	_, err = mongo.Client.Database("mflix").Collection("movies").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *Mongo) DeleteFilmId(objectId string) ([]model.Film, error) {
	objId, err := bson.ObjectIDFromHex(objectId)

	data, err := mongo.FindFilmId(objectId)
	if err != nil {
		return nil, err
	}
	_, err = mongo.Client.Database("mflix").Collection("movies").DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (mongo *Mongo) FindFilms(count int64) ([]model.Film, error) {
	var films []model.Film

	cur, err := mongo.Client.Database("mflix").Collection("movies").Find(context.Background(), bson.M{}, options.Find().SetLimit(count))
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		err := cur.All(context.Background(), &films)
		if err != nil {
			return nil, err
		}
	}

	return films, nil
}
