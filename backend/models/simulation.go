package models

import (
	"math/rand"
	"time"
)

type SimulationResult struct {
	Price    float64 `json:"price"`
	Slippage float64 `json:"slippage"`
	Fees     float64 `json:"fees"`
	Impact   float64 `json:"impact"`
	NetCost  float64 `json:"netCost"`
	Latency  float64 `json:"latency"`
}

func SimulateTrade(symbol string, usdAmount, volatility, feeRate float64) SimulationResult {
	start := time.Now()

	price := 95450.0 + rand.Float64()*10
	slippage := 0.001 * usdAmount
	fees := usdAmount * feeRate
	impact := volatility * 100
	netCost := slippage + fees + impact

	latency := float64(time.Since(start).Milliseconds())

	return SimulationResult{
		Price:    price,
		Slippage: slippage,
		Fees:     fees,
		Impact:   impact,
		NetCost:  netCost,
		Latency:  latency,
	}
}
