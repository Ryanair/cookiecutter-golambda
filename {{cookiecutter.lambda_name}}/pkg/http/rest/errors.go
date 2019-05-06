package rest

import (
	"encoding/json"
	"fmt"

	"github.com/Ryanair/goaws/lambda/apigw"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(msg string) *Error {
	return &Error{
		Message: msg,
	}
}

func (e Error) Marshal() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}

func ResponseError(httpStatus int, msg string) (*apigw.Response, error) {
	errMsg := NewError(msg).Marshal()
	return apigw.NewResponse(httpStatus, errMsg), nil
}

func ResponseInvalidRequest() (*apigw.Response, error) {
	errMsg := NewError("Invalid request body").Marshal()
	return apigw.NewResponse(apigw.StatusBadRequest, errMsg), nil
}

func ResponseValidationError(err error) (*apigw.Response, error) {
	errMsg := NewError(fmt.Sprintf("Invalid request body: %v", err)).Marshal()
	return apigw.NewResponse(apigw.StatusBadRequest, errMsg), nil
}

func ResponseInternalError() (*apigw.Response, error) {
	errMsg := NewError("Internal server error").Marshal()
	return apigw.NewResponse(apigw.StatusInternalServerError, errMsg), nil
}
