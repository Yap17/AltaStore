package adminauth

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/adminauth/request"
	"AltaStore/api/v1/adminauth/response"
	auth "AltaStore/business/adminauth"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service auth.Service
}

//NewController Construct item API controller
func NewController(service auth.Service) *Controller {
	return &Controller{
		service,
	}
}

//Login by given adminname and password will return JWT token
func (controller *Controller) AdminLogin(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	token, err := controller.service.AdminLogin(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	response := response.NewLoginResponse(token)

	cookie := &http.Cookie{}
	cookie.Name = "JWT"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true

	c.SetCookie(cookie)

	return c.JSON(common.SuccessResponseWithData(response))
}
