package api

import (
	"api-desatanggap/api/admin"
	"api-desatanggap/api/user"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserController  *user.Controller
	AdminController *admin.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.GET("/photo/:name", func(c echo.Context) error {
		name := fmt.Sprintf("utils/img/%s", c.Param("name"))
		fmt.Println(name)
		return c.Inline(name, name)
	})
	acc := e.Group("/users")
	acc.POST("/login", controller.UserController.LoginAccount)
	acc.POST("/registrations", controller.UserController.RegisterAccount)
	e.POST("/upload", controller.UserController.UploadPhoto)
	// e.POST("/register", controller.UserController.RegisterUser)
	// e.GET("/User", controller.UserController.FindUser)
	admin := e.Group("/administrators")
	admin.GET("/role", controller.AdminController.GetRole)
	admin.POST("/registrations", controller.AdminController.RegisterAdmin)
	admin.POST("/login", controller.AdminController.LoginAdmin)

}
