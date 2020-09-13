package db

import (
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
	return nil, errors.NewInternalServerError("db is not connected")
}
