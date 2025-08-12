package internal

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func TableRender() {
	renderCrypto()
	renderStock()
}

func renderCrypto() {
	table := tablewriter.NewWriter(os.Stdout)
	defer table.Render()
	tableHeader := []string{"Symbol", "Price", "+/- 1H", "+/- 24H"}
	table.Header(tableHeader)

	for _, cd := range getCryptoData() {
		priceRounded := strconv.FormatFloat(float64(cd.QuoteData.Price), 'f', 2, 32)
		cryptoDifference1Hour := strconv.FormatFloat(float64(cd.QuoteData.PercentChange1H), 'f', 2, 32)
		cryptoDifference24Hour := strconv.FormatFloat(float64(cd.QuoteData.PercentChange24H), 'f', 2, 32)
		result := []any{cd.Name, priceRounded, __colorizeStockPriceChange(cryptoDifference1Hour), __colorizeStockPriceChange(cryptoDifference24Hour)}
		table.Append(result...)
	}
}

func renderStock() {
	table := tablewriter.NewWriter(os.Stdout)
	defer table.Render()
	tableHeader := []string{"Symbol", "Open", "+/-", "High", "Low", "Previous Close"}
	table.Header(tableHeader)

	for _, sd := range getStockData() {
		stockDifference := strconv.FormatFloat(float64(*sd.Quote.D), 'f', 2, 32)
		result := []any{sd.Name, *sd.Quote.O, __colorizeStockPriceChange(stockDifference), *sd.Quote.H, *sd.Quote.L, *sd.Quote.Pc}
		table.Append(result...)
	}
}

func __colorizeStockPriceChange(change string) string {
	if strings.Contains(change, "-") {
		return "\033[31m" + change + "\033[0m"
	} else {
		return "\033[32m" + change + "\033[0m"
	}
}
