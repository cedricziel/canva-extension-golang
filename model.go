package canva

import (
	"encoding/json"
	"net/http"
)

const (
	TypeSuccess                    = "SUCCESS"
	TypeError                      = "ERROR"
	ErrorCodeInvalidRequest        = "INVALID_REQUEST"
	ErrorCodeInternalError         = "INTERNAL_ERROR"
	ErrorCodeConfigurationRequired = "CONFIGURATION_REQUIRED"
	ErrorCodeForbidden             = "FORBIDDEN"
	ErrorCodeNotFound              = "NOT_FOUND"
	ErrorCodeTimeout               = "TIMEOUT"
)

// Response is a response to a request coming from Canva.com
type Response struct {
	Type      string `json:"type"`
	Resources string `json:"resources,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	URL       string `json:"url,omitempty"`
}

// NewErrorResponse creates a new error response
// see: https://www.canva.com/developers/docs/content-extensions/error-handling/
func NewErrorResponse(errorCode string) Response {
	return Response{Type: TypeError, ErrorCode: errorCode}
}

// NewSuccessResponse creates an empty success response
func NewSuccessResponse() Response {
	return Response{Type: TypeSuccess}
}

// WriteSuccessResponse will write a correct success response
// Canva is really picky about response codes and expects
// a status of 200 and an inline success.
func WriteSuccessResponse(w http.ResponseWriter) {
	json, _ := json.Marshal(NewSuccessResponse())

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(json)
}
