package shopping

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/shopping/response"
	"AltaStore/business/shopping"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service shopping.Service
}

func NewController(service shopping.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetShoppingCartByUserId(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	shoppCart, err := c.service.GetShoppingCartByUserId(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetOneResponse(shoppCart)

	return ctx.JSON(common.SuccessResponseWithData(response))
}
