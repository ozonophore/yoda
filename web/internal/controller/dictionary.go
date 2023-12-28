package controller

import (
	"fmt"
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

func (c *Controller) UpdateWarehouse(ctx echo.Context) error {
	var params api.Warehouse
	ctx.Bind(&params)
	r, err := c.store.UpdateWarehouse(&params)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, r)
}

func (c *Controller) GetClusters(ctx echo.Context, params api.GetClustersParams) error {
	values, err := c.store.GetClusters(ctx.Request().Context(), params.Filter)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, values)
}

func (c *Controller) ExportWarehouses(ctx echo.Context, params api.ExportWarehousesParams) error {
	fileName := "warehouses_.xlsx"
	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	ctx.Response().WriteHeader(http.StatusOK)
	return c.dictService.ExportWarehouses(ctx.Response().Writer, params.Source, params.Cluster, params.Code)
}
