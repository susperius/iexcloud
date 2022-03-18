package iexcloud

import (
	"fmt"
	"time"
)

const (
	day   = 24 * time.Hour
	month = 30 * day
	year  = 365 * day

	timeFmt = "2006-01-02 15:04:05"
)

// QuoteResult holds the information about a symbol quote.
type QuoteResult struct {
	Symbol                 string  `json:"symbol"`
	CompanyName            string  `json:"companyName"`
	PrimaryExchange        string  `json:"primaryExchange"`
	CalculationPrice       string  `json:"calculationPrice"`
	Open                   float64 `json:"open"`
	OpenTime               int64   `json:"openTime"`
	OpenSource             string  `json:"openSource"`
	Close                  float64 `json:"close"`
	CloseTime              int64   `json:"closeTime"`
	CloseSource            string  `json:"closeSource"`
	High                   float64 `json:"high"`
	HighTime               int64   `json:"highTime"`
	HighSource             string  `json:"highSource"`
	Low                    float64 `json:"low"`
	LowTime                int64   `json:"lowTime"`
	LowSource              string  `json:"lowSource"`
	LatestPrice            float64 `json:"latestPrice"`
	LatestSource           string  `json:"latestSource"`
	LatestTime             string  `json:"latestTime"`
	LatestUpdate           int64   `json:"latestUpdate"`
	LatestVolume           int     `json:"latestVolume"`
	IexRealtimePrice       float64 `json:"iexRealtimePrice"`
	IexRealtimeSize        int     `json:"iexRealtimeSize"`
	IexLastUpdated         int64   `json:"iexLastUpdated"`
	DelayedPrice           float64 `json:"delayedPrice"`
	DelayedPriceTime       int64   `json:"delayedPriceTime"`
	OddLotDelayedPrice     float64 `json:"oddLotDelayedPrice"`
	OddLotDelayedPriceTime int64   `json:"oddLotDelayedPriceTime"`
	ExtendedPrice          float64 `json:"extendedPrice"`
	ExtendedChange         float64 `json:"extendedChange"`
	ExtendedChangePercent  float64 `json:"extendedChangePercent"`
	ExtendedPriceTime      int64   `json:"extendedPriceTime"`
	PreviousClose          float64 `json:"previousClose"`
	PreviousVolume         int     `json:"previousVolume"`
	Change                 float64 `json:"change"`
	ChangePercent          float64 `json:"changePercent"`
	Volume                 int     `json:"volume"`
	IexMarketPercent       float64 `json:"iexMarketPercent"`
	IexVolume              int     `json:"iexVolume"`
	AvgTotalVolume         int     `json:"avgTotalVolume"`
	IexBidPrice            float64 `json:"iexBidPrice"`
	IexBidSize             int     `json:"iexBidSize"`
	IexAskPrice            float64 `json:"iexAskPrice"`
	IexAskSize             int     `json:"iexAskSize"`
	IexOpen                float64 `json:"iexOpen"`
	IexOpenTime            int64   `json:"iexOpenTime"`
	IexClose               float64 `json:"iexClose"`
	IexCloseTime           int64   `json:"iexCloseTime"`
	MarketCap              int64   `json:"marketCap"`
	PeRatio                float64 `json:"peRatio"`
	Week52High             float64 `json:"week52High"`
	Week52Low              float64 `json:"week52Low"`
	YtdChange              float64 `json:"ytdChange"`
	LastTradeTime          int64   `json:"lastTradeTime"`
	Currency               string  `json:"currency"`
	IsUSMarketOpen         bool    `json:"isUSMarketOpen"`
}

// IntradayResult holds the symbol and the intraday items.
type IntradayResult struct {
	Symbol string
	Data   []IntradayItem
}

// IntradayItem holds a single intraday entry.
type IntradayItem struct {
	Average        float64 `json:"average"`
	ChangeOverTime float64 `json:"changeOverTime"`
	Close          float64 `json:"close"`
	Date           string  `json:"date"`
	High           float64 `json:"high"`
	Label          string  `json:"label"`
	Low            float64 `json:"low"`
	Minute         string  `json:"minute"`
	Notional       float64 `json:"notional"`
	NumberOfTrades int64   `json:"numberOfTrades"`
	Open           float64 `json:"open"`
	Volume         float64 `json:"volume"`
}

// SearchResults holds the search string and the received results.
type SearchResults struct {
	Search  string
	Results []SearchResultItem
}

// SearchResultItem holds a single search result.
type SearchResultItem struct {
	Cik            string `json:"cik"`
	Currency       string `json:"currency"`
	Exchange       string `json:"exchange"`
	ExchangeName   string `json:"exchangeName"`
	ExchangeSuffix string `json:"exchangeSuffix"`
	Figi           string `json:"figi"`
	IexID          string `json:"iexId"`
	Lei            string `json:"lei"`
	Name           string `json:"name"`
	Region         string `json:"region"`
	Sector         string `json:"sector"`
	SecurityName   string `json:"securityName"`
	SecurityType   string `json:"securityType"`
	Symbol         string `json:"symbol"`
	Type           string `json:"type"`
}

// NewsResultItem holds a single news result.
type NewsResultItem struct {
	Datetime int64 `json:"datetime"`
	// HasPaywall string `json:"hasPaywall"`
	Headline string `json:"headline"`
	Image    string `json:"image"`
	ImageURL string `json:"imageUrl"`
	Lang     string `json:"lang"`
	Provider string `json:"provider"`
	QmURL    string `json:"qmUrl"`
	Related  string `json:"related"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
	URL      string `json:"url"`
}

// NewsResults holds the symbol and the news received.
type NewsResults struct {
	Symbol string
	News   []NewsResultItem
}

// DividendResultItem holds information for a single dividend event.
type DividendResultItem struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	DeclaredDate string  `json:"declaredDate"`
	Description  string  `json:"description"`
	ExDate       string  `json:"exDate"`
	Flag         string  `json:"flag"`
	Frequency    string  `json:"frequency"`
	PaymentDate  string  `json:"paymentDate"`
	RecordDate   string  `json:"recordDate"`
	// the following entries are only available from the advanced dividends endpoint.
	SecurityType          string  `json:"securityType,omitempty"`
	Notes                 string  `json:"notes,omitempty"`
	Figi                  string  `json:"figi,omitempty"`
	LastUpdated           string  `json:"lastUpdated,omitempty"`
	CountryCode           string  `json:"countryCode,omitempty"`
	ParValue              float64 `json:"parValue,omitempty"`
	ParValueCurrency      string  `json:"parValueCurrency,omitempty"`
	NetAmount             float64 `json:"netAmount,omitempty"`
	GrossAmount           float64 `json:"grossAmount,omitempty"`
	Marker                string  `json:"marker,omitempty"`
	TaxRate               float64 `json:"taxRate,omitempty"`
	FromFactor            float64 `json:"fromFactor,omitempty"`
	ToFactor              float64 `json:"toFactor,omitempty"`
	ADRFee                float64 `json:"adrFee,omitempty"`
	Coupon                float64 `json:"coupon,omitempty"`
	DeclaredCurrencyCD    string  `json:"declaredCurrencyCD,omitempty"`
	DeclaredGrossAmount   float64 `json:"declaredGrossAmount,omitempty"`
	IsNetInvestmentIncome bool    `json:"isNetInvestmentIncome,omitempty"`
	IsDAP                 bool    `json:"isDAP,omitempty"`
	IsApproximate         bool    `json:"isApproximate,omitempty"`
	FxDate                string  `json:"fxDate,omitempty"`
	SecondPaymentDate     string  `json:"secondPaymentDate,omitempty"`
	SecondExDate          string  `json:"secondExDate,omitempty"`
	FiscalYearEndDate     string  `json:"fiscalYearEndDate,omitempty"`
	PeriodEndDate         string  `json:"periodEndDate,omitempty"`
	OptionalElectionDate  string  `json:"optionalElectionDate,omitempty"`
	ToDate                string  `json:"toDate,omitempty"`
	RegistrationDeadline  string  `json:"registrationDeadline,omitempty"`
	InstallmentPayDate    string  `json:"installmentPayDate,omitempty"`
	// TODO: refid is a number when calling .Dividends and a string for AdancedDividends.
	// RefID                 string  `json:"refid,omitempty"`
	Created string `json:"created,omitempty"`
}

// DividendResults holds the symbol and the received dividend event information.
type DividendResults struct {
	Symbol    string
	Dividends []DividendResultItem
}

// Duration specifies how far back historical data should be queried. The largest unit value is taken.
// For example if max is true the rest of the fields don't matter.
type Duration struct {
	Max    bool
	Years  int64
	Months int64
	Days   int64
}

func (d Duration) String() string {
	switch {
	case d.Max:
		return "max"
	case d.Years != 0:
		return fmt.Sprintf("%dy", d.Years)
	case d.Months != 0:
		return fmt.Sprintf("%dm", d.Months)
	case d.Days != 0:
		return fmt.Sprintf("%dd", d.Days)
	default:
		return "1d"
	}
}

// NewDuration returns an IEXCloud duration type from a time.Duration.
func NewDuration(dur time.Duration) Duration {
	switch {
	case dur < day:
		return Duration{Days: 1}
	case dur < month:
		days := int64(dur / day)
		return Duration{Days: days}
	case dur < year:
		months := int64(dur / month)
		return Duration{Months: months}
	default:
		years := int64(dur / year)
		return Duration{Years: years}
	}
}

type queryOption string

// Range sets a TimeSeries query option for the provided range.
func Range(dur Duration) queryOption {
	return queryOption("range=" + dur.String())
}

// Calendar sets a TimeSeries query option for the calendar value.
func Calendar(value bool) queryOption {
	if value {
		return "calendar=true"
	}
	return "calendar=false"
}

// Limit sets a TimeSeries query option to restrict the number of results.
func Limit(count uint32) queryOption {
	return queryOption(fmt.Sprintf("limit=%d", count))
}

// Subattribute sets a TimeSeries query option for the provided key values.
func Subattribute(keyValues map[string]string) queryOption {
	var pairs []string
	for k, v := range keyValues {
		pairs = append(pairs, k+"|"+v)
	}
	kvList := ""
	for i, p := range pairs {
		if i == len(pairs)-1 {
			kvList += p
			break
		}
		kvList += p + ","
	}
	return queryOption("subattribute=" + kvList)
}
