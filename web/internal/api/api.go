// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
	CookieAuthScopes = "CookieAuth.Scopes"
)

// Defines values for Permission.
const (
	DASHBOARD  Permission = "DASHBOARD"
	DICTIONARY Permission = "DICTIONARY"
	HOME       Permission = "HOME"
	ORDERS     Permission = "ORDERS"
	PROFILE    Permission = "PROFILE"
	SALES      Permission = "SALES"
)

// Defines values for ProductParamsGroupBy.
const (
	POSITION ProductParamsGroupBy = "POSITION"
)

// AuthInfo defines model for AuthInfo.
type AuthInfo struct {
	AccessToken *string `json:"access_token,omitempty"`
	Description *string `json:"description,omitempty"`
	Success     bool    `json:"success"`
}

// DictPosition defines model for DictPosition.
type DictPosition struct {
	// Barcode Штрихкод
	Barcode string `json:"barcode"`

	// Code1c Код 1С
	Code1c string `json:"code1c"`

	// Id ID строки
	Id int32 `json:"id"`

	// Marketplace Торговая точка
	Marketplace string `json:"marketplace"`

	// MarketplaceId Наименование точки
	MarketplaceId string `json:"marketplaceId"`

	// Name Наименование позиции
	Name string `json:"name"`

	// Org Организация
	Org string `json:"org"`
}

// DictPositions defines model for DictPositions.
type DictPositions struct {
	// Count Count of positions
	Count int32          `json:"count"`
	Items []DictPosition `json:"items"`
}

// ErrorData defines model for ErrorData.
type ErrorData struct {
	Description string `json:"description"`
	Success     bool   `json:"success"`
}

// LoginInfo defines model for LoginInfo.
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Order defines model for Order.
type Order struct {
	Balance         int32   `json:"balance"`
	Barcode         string  `json:"barcode"`
	Brand           string  `json:"brand"`
	Code1c          string  `json:"code1c"`
	ExternalCode    string  `json:"externalCode"`
	Id              int32   `json:"id"`
	Name            string  `json:"name"`
	OrderSum        float32 `json:"orderSum"`
	OrderedQuantity int32   `json:"orderedQuantity"`
	Org             string  `json:"org"`
	Source          string  `json:"source"`
	SupplierArticle string  `json:"supplierArticle"`
}

// OrderProduct defines model for OrderProduct.
type OrderProduct struct {
	Barcode                string  `json:"barcode"`
	Brand                  string  `json:"brand"`
	Code1c                 string  `json:"code1c"`
	ExternalCode           string  `json:"externalCode"`
	Id                     int32   `json:"id"`
	Name                   string  `json:"name"`
	OrderDate              *string `json:"orderDate,omitempty"`
	OrderQuantityCanceled  int32   `json:"orderQuantityCanceled"`
	OrderQuantityDelivered int32   `json:"orderQuantityDelivered"`
	OrderedQuantity        int32   `json:"orderedQuantity"`
	Org                    string  `json:"org"`
	Source                 string  `json:"source"`
	SupplierArticle        string  `json:"supplierArticle"`
}

// OrderProducts defines model for OrderProducts.
type OrderProducts struct {
	// Count Count of stocks
	Count int32          `json:"count"`
	Items []OrderProduct `json:"items"`
}

// Orders defines model for Orders.
type Orders struct {
	// Count Count of stocks
	Count int32   `json:"count"`
	Items []Order `json:"items"`
}

// PageProductParams defines model for PageProductParams.
type PageProductParams struct {
	Filter *string  `json:"filter,omitempty"`
	Limit  int32    `json:"limit"`
	Offset int32    `json:"offset"`
	Source []string `json:"source"`
}

// Permission defines model for Permission.
type Permission string

// ProductParams defines model for ProductParams.
type ProductParams struct {
	DateFrom openapi_types.Date    `json:"dateFrom"`
	DateTo   openapi_types.Date    `json:"dateTo"`
	Filter   *string               `json:"filter,omitempty"`
	GroupBy  *ProductParamsGroupBy `json:"groupBy,omitempty"`
	Limit    *int32                `json:"limit,omitempty"`
	Offset   *int32                `json:"offset,omitempty"`
}

// ProductParamsGroupBy defines model for ProductParams.GroupBy.
type ProductParamsGroupBy string

// Profile defines model for Profile.
type Profile struct {
	// Email Email пользователя
	Email string `json:"email"`

	// Name Имя пользователя
	Name string `json:"name"`

	// Permissions Права пользователя
	Permissions []Permission `json:"permissions"`
}

// Sale defines model for Sale.
type Sale struct {
	Barcode         string  `json:"barcode"`
	Code1c          string  `json:"code1c"`
	Country         string  `json:"country"`
	ExternalCode    string  `json:"externalCode"`
	Id              int32   `json:"id"`
	Name            string  `json:"name"`
	Oblast          string  `json:"oblast"`
	Quantity        int32   `json:"quantity"`
	Region          string  `json:"region"`
	Source          string  `json:"source"`
	SupplierArticle string  `json:"supplierArticle"`
	TotalPrice      float64 `json:"total_price"`
}

// Sales defines model for Sales.
type Sales struct {
	// Count Count of saleses
	Count int32  `json:"count"`
	Items []Sale `json:"items"`
}

// Stock defines model for Stock.
type Stock struct {
	// Barcode Штрихкод
	Barcode string `json:"barcode"`

	// Marketplace Торговая точка
	Marketplace string `json:"marketplace"`

	// MarketplaceId ID торговой точки
	MarketplaceId string `json:"marketplace_id"`

	// Organization Организация
	Organization string `json:"organization"`

	// OrganizationId ID организации
	OrganizationId string `json:"organization_id"`

	// Quantity Stock quantity
	Quantity int32 `json:"quantity"`
}

// Stocks defines model for Stocks.
type Stocks struct {
	// Count Count of stocks
	Count int32   `json:"count"`
	Items []Stock `json:"items"`
}

// Warehouse defines model for Warehouse.
type Warehouse struct {
	// Cluster Наименование кластера
	Cluster string `json:"cluster"`

	// Code Код склада
	Code string `json:"code"`

	// Source Источник
	Source string `json:"source"`
}

// Warehouses defines model for Warehouses.
type Warehouses struct {
	// Count Count of positions
	Count int32       `json:"count"`
	Items []Warehouse `json:"items"`
}

// Error defines model for Error.
type Error = ErrorData

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError = AuthInfo

// GetWarehousesParams defines parameters for GetWarehouses.
type GetWarehousesParams struct {
	Source  *[]string `form:"source,omitempty" json:"source,omitempty"`
	Limit   *int32    `form:"limit,omitempty" json:"limit,omitempty"`
	Offset  *int32    `form:"offset,omitempty" json:"offset,omitempty"`
	Cluster *string   `form:"cluster,omitempty" json:"cluster,omitempty"`
	Code    *string   `form:"code,omitempty" json:"code,omitempty"`
}

// GetOrdersParams defines parameters for GetOrders.
type GetOrdersParams struct {
	Date   *openapi_types.Date `form:"date,omitempty" json:"date,omitempty"`
	Page   int32               `form:"page" json:"page"`
	Size   int32               `form:"size" json:"size"`
	Filter *string             `form:"filter,omitempty" json:"filter,omitempty"`
	Source *string             `form:"source,omitempty" json:"source,omitempty"`
}

// GetOrdersReportParams defines parameters for GetOrdersReport.
type GetOrdersReportParams struct {
	Date *openapi_types.Date `form:"date,omitempty" json:"date,omitempty"`
}

// GetSalesByMonthWithPaginationParams defines parameters for GetSalesByMonthWithPagination.
type GetSalesByMonthWithPaginationParams struct {
	Year  int32 `form:"year" json:"year"`
	Month int32 `form:"month" json:"month"`
	Page  int32 `form:"page" json:"page"`
	Size  int32 `form:"size" json:"size"`
}

// GetSalesByMonthReportParams defines parameters for GetSalesByMonthReport.
type GetSalesByMonthReportParams struct {
	Year  int32 `form:"year" json:"year"`
	Month int32 `form:"month" json:"month"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginInfo

// GetPositionsJSONRequestBody defines body for GetPositions for application/json ContentType.
type GetPositionsJSONRequestBody = PageProductParams

// AddWarehouseJSONRequestBody defines body for AddWarehouse for application/json ContentType.
type AddWarehouseJSONRequestBody = Warehouse

// GetOrdersProductJSONRequestBody defines body for GetOrdersProduct for application/json ContentType.
type GetOrdersProductJSONRequestBody = ProductParams

// ExportOrdersProductToExcelJSONRequestBody defines body for ExportOrdersProductToExcel for application/json ContentType.
type ExportOrdersProductToExcelJSONRequestBody = ProductParams

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/login)
	Login(ctx echo.Context) error

	// (GET /auth/profile)
	Profile(ctx echo.Context) error

	// (GET /auth/refresh)
	Refresh(ctx echo.Context) error

	// (POST /dictionaries/positions)
	GetPositions(ctx echo.Context) error

	// (GET /dictionaries/warehouses)
	GetWarehouses(ctx echo.Context, params GetWarehousesParams) error

	// (POST /dictionaries/warehouses)
	AddWarehouse(ctx echo.Context) error

	// (GET /orders)
	GetOrders(ctx echo.Context, params GetOrdersParams) error

	// (POST /orders/product)
	GetOrdersProduct(ctx echo.Context) error

	// (POST /orders/product/report)
	ExportOrdersProductToExcel(ctx echo.Context) error

	// (GET /orders/report)
	GetOrdersReport(ctx echo.Context, params GetOrdersReportParams) error

	// (GET /sales)
	GetSalesByMonthWithPagination(ctx echo.Context, params GetSalesByMonthWithPaginationParams) error

	// (GET /sales/report)
	GetSalesByMonthReport(ctx echo.Context, params GetSalesByMonthReportParams) error
	// Получение остатков товаров
	// (GET /stocks/{date})
	GetStocksDate(ctx echo.Context, date openapi_types.Date) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Profile converts echo context to params.
func (w *ServerInterfaceWrapper) Profile(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Profile(ctx)
	return err
}

// Refresh converts echo context to params.
func (w *ServerInterfaceWrapper) Refresh(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Refresh(ctx)
	return err
}

// GetPositions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPositions(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPositions(ctx)
	return err
}

// GetWarehouses converts echo context to params.
func (w *ServerInterfaceWrapper) GetWarehouses(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetWarehousesParams
	// ------------- Optional query parameter "source" -------------

	err = runtime.BindQueryParameter("form", true, false, "source", ctx.QueryParams(), &params.Source)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "cluster" -------------

	err = runtime.BindQueryParameter("form", true, false, "cluster", ctx.QueryParams(), &params.Cluster)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter cluster: %s", err))
	}

	// ------------- Optional query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, false, "code", ctx.QueryParams(), &params.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter code: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetWarehouses(ctx, params)
	return err
}

// AddWarehouse converts echo context to params.
func (w *ServerInterfaceWrapper) AddWarehouse(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddWarehouse(ctx)
	return err
}

// GetOrders converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrders(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetOrdersParams
	// ------------- Optional query parameter "date" -------------

	err = runtime.BindQueryParameter("form", true, false, "date", ctx.QueryParams(), &params.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	// ------------- Required query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, true, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Required query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, true, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Optional query parameter "source" -------------

	err = runtime.BindQueryParameter("form", true, false, "source", ctx.QueryParams(), &params.Source)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOrders(ctx, params)
	return err
}

// GetOrdersProduct converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrdersProduct(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOrdersProduct(ctx)
	return err
}

// ExportOrdersProductToExcel converts echo context to params.
func (w *ServerInterfaceWrapper) ExportOrdersProductToExcel(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ExportOrdersProductToExcel(ctx)
	return err
}

// GetOrdersReport converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrdersReport(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetOrdersReportParams
	// ------------- Optional query parameter "date" -------------

	err = runtime.BindQueryParameter("form", true, false, "date", ctx.QueryParams(), &params.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOrdersReport(ctx, params)
	return err
}

// GetSalesByMonthWithPagination converts echo context to params.
func (w *ServerInterfaceWrapper) GetSalesByMonthWithPagination(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSalesByMonthWithPaginationParams
	// ------------- Required query parameter "year" -------------

	err = runtime.BindQueryParameter("form", true, true, "year", ctx.QueryParams(), &params.Year)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter year: %s", err))
	}

	// ------------- Required query parameter "month" -------------

	err = runtime.BindQueryParameter("form", true, true, "month", ctx.QueryParams(), &params.Month)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter month: %s", err))
	}

	// ------------- Required query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, true, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Required query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, true, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSalesByMonthWithPagination(ctx, params)
	return err
}

// GetSalesByMonthReport converts echo context to params.
func (w *ServerInterfaceWrapper) GetSalesByMonthReport(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSalesByMonthReportParams
	// ------------- Required query parameter "year" -------------

	err = runtime.BindQueryParameter("form", true, true, "year", ctx.QueryParams(), &params.Year)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter year: %s", err))
	}

	// ------------- Required query parameter "month" -------------

	err = runtime.BindQueryParameter("form", true, true, "month", ctx.QueryParams(), &params.Month)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter month: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSalesByMonthReport(ctx, params)
	return err
}

// GetStocksDate converts echo context to params.
func (w *ServerInterfaceWrapper) GetStocksDate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "date" -------------
	var date openapi_types.Date

	err = runtime.BindStyledParameterWithLocation("simple", false, "date", runtime.ParamLocationPath, ctx.Param("date"), &date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStocksDate(ctx, date)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/auth/login", wrapper.Login)
	router.GET(baseURL+"/auth/profile", wrapper.Profile)
	router.GET(baseURL+"/auth/refresh", wrapper.Refresh)
	router.POST(baseURL+"/dictionaries/positions", wrapper.GetPositions)
	router.GET(baseURL+"/dictionaries/warehouses", wrapper.GetWarehouses)
	router.POST(baseURL+"/dictionaries/warehouses", wrapper.AddWarehouse)
	router.GET(baseURL+"/orders", wrapper.GetOrders)
	router.POST(baseURL+"/orders/product", wrapper.GetOrdersProduct)
	router.POST(baseURL+"/orders/product/report", wrapper.ExportOrdersProductToExcel)
	router.GET(baseURL+"/orders/report", wrapper.GetOrdersReport)
	router.GET(baseURL+"/sales", wrapper.GetSalesByMonthWithPagination)
	router.GET(baseURL+"/sales/report", wrapper.GetSalesByMonthReport)
	router.GET(baseURL+"/stocks/:date", wrapper.GetStocksDate)

}
