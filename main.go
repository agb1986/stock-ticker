package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/olekukonko/tablewriter"
)

var STOCKS = []string{"RGTI", "INTC"}

func main() {
	getStockQuotes()
}

func getStockQuotes() {
	var FINNHUB_API_KEY = os.Getenv("FINNHUB_API_KEY")
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", FINNHUB_API_KEY)
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	table := tablewriter.NewWriter(os.Stdout)
	defer table.Render()
	tableHeader := []string{"Open", "+/-", "High", "Low", "Previous Close"}
	table.Header(tableHeader)

	for _, stock := range STOCKS {
		quote, response, err := finnhubClient.Quote(context.Background()).Symbol(stock).Execute()
		if err != nil {
			fmt.Print(response.Status)
			panic(err)
		}
		stockDifference := strconv.FormatFloat(float64(*quote.D), 'f', 2, 32)
		result := []any{*quote.O, colorizeStockPriceChange(stockDifference), *quote.H, *quote.L, *quote.Pc}
		table.Append(result...)
	}
}

func colorizeStockPriceChange(change string) string {
	if strings.Contains(change, "-") {
		return "\033[31m" + change + "\033[0m"
	} else {
		return "\033[32m" + change + "\033[0m"
	}
}
