package errors

// StatusError provides the HTTP status code in addition to the error message.
type StatusError interface {
	error
	StatusCode() int
}

// ExtendedStatusError provides HTTP status and additional error information.
type ExtendedStatusError interface {
	StatusError
	ErrorCode() int
	ErrorInfo() map[string]interface{}
}

type statusError struct {
	statusCode int
	message    string
}

// NewStatusError creates a new StatusError with statusCode and message.
func NewStatusError(statusCode int, message string) StatusError {
	return &statusError{
		statusCode: statusCode,
		message:    message,
	}
}

// StatusCode implements StatusError.
func (e *statusError) StatusCode() int {
	return e.statusCode
}

// Error implements error.
func (e *statusError) Error() string {
	return e.message
}

type extendedStatusError struct {
	StatusError
	errorCode int
	errorInfo map[string]interface{}
}

// NewExtendedStatusError creates a new ExtendedStatusError with statusCode,
// message and additional error information.
func NewExtendedStatusError(statusCode int, message string, errorCode int, errorInfo map[string]interface{}) StatusError {
	return &extendedStatusError{
		StatusError: NewStatusError(statusCode, message),
		errorCode:   errorCode,
		errorInfo:   errorInfo,
	}
}

// ErrorCode implements ExtendedStatusError.
func (e *extendedStatusError) ErrorCode() int {
	return e.errorCode
}

// ErrorInfo implements ExtendedStatusError.
func (e *extendedStatusError) ErrorInfo() map[string]interface{} {
	return e.errorInfo
}
