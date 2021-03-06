package rest

import (
	"encoding/json"
	"time"
	
	"github.com/Ryanair/goaws/lambda/apigw"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"

	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/internal/logger"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/{{cookiecutter.lambda_name}}"
)

type createFlightRequest struct {
	CarrierCode      string    `json:"carrierCode"`
	Number           string    `json:"number"`
	Departure        string    `json:"departure"`
	DepatureDateTime time.Time `json:"departureDateTime"`
	Arrival          string    `json:"arrival"`
	ArrivalDateTime  time.Time `json:"arrivalDateTime"`
}

func (r *createFlightRequest) validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.CarrierCode, validation.Required),
		validation.Field(&r.Number, validation.Required),
		validation.Field(&r.Departure, validation.Required),
		validation.Field(&r.DepatureDateTime, validation.Required),
		validation.Field(&r.Arrival, validation.Required),
		validation.Field(&r.ArrivalDateTime, validation.Required),
	)
}

type createFlightResponse struct {
	ID string `json:"id"`
}

type flightCreator interface {
	CreateFlight(flight *{{cookiecutter.lambda_name}}.Flight) ({{cookiecutter.lambda_name}}.ID, error)
}

type CreationHandler struct {
	service flightCreator
}

func NewCreationHandler(service flightCreator) *CreationHandler {
	return &CreationHandler{service: service}
}

func (h *CreationHandler) Handle(req *apigw.Request) (*apigw.Response, error) {
	var out createFlightRequest
	if err := Unmarshal(req, &out); err != nil {
		logger.InvalidRequest(req.Body, errors.Wrap(err, "failed to unmarshal request body"))
		return ResponseInvalidRequest()
	}

	if err := out.validate(); err != nil {
		logger.InvalidRequest(out, errors.Wrap(err, "failed to validate request body"))
		return ResponseValidationError(err)
	}

	id, createErr := createFlight(out, h.service)
	if createErr != nil {
		logger.Errorf("cannot create flight, err: %v", createErr)
		return ResponseInternalError()
	}

	resp := createFlightResponse{ID: id}
	respBody, marshalErr := Marshal(resp)
	if marshalErr != nil {
		logger.InvalidResponse(resp, errors.Wrap(marshalErr, "failed to marshal response body"))
		return ResponseInternalError()
	}

	return apigw.NewResponse(apigw.StatusCreated, respBody), nil
}

func createFlight(req createFlightRequest, service flightCreator) ({{cookiecutter.lambda_name}}.ID, error) {
	flght := {{cookiecutter.lambda_name}}.NewFlight(req.CarrierCode, req.Number, req.Departure, req.Arrival, req.DepatureDateTime, req.ArrivalDateTime)
	id, creationErr := service.CreateFlight(&flght)
	if creationErr != nil {
		logger.Errorf("cannot create flight, err: %v", creationErr)
		return "", creationErr
	}
	logger.Infof("flight %v created successfully", flght)
	return id, creationErr
}

func Unmarshal(req *apigw.Request, out interface{}) error {
	data := []byte(req.Body)
	return json.Unmarshal(data, out)
}

func Marshal(in interface{}) (string, error) {
	bytes, err := json.Marshal(in)
	return string(bytes), err
}
