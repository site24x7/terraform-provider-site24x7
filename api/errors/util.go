package errors

import "net/http"

// IsStatusError returns true if err is of type StatusError.
func IsStatusError(err error) bool {
	_, ok := err.(StatusError)
	return ok
}

// IsExtendedStatusError returns true if err is of type ExtendedStatusError.
func IsExtendedStatusError(err error) bool {
	_, ok := err.(ExtendedStatusError)
	return ok
}

// HasStatusCode returns true if err has the given status code. If err is not
// of type StatusError, this will always return false.
func HasStatusCode(err error, code int) bool {
	if statusErr, ok := err.(StatusError); ok {
		return statusErr.StatusCode() == code
	}

	return false
}

// IsNotFound is a convenience method for checking if an error is a StatusError
// with status code 404.
func IsNotFound(err error) bool {
	return HasStatusCode(err, http.StatusNotFound)
}
