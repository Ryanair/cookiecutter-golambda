package {{cookiecutter.lambda_name}}_test 

import (
	"fmt"
	"testing"
	"time"

	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/{{cookiecutter.lambda_name}}"
	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/pkg/storage"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	inErr   error
	success bool
}{
	{nil, true},
	{fmt.Errorf("create flight error"), false},
}

func TestCreateFlight(t *testing.T) {
	departureDateTime, _ := time.Parse(time.RFC3339, "2019-04-25T18:55:00Z")
	arrivalDateTime, _ := time.Parse(time.RFC3339, "2019-04-25T21:55:00Z")
	flght := {{cookiecutter.lambda_name}}.NewFlight("FR", "FR-2345", "WRO", "DUB", departureDateTime, arrivalDateTime)

	for _, test := range tests {
		errFn := storage.CreateFlightErr(test.inErr)
		repository := storage.NewInMemoryDBAdapter(errFn)
		service := {{cookiecutter.lambda_name}}.NewService(repository)
		id, err := service.CreateFlight(&flght)
		assert.Equal(t, test.success, id != "" && err == nil)
	}
}
