package storage

import (
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/{{cookiecutter.lambda_name}}"
)

type InMemoryDBAdapter struct {
	flights         []Flight
	createFlightErr error
}

func NewInMemoryDBAdapter(options ...func(*InMemoryDBAdapter)) *InMemoryDBAdapter {
	db := InMemoryDBAdapter{}
	flights := make([]Flight, 0)
	db.flights = flights
	for _, option := range options {
		option(&db)
	}
	return &db
}

func (m *InMemoryDBAdapter) SaveFlight(flght *{{cookiecutter.lambda_name}}.Flight) ({{cookiecutter.lambda_name}}.ID, error) {
	item := NewFlight(flght.ID, flght.CarrierCode, flght.Number, flght.Departure, flght.Arrival, flght.DepartureDateTime, flght.ArrivalDateTime)
	if m.createFlightErr != nil {
		return "", m.createFlightErr
	}
	m.flights = append(m.flights, item)
	return flght.ID, nil
}

func CreateFlightErr(err error) func(*InMemoryDBAdapter) {
	return func(m *InMemoryDBAdapter) {
		m.createFlightErr = err
	}
}
