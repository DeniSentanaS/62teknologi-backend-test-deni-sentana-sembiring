package service

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/repository"
)

type CategoriesService interface {
	FindAll() ([]model.CategoriesEntity, error)
	FindById(Id int) (model.CategoriesEntity, error)
	Create(requestbody request.CategoriesRequest) (model.CategoriesEntity, error)
	Update(Id int, requestbody request.CategoriesRequest) (model.CategoriesEntity, error)
	Delete(Id int) (model.CategoriesEntity, error)
}

type categoriesService struct {
	categoriesRepository repository.CategoriesRepository
}

func NewCategoriesService(categoriesRepository repository.CategoriesRepository) *categoriesService {
	return &categoriesService{categoriesRepository}
}

func (s *categoriesService) FindAll() ([]model.CategoriesEntity, error) {
	businesses, err := s.categoriesRepository.FindAll()
	return businesses, err
}

func (s *categoriesService) FindById(Id int) (model.CategoriesEntity, error) {
	categories, err := s.categoriesRepository.FindById(Id)
	return categories, err
}

func (s *categoriesService) Create(categoriesRequest request.CategoriesRequest) (model.CategoriesEntity, error) {

	categories := model.CategoriesEntity{
		Alias: categoriesRequest.Alias,
		Title: categoriesRequest.Title,
	}

	newCategories, err := s.categoriesRepository.Create(categories)
	return newCategories, err
}

func (s *categoriesService) Update(Id int, categoriesRequest request.CategoriesRequest) (model.CategoriesEntity, error) {
	categories, err := s.categoriesRepository.FindById(Id)

	categories.Alias = categoriesRequest.Alias
	categories.Title = categoriesRequest.Title

	newCategories, err := s.categoriesRepository.Update(categories)
	return newCategories, err
}

func (s *categoriesService) Delete(Id int) (model.CategoriesEntity, error) {
	businesses, err := s.categoriesRepository.FindById(Id)
	newCategories, err := s.categoriesRepository.Delete(businesses)
	return newCategories, err
}
