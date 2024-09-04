package models

// CurrencyRate holds the exchange rates for various currencies against a base currency.
type CurrencyRate struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

// ConversionResult holds the result of a currency conversion operation.
type ConversionResult struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Result float64 `json:"result"`
}
