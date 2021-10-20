package user

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/user/request"
	"AltaStore/business/user"

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

// //GetItemByID Get item by ID echo handler
// func (controller *Controller) FindUserByID(c echo.Context) error {
// 	id := string(c.Param("id"))

// 	user, err := controller.service.FindUserByID(id)
// 	if err != nil {
// 		return c.JSON(common.NewErrorBusinessResponse(err))
// 	}

// 	response := response.NewGetUserResponse(*user)

// 	return c.JSON(common.NewSuccessResponse(response))
// }

// UpdateUser update existing user handler
func (controller *Controller) UpdateUser(c echo.Context) error {
	id := string(c.Param("id"))

	updateUserRequest := new(request.UpdateUserRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}
	user := *updateUserRequest.ToUpsertUserSpec()

	err := controller.service.UpdateUser(id, user, "modifier")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// DeleteUser delete existing user handler
func (controller *Controller) DeleteUser(c echo.Context) error {
	id := string(c.Param("id"))

	err := controller.service.DeleteUser(id, "deleter")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
