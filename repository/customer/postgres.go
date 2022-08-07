package customer

import (
	"api-desatanggap/business/customer"
	"api-desatanggap/repository"
	"time"

	"gorm.io/gorm"
)

type PosgresRepository struct {
	db *gorm.DB
}

func NewPosgresRepository(db *gorm.DB) *PosgresRepository {
	return &PosgresRepository{
		db: db,
	}
}

func (repo *PosgresRepository) FindAccountByEmail(email string) (*customer.Account, error) {
	var data customer.Account
	repo.db.Model(&repository.Account{}).First(&data)
	return &data, nil
}

func (repo *PosgresRepository) CreateAccount(Data *customer.RegAccount) (*int, error) {
	err := repo.db.Create(&repository.Account{Fullname: Data.Fullname, Email: Data.Email, Password: Data.Password, Role: Data.Role}).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *PosgresRepository) Createcustomer(Data *customer.Regcustomer) (*customer.Regcustomer, error) {
	date, _ := time.Parse("2006-01-02", Data.Tanggal_lahir)
	err := repo.db.Create(&repository.Customer{Nama: Data.Nama, Gender: Data.Gender, Tanggal_lahir: date}).Error
	if err != nil {
		return nil, err
	}
	return Data, nil
}

func (repo *PosgresRepository) Findcustomer() ([]customer.Customer, error) {
	var data []customer.Customer
	repo.db.Find(&data)
	for n, v := range data {
		if v.Gender == "0" {
			data[n].Gender = "Laki-Laki"
		}
		if v.Gender == "1" {
			data[n].Gender = "Perempuan"
		}
	}
	return data, nil
}

func (repo *PosgresRepository) Detail_customer(id int) (*customer.Detail_customer, error) {
	var data *customer.Detail_customer
	err := repo.db.Model(&repository.Customer{}).Where("id = ?", id).Preload("Jurusan").Preload("Hobi").First(&data).Error
	if err != nil {
		return nil, err
	}
	if data.Gender == "1" {
		data.Gender = "Perempuan"
	}
	if data.Gender == "0" {
		data.Gender = "Laki-Laki"
	}
	return data, nil
}
