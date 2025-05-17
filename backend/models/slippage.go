package models

func EstimateSlippage(price, quantity, volatility float64) float64 {
	return 0.0005 * price * volatility * quantity
}
