package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCryptoSymbolsSuccess(t *testing.T) {
	crypto_symbols := getCryptoSymbols()
	assert.NotEmpty(t, crypto_symbols)
	assert.IsType(t, []string{}, crypto_symbols)
}

func TestGetStockIsinSuccess(t *testing.T) {
	crypto_symbols := getStockIsin()
	assert.NotEmpty(t, crypto_symbols)
	assert.IsType(t, []string{}, crypto_symbols)
}
