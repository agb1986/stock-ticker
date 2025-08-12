package main

import (
	"fmt"
	"time"
	// "github.com/olekukonko/tablewriter"
)

var STOCKS = []string{"RGTI", "INTC"}

func main() {
	now := time.Now()
	fmt.Printf("\nFetching stocks @ %s\n", now.Format(time.DateTime))
}

func getStockQuotes() {
	// table := tablewriter.NewWriter(os.Stdout)
	// defer table.Render()
	// tableHeader := []string{"Symbol", "Open", "+/-", "High", "Low", "Previous Close"}
	// table.Header(tableHeader)
	// 	stockDifference := strconv.FormatFloat(float64(*quote.D), 'f', 2, 32)
	// 	result := []any{stock, *quote.O, colorizeStockPriceChange(stockDifference), *quote.H, *quote.L, *quote.Pc}
	// 	table.Append(result...)
}

// func colorizeStockPriceChange(change string) string {
// 	if strings.Contains(change, "-") {
// 		return "\033[31m" + change + "\033[0m"
// 	} else {
// 		return "\033[32m" + change + "\033[0m"
// 	}
// }
