package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Hardcode useAPI to false
var useAPI = false

func RatesHandler(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	if base == "" {
		base = "USD" // Default to USD if no base is specified
	}
	targetParam := r.URL.Query().Get("target") // Get the 'target' query parameter
	targets := strings.Split(targetParam, ",") // Split the string into a slice by comma

	if len(targets) == 1 && targets[0] == "" {
		targets = []string{} // Handle the case where no target is specified
	}

	log.Printf("Received base: %s, targets: %v", base, targets) // Debugging log

	rates, err := FetchRates(base, targets, useAPI)
	if err != nil {
		http.Error(w, "Failed to fetch rates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amountStr := r.URL.Query().Get("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	log.Printf("Converting from %s to %s amount %f", from, to, amount) // Debugging log

	rates, err := FetchRates(from, []string{to}, useAPI)
	if err != nil {
		http.Error(w, "Failed to fetch rates", http.StatusInternalServerError)
		return
	}

	rate, exists := rates.Rates[to]
	if !exists {
		http.Error(w, "Conversion rate not found", http.StatusBadRequest)
		return
	}

	convertedAmount := amount * rate
	result := map[string]float64{"convertedAmount": convertedAmount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
