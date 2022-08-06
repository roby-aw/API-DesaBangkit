package customer

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Createcustomer(Data *Regcustomer) (*Regcustomer, error)
	Findcustomer() ([]Customer, error)
	Detail_customer(id int) (*Detail_customer, error)
}

type Service interface {
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

func (s *service) Createcustomer(Data *Regcustomer) (*Regcustomer, error) {
	return s.repository.Createcustomer(Data)
}

func (s *service) Findcustomer() ([]Customer, error) {
	return s.repository.Findcustomer()
}

func (s *service) Detail_customer(id int) (*Detail_customer, error) {
	return s.repository.Detail_customer(id)
}
