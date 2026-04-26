package api

import (
	"encoding/json"
	mdb "film-web-service/db/mongo"
	rdb "film-web-service/db/redis"
	"film-web-service/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFilmId(gin *gin.Context) {
	p := gin.Param("id")

	mongoClient, err := mdb.Init()
	defer mongoClient.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := mongoClient.FindFilmId(p)
	if err != nil {
		return
	}
	gin.IndentedJSON(http.StatusOK, data)
}

func PostFilm(gin *gin.Context) {
	var data model.Film

	err := gin.BindJSON(&data)
	if err != nil {
		return
	}

	mongoClient, err := mdb.Init()
	defer mongoClient.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.PostFilm(data)
	if err != nil {
		return
	}
	gin.IndentedJSON(http.StatusCreated, data)
}

func PutFilmId(gin *gin.Context) {
	p := gin.Param("id")
	var data model.Film

	err := gin.BindJSON(&data)
	if err != nil {
		return
	}

	mongoClient, err := mdb.Init()
	defer mongoClient.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.PutFilmId(p, data)
	if err != nil {
		return
	}
	gin.IndentedJSON(http.StatusOK, data)
}

func DeleteFilmId(gin *gin.Context) {
	p := gin.Param("id")

	mongoClient, err := mdb.Init()
	defer mongoClient.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := mongoClient.DeleteFilmId(p)
	if err != nil {
		return
	}
	gin.JSON(http.StatusOK, data)
}

func SendMessages(ctx *gin.Context) {
	type Messages struct {
		Limit int64 `json:"limit"`
	}
	var messages Messages

	err := ctx.BindJSON(&messages)
	if err != nil {
		return
	}
	mongoClient, err := mdb.Init()
	defer mongoClient.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := mongoClient.FindFilms(messages.Limit)
	if err != nil {
		return
	}
	redisClient, err := rdb.Init()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	redisClient.Producer("queue", jsonData)
	ctx.IndentedJSON(http.StatusOK, data)
}
