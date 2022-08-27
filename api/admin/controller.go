package admin

import (
	adminBusiness "api-desatanggap/business/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) RegisterAdmin(c echo.Context) error {
	Data := adminBusiness.RegAdmin{}
	c.Bind(&Data)
	_, err := Controller.service.CreateAdmin(&Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register account",
	})
}

func (Controller *Controller) LoginAdmin(c echo.Context) error {
	Data := adminBusiness.AuthLogin{}
	c.Bind(&Data)
	result, err := Controller.service.LoginAdmin(&Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"data":     result,
	})
}
func (Controller Controller) GetRole(c echo.Context) error {
	result, err := Controller.service.GetRole()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get role",
		"data":     result,
	})
}

func (Controller *Controller) CreateCooperation(c echo.Context) error {
	Data := adminBusiness.RegCooperation{}
	c.Bind(&Data)
	_, err := Controller.service.CreateCooperation(&Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success create cooperation",
	})
}
