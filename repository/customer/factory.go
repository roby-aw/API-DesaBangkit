package customer

import (
	"api-desatanggap/business/customer"
	"api-desatanggap/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) customer.Repository {
	customerRepo := NewPosgresRepository(dbCon.Postgres)
	return customerRepo
}
