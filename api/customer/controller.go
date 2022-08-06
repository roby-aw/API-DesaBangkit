package customer

import (
	customerBusiness "api-desatanggap/business/customer"
	"api-desatanggap/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service customerBusiness.Service
}

func NewController(service customerBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) Registercustomer(c echo.Context) error {
	Data := customerBusiness.Regcustomer{}
	c.Bind(&Data)
	result, err := Controller.service.Createcustomer(&Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success create data",
		"data":     result,
	})
}

func (Controller *Controller) Findcustomer(c echo.Context) error {
	result, err := Controller.service.Findcustomer()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get all data customer",
		"data":     result,
	})
}

func (Controller *Controller) Detail_customer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	result, err := Controller.service.Detail_customer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get detail customer",
		"data":     result,
	})
}

func (Controller *Controller) UploadPhoto(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	file, err := c.FormFile("file")
	err = utils.Upload(name, email, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success upload",
	})
}
