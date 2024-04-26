package utils

import (
	"fmt"
	"log"
	"mmclinton/tenk/config"
	"strings"
)

const baseURL = "https://financialmodelingprep.com/api/v3/sec_filings/"
const reportType = "?type=10-k"

func ApiUrlBuilder(symbol *string) (url string) {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading the configuration file: %v", err)
	}

	apiKey := fmt.Sprint("&apikey=" + config.ApiKey)
	*symbol = strings.ToUpper(*symbol)
	*symbol = strings.TrimSpace(*symbol)
	return fmt.Sprint(baseURL, *symbol, reportType, apiKey)
}
