package category

import (
	"AltaStore/business/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service category.Service
}

func NewController(service category.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAllCategory(ctx echo.Context) error {
	categories, err := c.service.GetAllCategory()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "")
	}

	return ctx.JSON(http.StatusOK, categories)

}
