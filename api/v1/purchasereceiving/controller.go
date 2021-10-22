package purchasereceiving

import (
	"AltaStore/api/common"
	"AltaStore/api/v1/purchasereceiving/request"
	"AltaStore/api/v1/purchasereceiving/response"
	"AltaStore/business/purchasereceiving"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service purchasereceiving.Service
}

func NewController(service purchasereceiving.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertPurchaseReceiving(ctx echo.Context) error {
	var err error
	userId := ctx.QueryParam("userid")

	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	insertData := new(request.InsertPurchaseReceivingRequest)

	if err = ctx.Bind(insertData); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.InsertPurchaseReceiving(insertData.ToPurchaseReceivingSpec(), userId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdatePurchaseReceiving(ctx echo.Context) error {
	var err error
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	userId := ctx.QueryParam("userid")

	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	updateData := new(request.UpdatePurchaseReceivingRequest)
	if err = ctx.Bind(updateData); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.UpdatePurchaseReceiving(id, updateData.ToPurchaseReceivingSpec(), userId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeletePurchaseReceiving(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	userid := ctx.QueryParam("userid")

	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}
	if _, err = uuid.Parse(userid); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	if err = c.service.DeletePurchaseReceiving(id, userid); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) GetAllPurchaseReceiving(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id != "" {
		if _, err := uuid.Parse(id); err != nil {
			return ctx.JSON(common.BadRequestResponse())
		}
	}

	userId := ctx.QueryParam("userid")

	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	code := ctx.QueryParam("code")
	data, err := c.service.GetAllPurchaseReceivingByParameter(code, userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(
		common.SuccessResponseWithData(
			response.GetAll(data).PurchaseReceivings,
		),
	)
}

func (c *Controller) FindPurchaseReceivingById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	userId := ctx.QueryParam("userid")

	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	data, err := c.service.GetPurchaseReceivingById(id, userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetById(*data)

	return ctx.JSON(common.SuccessResponseWithData(response))
}
