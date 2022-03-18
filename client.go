// package iexcloud implements a client to receive data from the iexcloud API.
package iexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// environment is used to set the endpoint.
type environment string

const (
	// Prod is used for real world data and apps running in prod.
	Prod environment = "cloud"
	// Sandbox is used for testing purposes.
	Sandbox environment = "sandbox"
)

const (
	fmtBaseURL    = "https://%s.iexapis.com/stable/"
	fmtTokenParam = "?token=%s"
	symbolsParam  = "&symbols=%s"
	fmtQuote      = "stock/%s/quote"
	fmtIntraday   = "stock/%s/intraday-prices"
	fmtHistorical = "stock/%s/chart"
	fmtDividend   = "stock/%s/dividends"
	timeSeries    = "time-series/"
)

type IexError struct {
	Code    int
	Message string
}

func (i *IexError) Error() string {
	return fmt.Sprintf("status code: %v => message: %s", i.Code, i.Message)
}

// Client holds the information and functions to make API calls.
type Client struct {
	tokenParam string
	baseURL    string
	httpClient *http.Client
}

// New returns a new client.
func New(apiKey string, env environment) *Client {
	return &Client{
		tokenParam: fmt.Sprintf(fmtTokenParam, apiKey),
		baseURL:    fmt.Sprintf(fmtBaseURL, env),
		httpClient: http.DefaultClient,
	}
}

func (c *Client) callAndUnmarshal(ctx context.Context, url string, wrapped interface{}) error {
	// log.Printf("calling iexcloud api: %s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return &IexError{Code: resp.StatusCode, Message: string(body)}
	}

	if err = json.Unmarshal(body, wrapped); err != nil {
		return err
	}

	return nil
}

// Quote returns the market information for the provided symbol.
func (c *Client) Quote(ctx context.Context, symbol string) (*QuoteResult, error) {
	callURL := c.baseURL + fmt.Sprintf(fmtQuote, symbol) + c.tokenParam

	res := &QuoteResult{}
	err := c.callAndUnmarshal(ctx, callURL, res)

	return res, err
}

// IntradayPrices provides Intraday information for a symbol.
func (c *Client) IntradayPrices(ctx context.Context, symbol string) (*IntradayResult, error) {
	callURL := c.baseURL + fmt.Sprintf(fmtIntraday, symbol) + c.tokenParam

	res := []IntradayItem{}
	err := c.callAndUnmarshal(ctx, callURL, &res)

	return &IntradayResult{Symbol: symbol, Data: res}, err
}

func (c *Client) HistoricalPrices(ctx context.Context, symbol string, duration Duration) (*IntradayResult, error) {

	callURL := c.baseURL + fmt.Sprintf(fmtHistorical, symbol)
	callURL += duration.String()

	callURL += c.tokenParam

	res := []IntradayItem{}
	err := c.callAndUnmarshal(ctx, callURL, &res)

	return &IntradayResult{Symbol: symbol, Data: res}, err
}

func (c *Client) Search(ctx context.Context, query string) (*SearchResults, error) {
	callURL := c.baseURL + "search/" + query + c.tokenParam

	res := []SearchResultItem{}
	err := c.callAndUnmarshal(ctx, callURL, &res)

	return &SearchResults{Search: query, Results: res}, err
}

// Dividends returns the basic dividend information
func (c *Client) Dividends(ctx context.Context, symbol string, duration Duration) (*DividendResults, error) {
	callURL := c.baseURL + path.Join(fmt.Sprintf(fmtDividend, symbol), duration.String())
	callURL += c.tokenParam

	res := []DividendResultItem{}
	err := c.callAndUnmarshal(ctx, callURL, &res)

	return &DividendResults{Symbol: symbol, Dividends: res}, err
}

// TimeSeries returns the data from a user provided id, key, subkey triplet in wrapped. This can be used
// for otherwise unsupported id endpoints.
func (c *Client) TimeSeries(ctx context.Context, id, key, subkey string, wrapped interface{}, queryOpts ...queryOption) error {
	callURL := c.baseURL + path.Join(timeSeries, id, key, subkey)
	opts := c.tokenParam
	for _, o := range queryOpts {
		opts += "&" + string(o)
	}
	callURL += opts
	return c.callAndUnmarshal(ctx, callURL, wrapped)
}

// News returns news entries for the symbol and options provided.
func (c *Client) News(ctx context.Context, symbol string, queryOpts ...queryOption) (*NewsResults, error) {
	res := []NewsResultItem{}
	if err := c.TimeSeries(ctx, "news", symbol, "", &res, queryOpts...); err != nil {
		return nil, err
	}
	return &NewsResults{Symbol: symbol, News: res}, nil
}

func (c *Client) AdvancedDividends(ctx context.Context, symbol string, queryOpts ...queryOption) (*DividendResults, error) {
	res := []DividendResultItem{}
	if err := c.TimeSeries(ctx, "advanced_dividends", symbol, "", &res, queryOpts...); err != nil {
		return nil, err
	}
	return &DividendResults{Symbol: symbol, Dividends: res}, nil
}
