package repository

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"

	"gorm.io/gorm"
)

type LocationsRepository interface {
	FindAll() ([]model.LocationsEntity, error)
	FindById(Id int) (model.LocationsEntity, error)
	Create(model model.LocationsEntity) (model.LocationsEntity, error)
	Update(model model.LocationsEntity) (model.LocationsEntity, error)
	Delete(model model.LocationsEntity) (model.LocationsEntity, error)
}

type locationsRepository struct {
	db *gorm.DB
}

func NewLocationsRepository(db *gorm.DB) *locationsRepository {
	return &locationsRepository{db}
}

func (r *locationsRepository) FindAll() ([]model.LocationsEntity, error) {
	var locations []model.LocationsEntity
	err := r.db.Find(&locations).Error

	return locations, err
}

func (r *locationsRepository) FindById(Id int) (model.LocationsEntity, error) {
	var locations model.LocationsEntity
	err := r.db.Find(&locations, Id).Error
	return locations, err
}

func (r *locationsRepository) Create(locations model.LocationsEntity) (model.LocationsEntity, error) {
	err := r.db.Create(&locations).Error
	return locations, err
}

func (r *locationsRepository) Update(locations model.LocationsEntity) (model.LocationsEntity, error) {
	err := r.db.Save(&locations).Error
	return locations, err
}

func (r *locationsRepository) Delete(locations model.LocationsEntity) (model.LocationsEntity, error) {
	err := r.db.Delete(&locations).Error
	return locations, err
}
