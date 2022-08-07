package api

import (
	"api-desatanggap/api/customer"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	CustomerController *customer.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.GET("/photo/:name", func(c echo.Context) error {
		name := fmt.Sprintf("utils/img/%s", c.Param("name"))
		fmt.Println(name)
		return c.Inline(name, name)
	})
	acc := e.Group("/users")
	acc.POST("/login", controller.CustomerController.LoginAccount)
	acc.POST("/registrations", controller.CustomerController.RegisterAccount)
	e.POST("/upload", controller.CustomerController.UploadPhoto)
	e.POST("/register", controller.CustomerController.Registercustomer)
	e.GET("/customer", controller.CustomerController.Findcustomer)
	e.GET("/detcustomer/:id", controller.CustomerController.Detail_customer)
}
