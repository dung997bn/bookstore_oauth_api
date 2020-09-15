package app

import (
	"github.com/dung997bn/bookstore_oauth_api/src/clients/cassandra"
	"github.com/dung997bn/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/dung997bn/bookstore_oauth_api/src/http"
	"github.com/dung997bn/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication func
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	dbRepository := db.New()
	atService := accesstoken.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	//route
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	//run
	router.Run(":8080")
}
