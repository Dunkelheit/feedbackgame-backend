package error

import (
	"net/http"
)

// Our error codes
const (
	InternalServerError         = 5001
	InvalidAuthenticationTokens = 9001
	InsufficientPermissions     = 9002
)

// Error when something went wrong
// TODO: Consider make private
type Error struct {
	Error          error  `json:"-"`
	Code           int    `json:"code"`
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

// NewError creates a new error
func NewError(code int, err error) *Error {
	e := new(Error)
	e.Code = code
	if err != nil {
		e.Error = err
	}
	e.Defaults()
	return e
}

// Defaults sets the default thingies based on the code
func (e *Error) Defaults() {
	switch e.Code {
	case InternalServerError:
		e.HTTPStatusCode = http.StatusInternalServerError
		e.Message = "An error has occurred"
	case InvalidAuthenticationTokens:
		e.HTTPStatusCode = http.StatusUnauthorized
		e.Message = "Invalid authentication tokens"
	case InsufficientPermissions:
		e.HTTPStatusCode = http.StatusUnauthorized
		e.Message = "You are not allowed to perform this action"
	}
}
