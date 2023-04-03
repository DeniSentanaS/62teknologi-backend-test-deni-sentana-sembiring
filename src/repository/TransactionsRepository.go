package repository

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"

	"gorm.io/gorm"
)

type TransactionsRepository interface {
	FindAll() ([]model.TransactionsEntity, error)
	FindById(Id int) (model.TransactionsEntity, error)
	Create(model model.TransactionsEntity) (model.TransactionsEntity, error)
	Update(model model.TransactionsEntity) (model.TransactionsEntity, error)
	Delete(model model.TransactionsEntity) (model.TransactionsEntity, error)
}

type transactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) *transactionsRepository {
	return &transactionsRepository{db}
}

func (r *transactionsRepository) FindAll() ([]model.TransactionsEntity, error) {
	var transaction []model.TransactionsEntity
	err := r.db.Find(&transaction).Error

	return transaction, err
}

func (r *transactionsRepository) FindById(Id int) (model.TransactionsEntity, error) {
	var transaction model.TransactionsEntity
	err := r.db.Find(&transaction, Id).Error
	return transaction, err
}

func (r *transactionsRepository) Create(transaction model.TransactionsEntity) (model.TransactionsEntity, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *transactionsRepository) Update(transaction model.TransactionsEntity) (model.TransactionsEntity, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *transactionsRepository) Delete(transaction model.TransactionsEntity) (model.TransactionsEntity, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}
