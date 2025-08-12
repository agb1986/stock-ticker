package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var CRYPTO_URL = "https://pro-api.coinmarketcap.com"

// https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyQuotesLatest

func getCryptoData() []CryptoData {
	client := &http.Client{}
	var cryptoData []CryptoData

	for _, item := range getCryptoSymbols() {
		// API request to CoinMarketCap for cryptocurrency data
		url := fmt.Sprintf("%s/v2/cryptocurrency/quotes/latest", CRYPTO_URL)
		req, err := http.NewRequest("GET", url, nil)
		handleError(err, "getCryptoData")
		req.Header.Add("x-cmc_pro_api_key", getCmcToken())
		req.Header.Add("accept", "application/json")
		q := req.URL.Query()
		q.Add("symbol", item)
		q.Add("convert", "GBP")
		req.URL.RawQuery = q.Encode()
		res, err := client.Do(req)
		handleError(err, "getCryptoData")
		body, err := io.ReadAll(res.Body)
		handleError(err, "getCryptoData")

		// Parse the JSON response
		var data any
		json.Unmarshal([]byte(string(body)), &data)
		jsonPath := fmt.Sprintf("$.data.%s[0].quote.GBP", item)
		parsedData := getDataByJsonPath(jsonPath, data)
		marshaledData, err := json.Marshal(parsedData)
		handleError(err, "getCryptoData")
		var quoteData QuoteData
		json.Unmarshal(marshaledData, &quoteData)
		var cd CryptoData
		cd.Name = item
		cd.QuoteData = quoteData
		cryptoData = append(cryptoData, cd)
	}

	return cryptoData
}

type CryptoData struct {
	Name      string
	QuoteData QuoteData
}

// Nested data object within CoinMarketCap response
type QuoteData struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	Tvl                   any       `json:"tvl"`
	LastUpdated           time.Time `json:"last_updated"`
}
