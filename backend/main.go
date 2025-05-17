package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trade-simulator/models"
	"trade-simulator/utils"
	"trade-simulator/ws"
)

var (
	symbol     = flag.String("symbol", "BTC-USDT-SWAP", "Trading symbol")
	volatility = flag.Float64("vol", 0.02, "Market volatility")
	feeRate    = flag.Float64("fee", 0.001, "Fee rate")
	usdAmount  = flag.Float64("usd", 100.0, "USD trade amount")
)

type SimulationRequest struct {
	Symbol     string  `json:"symbol"`
	UsdAmount  float64 `json:"usdAmount"`
	Volatility float64 `json:"volatility"`
	FeeRate    float64 `json:"feeRate"`
}

type SimulationResponse struct {
	Price    float64 `json:"price"`
	Slippage float64 `json:"slippage"`
	Fees     float64 `json:"fees"`
	Impact   float64 `json:"impact"`
	NetCost  float64 `json:"netCost"`
	Latency  float64 `json:"latency"`
}

func handleSimulate(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var req SimulationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	/*resp := SimulationResponse{
		Price:    95450.5,
		Slippage: 1.25,
		Fees:     req.UsdAmount * req.FeeRate,
		Impact:   req.Volatility * 100,
		NetCost:  1.25 + (req.UsdAmount * req.FeeRate) + (req.Volatility * 100),
		Latency:  float64(time.Since(start).Milliseconds()),
	}
	json.NewEncoder(w).Encode(resp)*/
	result := models.SimulateTrade(
		req.Symbol,
		req.UsdAmount,
		req.Volatility,
		req.FeeRate,
	)
	result.Latency = float64(time.Since(start).Microseconds())

	json.NewEncoder(w).Encode(result)

}

func main() {
	flag.Parse()
	utils.InitCSVLogger()
	defer utils.CloseCSVLogger()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		utils.CloseCSVLogger()
		os.Exit(1)
	}()

	go ws.Start(*symbol, *usdAmount, *volatility, *feeRate)

	http.HandleFunc("/simulate", handleSimulate)
	log.Println("Backend started: WebSocket + HTTP API running on : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/ws/live", ws.HandleWebSocket)
}
