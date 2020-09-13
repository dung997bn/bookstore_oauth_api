package accesstoken

import (
	"strings"

	"github.com/dung997bn/bookstore_oauth_api/src/utils/errors"
)

//Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

//Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

//NewService func
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

//GetById func
func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
