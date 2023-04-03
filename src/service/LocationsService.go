package service

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/repository"
)

type LocationsService interface {
	FindAll() ([]model.LocationsEntity, error)
	FindById(Id int) (model.LocationsEntity, error)
	Create(requestbody request.LocationsRequest) (model.LocationsEntity, error)
	Update(Id int, requestbody request.LocationsRequest) (model.LocationsEntity, error)
	Delete(Id int) (model.LocationsEntity, error)
}

type locationsService struct {
	locationsRepository repository.LocationsRepository
}

func NewLocationsService(locationsRepository repository.LocationsRepository) *locationsService {
	return &locationsService{locationsRepository}
}

func (s *locationsService) FindAll() ([]model.LocationsEntity, error) {
	locations, err := s.locationsRepository.FindAll()
	return locations, err
}

func (s *locationsService) FindById(Id int) (model.LocationsEntity, error) {
	locations, err := s.locationsRepository.FindById(Id)
	return locations, err
}

func (s *locationsService) Create(locationsRequest request.LocationsRequest) (model.LocationsEntity, error) {

	locations := model.LocationsEntity{
		Address1:  locationsRequest.Address1,
		Address2:  locationsRequest.Address2,
		Address3:  locationsRequest.Address3,
		City:      locationsRequest.City,
		ZipCode:   locationsRequest.ZipCode,
		State:     locationsRequest.State,
		Latitude:  locationsRequest.Latitude,
		Longitude: locationsRequest.Longitude,
	}

	newLocations, err := s.locationsRepository.Create(locations)
	return newLocations, err
}

func (s *locationsService) Update(Id int, locationsRequest request.LocationsRequest) (model.LocationsEntity, error) {
	locations, err := s.locationsRepository.FindById(Id)

	locations.Address1 = locationsRequest.Address1
	locations.Address2 = locationsRequest.Address2
	locations.Address3 = locationsRequest.Address3
	locations.City = locationsRequest.City
	locations.ZipCode = locationsRequest.ZipCode
	locations.State = locationsRequest.State
	locations.Latitude = locationsRequest.Latitude
	locations.Longitude = locationsRequest.Longitude

	newLocations, err := s.locationsRepository.Update(locations)
	return newLocations, err
}

func (s *locationsService) Delete(Id int) (model.LocationsEntity, error) {
	locations, err := s.locationsRepository.FindById(Id)
	newLocations, err := s.locationsRepository.Delete(locations)
	return newLocations, err
}
