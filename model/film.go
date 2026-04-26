package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Film struct {
	ID               bson.ObjectID `bson:"_id"`
	Awards           Awards        `bson:"awards"`
	Cast             []string      `bson:"cast"`
	Countries        []string      `bson:"countries"`
	Directors        []string      `bson:"directors"`
	Fullplot         string        `bson:"fullplot"`
	Genres           []string      `bson:"genres"`
	Imdb             Imdb          `bson:"imdb"`
	Languages        []string      `bson:"languages"`
	Lastupdated      string        `bson:"lastupdated"`
	Metacritic       int           `bson:"metacritic"`
	NumMflixComments int           `bson:"num_mflix_comments"`
	Plot             string        `bson:"plot"`
	Poster           string        `bson:"poster"`
	Rated            string        `bson:"rated"`
	Released         bson.DateTime `bson:"released"`
	Runtime          int           `bson:"runtime"`
	Title            string        `bson:"title"`
	Tomatoes         Tomatoes      `bson:"tomatoes"`
	Type             string        `bson:"type"`
	Writers          []string      `bson:"writers"`
	Year             int           `bson:"year"`
}

type Awards struct {
	Wins        int    `bson:"wins"`
	Nominations int    `bson:"nominations"`
	Text        string `bson:"text"`
}
type Imdb struct {
	Rating float64 `bson:"rating"`
	Votes  int     `bson:"votes"`
	ID     int     `bson:"id"`
}
type Viewer struct {
	Rating     float64 `bson:"rating"`
	NumReviews int     `bson:"numReviews"`
	Meter      int     `bson:"meter"`
}

type Tomatoes struct {
	Viewer      Viewer        `bson:"viewer"`
	LastUpdated bson.DateTime `bson:"lastUpdated"`
}
