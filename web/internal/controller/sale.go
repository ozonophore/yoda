package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
)

func (c *Controller) GetSalesByMonthReport(ctx echo.Context, params api.GetSalesByMonthReportParams) error {
	month := params.Month
	year := params.Year
	fileName := fmt.Sprintf("sales_%d_%d.xlsx", year, month)
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	return c.saleService.PrepareAndReturnExcel(ctx.Response().Writer, uint16(year), uint8(month))
}

func (c *Controller) GetSalesByMonthWithPagination(ctx echo.Context, params api.GetSalesByMonthWithPaginationParams) error {
	page := params.Page
	size := params.Size
	month := params.Month
	year := params.Year
	sales, err := c.saleService.GetSale(uint16(year), uint8(month), page, size)
	if err != nil {
		return err
	}

	return ctx.JSON(200, sales)
}
