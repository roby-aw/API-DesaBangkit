package user

import (
	"api-desatanggap/utils"
	"errors"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAccountByEmail(email string) (*Account, error)
	CreateToken(Data *Account) (*string, error)
	CreateAccount(Data *RegAccount) (*Account, error)
	Createcustomer(Data *Regcustomer) (*Regcustomer, error)
	Findcustomer() ([]Customer, error)
	GetRole() ([]*Role, error)
	SendVerification(email string) (*string, error)
	ValidationEmail(Data string) error
	CreateCodeOtp(Email, Code string) error
	VerificationAccount(code string) error
}

type Service interface {
	FindAccountByEmail(email string) (*Account, error)
	LoginAccount(Data *AuthLogin) (*ResLogin, error)
	CreateAccount(Data *RegAccount) (*Account, error)
	Findcustomer() ([]Customer, error)
	GetRole() ([]*Role, error)
	SendVerification(email string) (*string, error)
	ValidationEmail(Data string) error
	VerificationAccount(code string) error
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

func (s *service) CreateAccount(Data *RegAccount) (*Account, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	data, _ := s.repository.FindAccountByEmail(Data.Email)
	if data != nil {
		return nil, errors.New("Email already used")
	}

	result, err := s.repository.CreateAccount(Data)
	if err != nil {
		return nil, err
	}
	code, err := s.repository.SendVerification(Data.Email)
	if err != nil {
		return nil, err
	}
	err = s.repository.CreateCodeOtp(Data.Email, *code)
	// err = s.ValidationEmail(result.Email)
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}

func (s *service) LoginAccount(Data *AuthLogin) (*ResLogin, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	Acc, err := s.repository.FindAccountByEmail(Data.Email)
	if err != nil {
		return nil, errors.New("wrong email")
	}
	err = utils.VerifyPassword(Acc.Password, Data.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}
	if Acc.IsVerified != true {
		return nil, errors.New("need to verification email")
	}
	token, err := s.repository.CreateToken(Acc)
	Response := &ResLogin{
		Account: *Acc,
		Token:   *token,
	}
	return Response, nil
}

func (s *service) FindAccountByEmail(email string) (*Account, error) {
	return s.repository.FindAccountByEmail(email)
}

// func (s *service) Createcustomer(Data *Regcustomer) (*Account, error) {
// 	return s.repository.Createcustomer(Data)
// }

func (s *service) Findcustomer() ([]Customer, error) {
	return s.repository.Findcustomer()
}

func (s *service) GetRole() ([]*Role, error) {
	return s.repository.GetRole()
}

func (s *service) SendVerification(email string) (*string, error) {
	return s.repository.SendVerification(email)
}

func (s *service) ValidationEmail(Data string) error {
	return s.repository.ValidationEmail(Data)
}

func (s *service) VerificationAccount(code string) error {
	data, err := utils.DecodeBase64(code)
	if err != nil {
		return err
	}
	return s.repository.VerificationAccount(string(data))
}
