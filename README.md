# Short Description
This library provides a client to retrieve data from the [Iexcloud](https://www.iexcloud.com) API.
The client only supports a subset of the available API endpoints. Namely,

* **Quote** to retrieve symbol information at this point in time.
* **IntradayPrices** to retrieve symbol information for the last 24 hours. Resolution is dependend on security type.
* **HistoricalPrices** to retrieve symbol information for a longer period of time.
* **Search** queries the search endpoint.
* **Dividends** returns basic dividend information (only supports US securities).
* **AdvancedDividends** returns detailed dividend information (supports International securities).
* **News** returns news information for the symbol.
* **TimeSeries** provides a wrapper function to query a TimeSeries based API point.

More information on the API itself can be found [here](https://iexcloud.io/docs/api/).