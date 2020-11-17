package app

import (
	"bookstore-ouath-api/src/clients/cassandra"
	"bookstore-ouath-api/src/domain/access_token"
	"bookstore-ouath-api/src/http"
	"bookstore-ouath-api/src/repository/db"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// try connecting to cassandra on application start
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()
	repository := db.NewRepository()
	acTokenService := access_token.NewService(repository)
	handler := http.NewAccessTokenHandler(acTokenService)

	router.GET("/oauth/access_token/:access_token_id", handler.GetById)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
