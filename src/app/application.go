package app

import (
	"github.com/dung997bn/bookstore_oauth_api/src/http"
	"github.com/dung997bn/bookstore_oauth_api/src/repository/db"
	"github.com/dung997bn/bookstore_oauth_api/src/repository/rest"
	"github.com/dung997bn/bookstore_oauth_api/src/services/accesstokenservice"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication func
func StartApplication() {
	atService := accesstokenservice.NewService(rest.NewRestUsersRepository(), db.NewDbRepository())
	atHandler := http.NewAccessTokenHandler(atService)
	//route
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	//run
	router.Run(":8080")
}
