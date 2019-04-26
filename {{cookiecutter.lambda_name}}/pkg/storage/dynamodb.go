package storage

import (
	"github.com/Ryanair/goaws/dynamodb"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/{{cookiecutter.lambda_name}}"
)

type DynamoDBAdapter struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBAdapter(client *dynamodb.Client, tableName string) *DynamoDBAdapter {
	return &DynamoDBAdapter{
		client:    client,
		tableName: tableName,
	}
}

func (r *DynamoDBAdapter) SaveFlight(flght *{{cookiecutter.lambda_name}}.Flight) ({{cookiecutter.lambda_name}}.ID, error) {
	item := NewFlight(flght.ID, flght.CarrierCode, flght.Number, flght.Departure, flght.Arrival, flght.DepartureDateTime, flght.ArrivalDateTime)
	return item.ID, r.client.Put(item, r.tableName)
}