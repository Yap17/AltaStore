package shopping

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/shopping/request"
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

func (c *Controller) NewShoppingCart(ctx echo.Context) error {
	var err error

	payload := new(request.InsertShoppingCart)

	if err = ctx.Bind(payload); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if _, err = uuid.Parse(payload.UserId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	result, err := c.service.NewShoppingCart(payload.UserId, payload.Description)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetOneResponse(result)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) GetShopCartDetailById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	itemDetail, err := c.service.GetShopCartDetailById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(itemDetail))
}

func (c *Controller) NewItemInShopCart(ctx echo.Context) error {
	var item = new(request.DetailItemInShopCart)

	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := ctx.Bind(item); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.NewItemInShopCart(id, item)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) ModifyItemInShopCart(ctx echo.Context) error {
	var item = new(request.DetailItemInShopCart)

	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := ctx.Bind(item); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.ModifyItemInShopCart(id, item)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())

}

func (c *Controller) DeleteItemInShopCart(ctx echo.Context) error {
	var err1, err2 error

	id := ctx.Param("id")
	productId := ctx.Param("productid")

	_, err1 = uuid.Parse(id)
	_, err2 = uuid.Parse(productId)

	if err1 != nil || err2 != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.DeleteItemInShopCart(id, productId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
