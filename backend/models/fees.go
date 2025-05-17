package models

import (
	"strconv"
)

func EstimatesFees(price, quantity, feeRate float64) float64 {
	return price * quantity * feeRate
}

func ParsePrice(raw string) float64 {
	price, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0.0 // fallback or you could log fatal
	}
	return price
}
