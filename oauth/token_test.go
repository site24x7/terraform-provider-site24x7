package oauth

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestTokenSource_Token(t *testing.T) {
	brokenToken := &oauth2.Token{
		AccessToken: "foobar",
		Expiry:      time.Now().Add(maxExpiresIn * 1000), // off by one magnitude
	}

	ts := &tokenSource{
		delegate: oauth2.StaticTokenSource(brokenToken),
	}

	// first call fixes the expiry time
	token, err := ts.Token()
	require.NoError(t, err)

	expiry := token.Expiry

	assert.True(t, time.Until(expiry) <= maxExpiresIn)
	assert.Equal(t, token, brokenToken)

	// subsequent calls do not update the exipry time
	token, err = ts.Token()
	require.NoError(t, err)

	assert.Equal(t, expiry, token.Expiry)
}

func TestTokenSource_Token_expired(t *testing.T) {
	expiredToken := &oauth2.Token{
		AccessToken: "foobar",
		Expiry:      time.Now(),
	}

	ts := &tokenSource{
		delegate: oauth2.StaticTokenSource(expiredToken),
	}

	token, err := ts.Token()
	require.NoError(t, err)

	assert.False(t, token.Valid())
}

type badTokenSource struct{}

func (*badTokenSource) Token() (*oauth2.Token, error) {
	return nil, errors.New("whoops")
}

func TestTokenSource_Token_error(t *testing.T) {
	ts := &tokenSource{
		delegate: &badTokenSource{},
	}

	_, err := ts.Token()
	require.Error(t, err)
	assert.Equal(t, errors.New("whoops"), err)
}
