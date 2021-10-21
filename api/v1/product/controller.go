package product

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/product/request"
	"AltaStore/api/v1/product/response"
	"AltaStore/business/product"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAllProduct(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id != "" {
		if _, err := uuid.Parse(id); err != nil {
			return ctx.JSON(common.BadRequestResponse())
		}
	}

	isActive := ctx.QueryParam("isactive")
	categoryName := ctx.QueryParam("categoryname")
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	product, err := c.service.GetAllProductByParameter(id, isActive, categoryName, code, name)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(
		common.SuccessResponseWithData(
			response.GetAll(product).Products,
		),
	)
}

func (c *Controller) FindProductById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	product, err := c.service.FindProductById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetById(*product)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) InsertProduct(ctx echo.Context) error {
	var err error

	insertProduct := new(request.InsertProductRequest)

	if err = ctx.Bind(insertProduct); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.InsertProduct(insertProduct.ToProductSpec()); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdateProduct(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	updateProduct := new(request.UpdateProductRequest)
	if err = ctx.Bind(updateProduct); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.UpdateProduct(id, updateProduct.ToProductSpec()); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteProduct(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	userid := ctx.QueryParam("userid")

	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}
	if _, err = uuid.Parse(userid); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.DeleteProduct(id, userid); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
