package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
	"net/http"
)

func (c *Controller) GetPositions(ctx echo.Context) error {
	var params api.PageProductParams
	ctx.Bind(&params)
	positions, err := c.dictService.GetPositions(params.Offset, params.Limit, params.Source, params.Filter)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, positions)
}

func (c *Controller) GetWarehouses(ctx echo.Context, params api.GetWarehousesParams) error {
	limit := params.Limit
	offset := params.Offset
	var source *[]string
	if params.Source == nil {
		source = &[]string{}
	} else {
		source = params.Source
	}
	whs, err := c.dictService.GetWarehouses(*offset, *limit, *source, params.Code, params.Cluster)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, whs)
}

func (c *Controller) AddWarehouse(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
