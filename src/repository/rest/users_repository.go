package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dung997bn/bookstore_oauth_api/src/domain/users"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://127.0.0.1:8081",
		Timeout: 100 * time.Millisecond,
	}
)

//RestUsersRepository interface
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, resterrors.RestErr)
}

type usersRepository struct{}

//NewRestUsersRepository func
func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, resterrors.RestErr) {
	requestBody := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(requestBody)
	fmt.Println(string(bytes))

	response := usersRestClient.Post("/users/login", requestBody)
	if response == nil || response.Response == nil {
		return nil, resterrors.NewInternalServerError("Invalid restclient response when trying to login user", errors.New("server error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			fmt.Println(err)
			return nil, resterrors.NewInternalServerError("Invalid error interface when trying to login user", errors.New("server error"))
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error when trying unmarshal users response", errors.New("server error"))
	}
	return &user, nil
}
