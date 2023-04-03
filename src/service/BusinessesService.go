package service

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/repository"
)

type BusinessesService interface {
	FindAll() ([]model.BusinessesEntity, error)
	FindById(Id int) (model.BusinessesEntity, error)
	Create(requestbody request.BusinessesRequest) (model.BusinessesEntity, error)
	Update(Id int, requestbody request.BusinessesRequest) (model.BusinessesEntity, error)
	Delete(Id int) (model.BusinessesEntity, error)
}

type businessesService struct {
	businessesRepository repository.BusinessesRepository
}

func NewBusinessesService(businessesRepository repository.BusinessesRepository) *businessesService {
	return &businessesService{businessesRepository}
}

func (s *businessesService) FindAll() ([]model.BusinessesEntity, error) {
	businesses, err := s.businessesRepository.FindAll()
	return businesses, err
}

func (s *businessesService) FindById(Id int) (model.BusinessesEntity, error) {
	businesses, err := s.businessesRepository.FindById(Id)
	return businesses, err
}

func (s *businessesService) Create(businessesRequest request.BusinessesRequest) (model.BusinessesEntity, error) {
	price, _ := businessesRequest.ReviewCount.Int64()

	businesses := model.BusinessesEntity{
		Alias:          businessesRequest.Alias,
		Name:           businessesRequest.Name,
		ImageUrl:       businessesRequest.ImageUrl,
		IsClosed:       businessesRequest.IsClosed,
		Url:            businessesRequest.Url,
		ReviewCount:    int(price),
		Rating:         businessesRequest.Rating,
		Distance:       businessesRequest.Distance,
		CategoriesId:   businessesRequest.CategoriesId,
		LocationsId:    businessesRequest.LocationsId,
		TransactionsId: businessesRequest.TransactionsId,
	}

	newBusinesses, err := s.businessesRepository.Create(businesses)
	return newBusinesses, err
}

func (s *businessesService) Update(Id int, businessesRequest request.BusinessesRequest) (model.BusinessesEntity, error) {
	businesses, err := s.businessesRepository.FindById(Id)

	price, _ := businessesRequest.ReviewCount.Int64()

	businesses.Alias = businessesRequest.Alias
	businesses.Name = businessesRequest.Name
	businesses.ImageUrl = businessesRequest.ImageUrl
	businesses.IsClosed = businessesRequest.IsClosed
	businesses.Url = businessesRequest.Url
	businesses.ReviewCount = int(price)
	businesses.Rating = businessesRequest.Rating
	businesses.Distance = businessesRequest.Distance

	newBusinesses, err := s.businessesRepository.Update(businesses)
	return newBusinesses, err
}

func (s *businessesService) Delete(Id int) (model.BusinessesEntity, error) {
	businesses, err := s.businessesRepository.FindById(Id)
	newBusinesses, err := s.businessesRepository.Delete(businesses)
	return newBusinesses, err
}
