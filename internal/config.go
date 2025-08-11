package internal

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func getCryptoSymbols() []string {
	cwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/config/symbols.yaml", cwd)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return viper.GetStringSlice("symbols.crypto")
}

func getStockIsin() []string {
	cwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/config/symbols.yaml", cwd)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return viper.GetStringSlice("symbols.stock")
}

func getCmcToken() string {
	return os.Getenv("CMC_API_KEY")
}

func getFinnhubToken() string {
	return os.Getenv("FINNHUB_API_KEY")
}
