package api

import (
	"log"
	"server/internal/models" // Adjust the import path based on your project structure
)

// Define default fixed exchange rates (USD as base)
var defaultRates = map[string]float64{
	"USD": 1.00,  // Base currency, USD to USD rate is obviously 1
	"EUR": 0.85,  // Euro
	"JPY": 110.0, // Japanese Yen
	"GBP": 0.75,  // British Pound Sterling
	"AUD": 1.40,  // Australian Dollar
	"CAD": 1.25,  // Canadian Dollar
	"CHF": 0.91,  // Swiss Franc
	"CNY": 6.45,  // Chinese Yuan
	"SEK": 8.75,  // Swedish Krona
	"NZD": 1.50,  // New Zealand Dollar
}

// fetchDefaultRates retrieves rates using the internal defaults based on the target currencies.
func fetchDefaultRates(base string, targets []string) models.CurrencyRate {
	// Assume base is USD for simplicity
	filteredRates := make(map[string]float64)

	log.Printf("Fetching default rates for base: %s with targets: %v", base, targets)

	if len(targets) == 0 {
		log.Println("No specific targets provided, returning all default rates.")
		return models.CurrencyRate{
			Success: true,
			Base:    base,
			Rates:   defaultRates,
		}
	}

	for _, target := range targets {
		if rate, ok := defaultRates[target]; ok {
			filteredRates[target] = rate
			log.Printf("Added rate for %s: %f", target, rate)
		} else {
			log.Printf("No default rate found for %s", target)
		}
	}

	if len(filteredRates) == 0 {
		log.Println("No rates found for the given targets, returning empty rates map.")
	} else {
		log.Printf("Returning filtered rates: %v", filteredRates)
	}

	return models.CurrencyRate{
		Success: true,
		Base:    base,
		Rates:   filteredRates,
	}
}
