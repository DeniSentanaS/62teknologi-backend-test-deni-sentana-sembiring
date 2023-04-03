package service

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/repository"
)

type TransactionsService interface {
	FindAll() ([]model.TransactionsEntity, error)
	FindById(Id int) (model.TransactionsEntity, error)
	Create(requestbody request.TransactionsRequest) (model.TransactionsEntity, error)
	Update(Id int, requestbody request.TransactionsRequest) (model.TransactionsEntity, error)
	Delete(Id int) (model.TransactionsEntity, error)
}

type transactionsService struct {
	transactionsRepository repository.TransactionsRepository
}

func NewTransactionsService(transactionsRepository repository.TransactionsRepository) *transactionsService {
	return &transactionsService{transactionsRepository}
}

func (s *transactionsService) FindAll() ([]model.TransactionsEntity, error) {
	transactions, err := s.transactionsRepository.FindAll()
	return transactions, err
}

func (s *transactionsService) FindById(Id int) (model.TransactionsEntity, error) {
	transactions, err := s.transactionsRepository.FindById(Id)
	return transactions, err
}

func (s *transactionsService) Create(transactionsRequest request.TransactionsRequest) (model.TransactionsEntity, error) {

	transactions := model.TransactionsEntity{
		Transaction: transactionsRequest.Transaction,
	}

	newTransactions, err := s.transactionsRepository.Create(transactions)
	return newTransactions, err
}

func (s *transactionsService) Update(Id int, transactionsRequest request.TransactionsRequest) (model.TransactionsEntity, error) {
	transactions, err := s.transactionsRepository.FindById(Id)

	transactions.Transaction = transactionsRequest.Transaction

	newTransactions, err := s.transactionsRepository.Update(transactions)
	return newTransactions, err
}

func (s *transactionsService) Delete(Id int) (model.TransactionsEntity, error) {
	transactions, err := s.transactionsRepository.FindById(Id)
	newTransactions, err := s.transactionsRepository.Delete(transactions)
	return newTransactions, err
}
