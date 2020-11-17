package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConstantInAccessToken(t *testing.T) {
	assert.EqualValues(t, expirationTime, 24, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.Empty(t, at.AccessToken, "New access token should not have defined access token id")

	assert.True(t, at.UserId == 0, "New access token should not have associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "Access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "Access token expire after 3 hours")
}
