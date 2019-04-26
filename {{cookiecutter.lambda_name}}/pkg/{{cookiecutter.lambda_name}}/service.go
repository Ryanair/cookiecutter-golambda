package {{cookiecutter.lambda_name}}

import "github.com/pkg/errors"

type repository interface {
	SaveFlight(flght *Flight) (ID, error)
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateFlight(flight *Flight) (ID, error) {
	id, err := s.repository.SaveFlight(flight)
	return id, errors.Wrap(err, "cannot create flight")
}