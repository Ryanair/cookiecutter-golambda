package {{cookiecutter.lambda_name}}

import (
	"time"

	"github.com/rs/xid"
)

type ID = string

type Flight struct {
	ID                ID
	CarrierCode       string
	Number            string
	Departure         string
	DepartureDateTime time.Time
	Arrival           string
	ArrivalDateTime   time.Time
}

func NewFlight(carrierCode, number, departure, arrival string, departureDateTime, arrivalDateTime time.Time) Flight {
	guid := xid.New()
	return Flight{
		ID:                guid.String(),
		CarrierCode:       carrierCode,
		Number:            number,
		Departure:         departure,
		DepartureDateTime: departureDateTime,
		Arrival:           arrival,
		ArrivalDateTime:   arrivalDateTime,
	}
}
