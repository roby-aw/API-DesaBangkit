package modules

import (
	"api-desatanggap/api"
	customerApi "api-desatanggap/api/customer"
	customerBusiness "api-desatanggap/business/customer"
	"api-desatanggap/config"
	customerRepo "api-desatanggap/repository/customer"
	"api-desatanggap/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {
	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	controller := api.Controller{
		CustomerController: customerPermitController,
	}
	return controller
}
