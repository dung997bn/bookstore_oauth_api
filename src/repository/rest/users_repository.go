package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dung997bn/bookstore_oauth_api/src/domain/users"
	"github.com/dung997bn/bookstore_oauth_api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://bookstore.com/users/login",
		Timeout: 100 * time.Millisecond,
	}
)

//RestUsersRepository interface
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

//NewRepository func
func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	requestBody := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(requestBody)
	fmt.Println(string(bytes))

	response := usersRestClient.Post("/users/login", requestBody)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying unmarshal users response")
	}
	return &user, nil
}
