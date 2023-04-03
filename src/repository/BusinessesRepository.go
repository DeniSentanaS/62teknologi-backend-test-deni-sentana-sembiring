package repository

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"

	"gorm.io/gorm"
)

type BusinessesRepository interface {
	FindAll() ([]model.BusinessesEntity, error)
	FindById(Id int) (model.BusinessesEntity, error)
	Create(model model.BusinessesEntity) (model.BusinessesEntity, error)
	Update(model model.BusinessesEntity) (model.BusinessesEntity, error)
	Delete(model model.BusinessesEntity) (model.BusinessesEntity, error)
}

type businessesRepository struct {
	db *gorm.DB
}

func NewBusinessesRepository(db *gorm.DB) *businessesRepository {
	return &businessesRepository{db}
}

func (r *businessesRepository) FindAll() ([]model.BusinessesEntity, error) {
	var businesses []model.BusinessesEntity
	err := r.db.Find(&businesses).Error

	return businesses, err
}

func (r *businessesRepository) FindById(Id int) (model.BusinessesEntity, error) {
	var businesses model.BusinessesEntity
	err := r.db.Find(&businesses, Id).Error
	return businesses, err
}

func (r *businessesRepository) Create(businesses model.BusinessesEntity) (model.BusinessesEntity, error) {
	err := r.db.Create(&businesses).Error
	return businesses, err
}

func (r *businessesRepository) Update(businesses model.BusinessesEntity) (model.BusinessesEntity, error) {
	err := r.db.Save(&businesses).Error
	return businesses, err
}

func (r *businessesRepository) Delete(businesses model.BusinessesEntity) (model.BusinessesEntity, error) {
	err := r.db.Delete(&businesses).Error
	return businesses, err
}
