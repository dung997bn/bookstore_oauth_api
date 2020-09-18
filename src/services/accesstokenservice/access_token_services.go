package accesstokenservice

import (
	"strings"

	"github.com/dung997bn/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/dung997bn/bookstore_oauth_api/src/repository/db"
	"github.com/dung997bn/bookstore_oauth_api/src/repository/rest"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

//Service type
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *resterrors.RestErr)
	Create(accesstoken.AccessTokenRequest) (*accesstoken.AccessToken, *resterrors.RestErr)
	UpdateExpirationTime(accesstoken.AccessToken) *resterrors.RestErr
}

type service struct {
	restUserRepo rest.RestUsersRepository
	dbRepo       db.DbRepository
}

//NewService func
func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUserRepo: usersRepo,
		dbRepo:       dbRepo,
	}
}

func (s *service) GetByID(accessTokenID string) (*accesstoken.AccessToken, *resterrors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, resterrors.NewBadRequestError("Invalid accesstoken id")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request accesstoken.AccessTokenRequest) (*accesstoken.AccessToken, *resterrors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := accesstoken.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at accesstoken.AccessToken) *resterrors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
