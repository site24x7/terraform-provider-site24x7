package oauth

import (
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// maxExpiresIn is the maximum value for the expires_in field that is treated
// as ok. Any value larger than that is considered as a broken expiry duration
// and will be handled by tokenSource.
const maxExpiresIn = 1 * time.Hour

// tokenSource only exists due to the fact that the Site24x7 OAuth
// Authorization Server is not compliant with RFC 6749. It returns milliseconds
// in the expires_in field although it must return an expiry in seconds or omit
// the field altogether (see https://tools.ietf.org/html/rfc6749#section-5.1).
type tokenSource struct {
	delegate oauth2.TokenSource
	mu       sync.Mutex
}

// Token implements oauth2.TokenSource.
func (s *tokenSource) Token() (*oauth2.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	token, err := s.delegate.Token()
	if err != nil {
		return nil, err
	}

	// The Site24x7 OAuth Authorization Server hands out access tokens with a
	// lifetime of one hour. Since it reports milliseconds instead of seconds
	// in the expires_in field, the automatic token refresh mechanism is
	// broken.
	// We detect tokens with a wrong expiry based on the fact that it is way
	// beyond one hour from now.
	if time.Until(token.Expiry) > maxExpiresIn {
		// This is a hack which is better than forking golang.org/x/oauth2 and
		// "fixing" it by violating the RFC:
		// Because token is a pointer which is shared with the token source
		// that we are delegating to, we can manipulate it to fix validity
		// checks in our delegates. Subsequent calls to s.delegate.Token() will
		// check the validity based on the fixed expiry date and refresh the
		// token if needed.
		token.Expiry = time.Now().Add(maxExpiresIn)
	}

	return token, nil
}
