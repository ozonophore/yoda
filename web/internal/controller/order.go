package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
