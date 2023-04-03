package repository

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"

	"gorm.io/gorm"
)

type CategoriesRepository interface {
	FindAll() ([]model.CategoriesEntity, error)
	FindById(Id int) (model.CategoriesEntity, error)
	Create(model model.CategoriesEntity) (model.CategoriesEntity, error)
	Update(model model.CategoriesEntity) (model.CategoriesEntity, error)
	Delete(model model.CategoriesEntity) (model.CategoriesEntity, error)
}

type categoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *categoriesRepository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) FindAll() ([]model.CategoriesEntity, error) {
	var categories []model.CategoriesEntity
	err := r.db.Find(&categories).Error

	return categories, err
}

func (r *categoriesRepository) FindById(Id int) (model.CategoriesEntity, error) {
	var categories model.CategoriesEntity
	err := r.db.Find(&categories, Id).Error
	return categories, err
}

func (r *categoriesRepository) Create(categories model.CategoriesEntity) (model.CategoriesEntity, error) {
	err := r.db.Create(&categories).Error
	return categories, err
}

func (r *categoriesRepository) Update(categories model.CategoriesEntity) (model.CategoriesEntity, error) {
	err := r.db.Save(&categories).Error
	return categories, err
}

func (r *categoriesRepository) Delete(categories model.CategoriesEntity) (model.CategoriesEntity, error) {
	err := r.db.Delete(*&categories).Error
	return categories, err
}
