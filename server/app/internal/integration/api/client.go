// Package integration provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/yoda/common/pkg/types"
)

// Barcode defines model for Barcode.
type Barcode struct {
	// Barcode Barcode
	Barcode string `json:"barcode"`

	// BarcodeID Barcode ID from 1c
	BarcodeID string `json:"barcodeID"`

	// Id ID from 1c
	Id string `json:"id"`

	// MarketId Market ID from 1c
	MarketId string `json:"marketId"`

	// OrgId Organization ID from 1c
	OrgId string `json:"orgId"`

	// UpdateAt Barcode update date
	UpdateAt types.CustomTime `json:"updateAt"`
}

// Barcodes defines model for Barcodes.
type Barcodes struct {
	// Count Count of barcodes
	Count int32     `json:"count"`
	Items []Barcode `json:"items"`
}

// Item defines model for Item.
type Item struct {
	// Id Item ID from 1c
	Id string `json:"id"`

	// Name Item name
	Name string `json:"name"`

	// UpdateAt Item update date
	UpdateAt types.CustomTime `json:"updateAt"`
}

// Items defines model for Items.
type Items struct {
	// Count Count of items
	Count int32  `json:"count"`
	Items []Item `json:"items"`
}

// Marketplace defines model for Marketplace.
type Marketplace struct {
	// Id Marketplace ID from 1c
	Id string `json:"id"`

	// Name Marketplace name
	Name     string           `json:"name"`
	UpdateAt types.CustomTime `json:"updateAt"`
}

// Marketplaces defines model for Marketplaces.
type Marketplaces struct {
	// Count Count of marketplaces
	Count int32         `json:"count"`
	Items []Marketplace `json:"items"`
}

// Organization defines model for Organization.
type Organization struct {
	// Id Organization ID from 1c
	Id string `json:"id"`

	// Inn Organization inn
	Inn *string `json:"inn,omitempty"`

	// Kpp Organization kpp
	Kpp *string `json:"kpp,omitempty"`

	// Name Organization name
	Name string `json:"name"`

	// UpdateAt Organization update date
	UpdateAt types.CustomTime `json:"updateAt"`
}

// Organizations defines model for Organizations.
type Organizations struct {
	// Count Count of organizations
	Count int32          `json:"count"`
	Items []Organization `json:"items"`
}

// Stock defines model for Stock.
type Stock struct {
	// Id Stock ID from 1c
	Id string `json:"id"`

	// Quantity Stock quantity
	Quantity float64 `json:"quantity"`
}

// Stocks defines model for Stocks.
type Stocks struct {
	// Count Count of stocks
	Count int32   `json:"count"`
	Items []Stock `json:"items"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetItems request
	GetItems(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetItemsBarcodes request
	GetItemsBarcodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetMarketplaces request
	GetMarketplaces(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOrganizations request
	GetOrganizations(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetStocks request
	GetStocks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetItems(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetItemsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetItemsBarcodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetItemsBarcodesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetMarketplaces(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMarketplacesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOrganizations(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOrganizationsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetStocks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStocksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetItemsRequest generates requests for GetItems
func NewGetItemsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/items")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetItemsBarcodesRequest generates requests for GetItemsBarcodes
func NewGetItemsBarcodesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/items/barcodes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetMarketplacesRequest generates requests for GetMarketplaces
func NewGetMarketplacesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/marketplaces")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetOrganizationsRequest generates requests for GetOrganizations
func NewGetOrganizationsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/organizations")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetStocksRequest generates requests for GetStocks
func NewGetStocksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/stocks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetItems request
	GetItemsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetItemsResponse, error)

	// GetItemsBarcodes request
	GetItemsBarcodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetItemsBarcodesResponse, error)

	// GetMarketplaces request
	GetMarketplacesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMarketplacesResponse, error)

	// GetOrganizations request
	GetOrganizationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganizationsResponse, error)

	// GetStocks request
	GetStocksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStocksResponse, error)
}

type GetItemsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Items
}

// Status returns HTTPResponse.Status
func (r GetItemsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetItemsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetItemsBarcodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Barcodes
}

// Status returns HTTPResponse.Status
func (r GetItemsBarcodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetItemsBarcodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMarketplacesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Marketplaces
}

// Status returns HTTPResponse.Status
func (r GetMarketplacesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetMarketplacesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOrganizationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Organizations
}

// Status returns HTTPResponse.Status
func (r GetOrganizationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOrganizationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetStocksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Stocks
}

// Status returns HTTPResponse.Status
func (r GetStocksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStocksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetItemsWithResponse request returning *GetItemsResponse
func (c *ClientWithResponses) GetItemsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetItemsResponse, error) {
	rsp, err := c.GetItems(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetItemsResponse(rsp)
}

// GetItemsBarcodesWithResponse request returning *GetItemsBarcodesResponse
func (c *ClientWithResponses) GetItemsBarcodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetItemsBarcodesResponse, error) {
	rsp, err := c.GetItemsBarcodes(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetItemsBarcodesResponse(rsp)
}

// GetMarketplacesWithResponse request returning *GetMarketplacesResponse
func (c *ClientWithResponses) GetMarketplacesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMarketplacesResponse, error) {
	rsp, err := c.GetMarketplaces(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMarketplacesResponse(rsp)
}

// GetOrganizationsWithResponse request returning *GetOrganizationsResponse
func (c *ClientWithResponses) GetOrganizationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganizationsResponse, error) {
	rsp, err := c.GetOrganizations(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOrganizationsResponse(rsp)
}

// GetStocksWithResponse request returning *GetStocksResponse
func (c *ClientWithResponses) GetStocksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStocksResponse, error) {
	rsp, err := c.GetStocks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetStocksResponse(rsp)
}

// ParseGetItemsResponse parses an HTTP response from a GetItemsWithResponse call
func ParseGetItemsResponse(rsp *http.Response) (*GetItemsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetItemsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Items
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetItemsBarcodesResponse parses an HTTP response from a GetItemsBarcodesWithResponse call
func ParseGetItemsBarcodesResponse(rsp *http.Response) (*GetItemsBarcodesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetItemsBarcodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Barcodes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetMarketplacesResponse parses an HTTP response from a GetMarketplacesWithResponse call
func ParseGetMarketplacesResponse(rsp *http.Response) (*GetMarketplacesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetMarketplacesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Marketplaces
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetOrganizationsResponse parses an HTTP response from a GetOrganizationsWithResponse call
func ParseGetOrganizationsResponse(rsp *http.Response) (*GetOrganizationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOrganizationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Organizations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetStocksResponse parses an HTTP response from a GetStocksWithResponse call
func ParseGetStocksResponse(rsp *http.Response) (*GetStocksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetStocksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Stocks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
