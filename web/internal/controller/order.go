package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yoda/web/internal/api"
	"net/http"
	"time"
)

func (c *Controller) GetOrdersReport(ctx echo.Context, params api.GetOrdersReportParams) error {
	reportDate := time.Now()
	if params.Date != nil {
		reportDate = params.Date.Time
	}
	fileName := fmt.Sprintf("orders_%s.xlsx", reportDate.Format(time.DateOnly))
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	err := c.orderService.PrepareAndReturnExcel(ctx.Response().Writer, reportDate)
	return err
}

func (c *Controller) GetOrders(ctx echo.Context, params api.GetOrdersParams) error {
	page := params.Page
	size := params.Size
	reportDate := time.Now()
	if params.Date != nil {
		reportDate = params.Date.Time
	}
	filter := ""
	if params.Filter != nil {
		filter = *params.Filter
	}
	sourse := ""
	if params.Source != nil {
		sourse = *params.Source
	}
	orders, err := c.orderService.GetOrders(reportDate, filter, sourse, page, size)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (c *Controller) GetOrdersProduct(ctx echo.Context) error {
	var params api.ProductParams
	err := ctx.Bind(&params)
	if err != nil {
		logrus.Error("GetOrdersProduct: ", err)
		return ctx.JSON(418, api.ErrorData{
			Success:     false,
			Description: err.Error(),
		})
	}
	offset := params.Offset
	limit := params.Limit
	order, err := c.orderService.GetOrdersProduct(params.DateFrom.Time, params.DateTo.Time, params.Filter, *offset, *limit, (*string)(params.GroupBy))
	if err != nil {
		logrus.Error("GetOrdersProduct: ", err)
		return ctx.JSON(418, api.ErrorData{
			Success:     false,
			Description: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK,
		order)
}

func (c *Controller) ExportOrdersProductToExcel(ctx echo.Context) error {
	var params api.ProductParams
	err := ctx.Bind(&params)
	if err != nil {
		logrus.Error("GetOrdersProduct: ", err)
		return ctx.JSON(418, api.ErrorData{
			Success:     false,
			Description: err.Error(),
		})
	}
	fileName := fmt.Sprintf("orders_%s_%s.xlsx", params.DateFrom.Time.Format("20060102"), params.DateTo.Time.Format("20060102"))
	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	ctx.Response().WriteHeader(http.StatusOK)
	return c.orderService.ExportOrderProductReport(ctx.Response().Writer, params.DateFrom.Time, params.DateTo.Time, params.Filter, (*string)(params.GroupBy))
}
