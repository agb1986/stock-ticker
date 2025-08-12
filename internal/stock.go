package internal

import (
	"context"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func __setupFinnhubClient() *finnhub.DefaultApiService {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", getFinnhubToken())
	return finnhub.NewAPIClient(cfg).DefaultApi
}

// https://finnhub.io/docs/api/symbol-search
func __getStockSymbolsByIsin() []string {
	var symbols []string

	finnhubClient := __setupFinnhubClient()
	for _, item := range getStockIsin() {
		searchResult, _, err := finnhubClient.SymbolSearch(context.Background()).Q(item).Execute()
		handleError(err, "__getStockSymbolsByIsin")
		symbols = append(symbols, *searchResult.GetResult()[0].Symbol)
	}

	return symbols
}

// https://finnhub.io/docs/api/quote
func getStockData() []StockData {
	var stockData []StockData

	finnhubClient := __setupFinnhubClient()
	for _, item := range __getStockSymbolsByIsin() {
		quote, _, err := finnhubClient.Quote(context.Background()).Symbol(item).Execute()
		handleError(err, "getStockData")
		var sd StockData
		sd.Name = item
		sd.Quote = quote
		stockData = append(stockData, sd)
	}

	return stockData
}

type StockData struct {
	Name  string
	Quote finnhub.Quote
}
