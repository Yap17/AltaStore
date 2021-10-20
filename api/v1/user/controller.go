package user

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/user/request"
	"AltaStore/api/v1/user/response"
	"AltaStore/business/user"

	uuid "github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service user.Service
}

//NewController Construct item API controller
func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

// InsertUser Create new user handler
func (controller *Controller) InsertUser(c echo.Context) error {
	insertUserRequest := new(request.InsertUserRequest)

	if err := c.Bind(insertUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	user := *insertUserRequest.ToUpsertUserSpec()

	err := controller.service.InsertUser(user, "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindUserByID(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := controller.service.FindUserByID(id.String())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetUserResponse(*user)

	return c.JSON(common.NewSuccessResponse(response))
}

// UpdateUser update existing user handler
func (controller *Controller) UpdateUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	updateUserRequest := new(request.UpdateUserRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}
	user := *updateUserRequest.ToUpsertUserSpec()

	err := controller.service.UpdateUser(id.String(), user, "modifier")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// UpdateUserPassword update existing user handler
func (controller *Controller) UpdateUserPassword(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	updateUserPasswordRequest := new(request.UpdateUserPasswordRequest)

	if err := c.Bind(updateUserPasswordRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}
	user := *updateUserPasswordRequest.ToUpsertUserSpec()

	err := controller.service.UpdateUserPassword(id.String(), user.NewPassword, user.OldPassword)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// DeleteUser delete existing user handler
func (controller *Controller) DeleteUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	err := controller.service.DeleteUser(id.String(), "deleter")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
