package db

import (
	"github.com/dung997bn/bookstore_oauth_api/src/clients/cassandra"
	"github.com/dung997bn/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/dung997bn/bookstore_oauth_api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, client_id, user_id, exprires from access_tokens where access_token = ? ;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, client_id, user_id, exprires) VALUES (?, ?, ?, ?) ;"
	queryUpdateExpires     = "UPDATE access_tokens set expires=? WHERE  access_token = ?"
)

//DbRepository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

//NewDbRepository func
func NewDbRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	//TODO: implement get access token from Cassandra
	// session, err := cassandra.GetSession()
	// if err != nil {
	// 	return nil, errors.NewInternalServerError(err.Error())
	// }
	// defer session.Close()

	var result accesstoken.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.ClientID, &result.UserID, &result.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (db *dbRepository) Create(ac accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, ac.AccessToken, ac.ClientID, ac.UserID, ac.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(ac accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, ac.Expires, ac.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
