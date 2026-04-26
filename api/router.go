package api

import (
	api "film-web-service/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/films/:id", api.GetFilmId)
	r.POST("/film", api.PostFilm)
	r.PUT("/films/:id", api.PutFilmId)
	r.DELETE("films/:id", api.DeleteFilmId)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
