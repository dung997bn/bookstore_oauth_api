package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/dung997bn/bookstore_oauth_api/src/utils/cryptoutils"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

//AccessTokenRequest type
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	//Used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	//Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

//Validate func
func (at *AccessTokenRequest) Validate() *resterrors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return resterrors.NewBadRequestError("Invalid grant type parameter")
	}
	return nil
}

//AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"exprires"`
}

//Validate func
func (at *AccessToken) Validate() *resterrors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return resterrors.NewBadRequestError("Invalid token id")
	}
	if at.UserID <= 0 {
		return resterrors.NewBadRequestError("Invalid user id")
	}
	if at.ClientID <= 0 {
		return resterrors.NewBadRequestError("Invalid client id")
	}
	if at.Expires <= 0 {
		return resterrors.NewBadRequestError("Invalid expriration time")
	}
	return nil
}

//GetNewAccessToken func
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

//IsExpired func
func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	return time.Unix(at.Expires, 0).Before(now)
}

//Generate func
func (at *AccessToken) Generate() {
	at.AccessToken = cryptoutils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
