package db

import (
	"github.com/dung997bn/bookstore_oauth_api/src/clients/cassandra"
	"github.com/dung997bn/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/dung997bn/bookstore_oauth_api/src/utils/errors"
)

//DbRepository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

//New func
func New() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	//TODO: implement get access token from Cassandra
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}

	defer session.Close()

	return nil, errors.NewInternalServerError("db is not connected")
}
