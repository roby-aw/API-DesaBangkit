package user

import (
	userBusiness "api-desatanggap/business/user"
	"api-desatanggap/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userBusiness.Service
}

func NewController(service userBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) RegisterAccount(c echo.Context) error {
	Data := userBusiness.RegAccount{}
	c.Bind(&Data)
	_, err := Controller.service.CreateAccount(&Data)
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

func (Controller *Controller) LoginAccount(c echo.Context) error {
	Data := userBusiness.AuthLogin{}
	c.Bind(&Data)
	result, err := Controller.service.LoginAccount(&Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"result":   result,
	})
}

// func (Controller *Controller) Registercustomer(c echo.Context) error {
// 	Data := userBusiness.Regcustomer{}
// 	c.Bind(&Data)
// 	result, err := Controller.service.Createcustomer(&Data)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"code":     400,
// 			"messages": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":     200,
// 		"messages": "success create data",
// 		"data":     result,
// 	})
// }

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
func (Controller *Controller) GetRole(c echo.Context) error {
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
		"result":   result,
	})
}

func (Controller *Controller) SmtpEmail(c echo.Context) error {
	email := c.QueryParam("email")
	utils.InitEmail(email)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success send email",
		"result":   email,
	})
}

func (Controller *Controller) VerificationAccount(c echo.Context) error {
	code := c.QueryParam("code")
	err := Controller.service.VerificationAccount(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success verification account",
	})
}
