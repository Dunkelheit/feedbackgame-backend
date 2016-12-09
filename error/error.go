package error

import (
	"fmt"
	"net/http"
)

// Our error codes
const (
	InvalidRequest              = 4001
	EntityNotFound              = 4002
	InternalServerError         = 5001
	InvalidAuthenticationTokens = 9001
	InsufficientPermissions     = 9002
)

// APIError when something went wrong
type APIError struct {
	Code           int    `json:"code"`
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

func (err *APIError) Error() string {
	return fmt.Sprintf("[ERR %d] (%d) %s", err.Code, err.HTTPStatusCode, err.Message)
}

// ForCode creates a new error for a specific code
func ForCode(code int, err error) (apiError *APIError) {
	apiError = new(APIError)
	apiError.Code = code
	apiError.defaults()
	return
}

// Defaults sets the default thingies based on the code
func (err *APIError) defaults() {
	switch err.Code {
	case InvalidRequest:
		err.HTTPStatusCode = http.StatusBadRequest
		err.Message = "Your request is invalid"
	case EntityNotFound:
		err.HTTPStatusCode = http.StatusNotFound
		err.Message = "Entity not found"
	case InternalServerError:
		err.HTTPStatusCode = http.StatusInternalServerError
		err.Message = "An error has occurred"
	case InvalidAuthenticationTokens:
		err.HTTPStatusCode = http.StatusUnauthorized
		err.Message = "Invalid authentication tokens"
	case InsufficientPermissions:
		err.HTTPStatusCode = http.StatusUnauthorized
		err.Message = "You are not allowed to perform this action"
	}
}
