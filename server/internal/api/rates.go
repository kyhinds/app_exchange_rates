package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/internal/models"
)

// FetchRates attempts to get rates from the API and falls back to default rates if necessary.
func FetchRates(base string, targets []string, useAPI bool) (models.CurrencyRate, error) {
	if useAPI {
		rates, err := fetchFromAPI(base, targets)
		if err == nil {
			return rates, nil
		}
		log.Printf("API call failed: %v, using default rates.", err)
	}
	return fetchDefaultRates(base, targets), nil
}

// fetchFromAPI fetches rates from an external API.
func fetchFromAPI(base string, targets []string) (models.CurrencyRate, error) {
	apiKey := "your_api_key_here"
	url := fmt.Sprintf("https://api.exchangeratesapi.io/v1/latest?access_key=%s&base=%s", apiKey, base)
	resp, err := http.Get(url)
	if err != nil {
		return models.CurrencyRate{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CurrencyRate{}, err
	}

	var rates models.CurrencyRate
	if err := json.Unmarshal(body, &rates); err != nil {
		return models.CurrencyRate{}, err
	}
	rates = filterRates(rates, targets)
	return rates, nil
}

// filterRates filters the fetched rates to only include the specified target currencies.
func filterRates(rates models.CurrencyRate, targets []string) models.CurrencyRate {
	filteredRates := make(map[string]float64)
	for _, currency := range targets {
		if rate, found := rates.Rates[currency]; found {
			filteredRates[currency] = rate
		}
	}
	rates.Rates = filteredRates
	return rates
}
