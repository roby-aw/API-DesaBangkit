package customer

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAccountByEmail(email string) (*Account, error)
	CreateAccount(Data *RegAccount) (*int, error)
	Createcustomer(Data *Regcustomer) (*Regcustomer, error)
	Findcustomer() ([]Customer, error)
	Detail_customer(id int) (*Detail_customer, error)
}

type Service interface {
	FindAccountByEmail(email string) (*Account, error)
	CreateAccount(Data *RegAccount) (*int, error)
	Createcustomer(Data *Regcustomer) (*Regcustomer, error)
	Findcustomer() ([]Customer, error)
	Detail_customer(id int) (*Detail_customer, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) CreateAccount(Data *RegAccount) (*int, error) {
	data, err := s.repository.FindAccountByEmail(Data.Email)
	if err != nil || data.Email != "" {
		return nil, errors.New("Email already used")
	}
	return s.repository.CreateAccount(Data)
}

func (s *service) FindAccountByEmail(email string) (*Account, error) {
	return s.repository.FindAccountByEmail(email)
}

func (s *service) Createcustomer(Data *Regcustomer) (*Regcustomer, error) {
	return s.repository.Createcustomer(Data)
}

func (s *service) Findcustomer() ([]Customer, error) {
	return s.repository.Findcustomer()
}

func (s *service) Detail_customer(id int) (*Detail_customer, error) {
	return s.repository.Detail_customer(id)
}
