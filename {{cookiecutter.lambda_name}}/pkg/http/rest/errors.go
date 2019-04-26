package rest

import (
	"encoding/json"

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
