package userauth

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/userauth/request"
	"AltaStore/api/v1/userauth/response"
	auth "AltaStore/business/userauth"
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

//Login by given username and password will return JWT token
func (controller *Controller) UserLogin(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(common.BadRequestResponse())
	}

	token, err := controller.service.UserLogin(loginRequest.Email, loginRequest.Password)
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
