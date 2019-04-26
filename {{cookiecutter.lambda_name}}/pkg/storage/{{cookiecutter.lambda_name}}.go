package storage

import (
	"time"
)

type Flight struct {
	ID                string    `dynamodbav:"id"`
	CarrierCode       string    `dynamodbav:"carrierCode"`
	Number            string    `dynamodbav:"number"`
	Departure         string    `dynamodbav:"departure"`
	DepartureDateTime time.Time `dynamodbav:"departureDateTime"`
	Arrival           string    `dynamodbav:"arrival"`
	ArrivalDateTime   time.Time `dynamodbav:"arrivalDateTime"`
}

func NewFlight(id, carrierCode, number, departure, arrival string, departureDateTime, arrivalDateTime time.Time) Flight {
	return Flight{
		ID:                id,
		CarrierCode:       carrierCode,
		Number:            number,
		Departure:         departure,
		DepartureDateTime: departureDateTime,
		Arrival:           arrival,
		ArrivalDateTime:   arrivalDateTime,
	}
}
