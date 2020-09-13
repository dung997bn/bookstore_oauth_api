package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expired time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Access token should not be nil")
	}
	//way1
	if at.AccessToken != "" {
		t.Error("New access token should not have defined access token id")
	}
	//way2
	assert.EqualValues(t, "", at.AccessToken, "New access token should not have defined access token id")

	if at.UserID != 0 {
		t.Error("New access token should not have defined user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("Empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token expiring three hourd from now should not be expired")
	}
}
