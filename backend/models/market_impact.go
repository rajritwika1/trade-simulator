package models

func EstimatesMarketImpact(volume, volatility float64) float64 {
	return 0.001 * volume * volatility
}
