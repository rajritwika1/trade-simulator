package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"trade-simulator/models"
	"trade-simulator/utils"

	"nhooyr.io/websocket"
)

type Orderbook struct {
	Timestamp string     `json:"timestamp"`
	Exchange  string     `json:"exchange"`
	Symbol    string     `json:"symbol"`
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
}

func Start(symbol string, usdAmount, volatility, feeRate float64) {
	ctx := context.Background()
	url := "wss://ws.gomarket-cpp.goquant.io/ws/l2-orderbook/okx/" + symbol
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Fatal("websocket error:", err)
	}

	defer c.Close(websocket.StatusNormalClosure, "done")

	for {
		_, data, err := c.Read(ctx)
		if err != nil {
			log.Println("Read error:", err)
			continue
		}

		var ob Orderbook
		if err := json.Unmarshal(data, &ob); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		processTick(ob, usdAmount, volatility, feeRate)

	}
}

func processTick(ob Orderbook, usdAmount, volatility, feeRate float64) {
	if len(ob.Asks) == 0 {
		return
	}
	timer := utils.NewTimer()
	defer timer.Stop()

	price := models.ParsePrice(ob.Asks[0][0])
	quantity := usdAmount / price
	//volatility := 0.02
	//feeRate := 0.001

	slippage := models.EstimateSlippage(price, quantity, volatility)
	fees := models.EstimatesFees(price, quantity, feeRate)
	impact := models.EstimatesMarketImpact(quantity, volatility)

	netCost := slippage + fees + impact

	fmt.Printf("\n✅ Tick: %s\n", ob.Timestamp)
	fmt.Printf("🔹 Price: %.2f\n🔹 Qty: %.6f\n", price, quantity)
	fmt.Printf("🔹 Slippage: $%.4f\n🔹 Fees: $%.4f\n🔹 Impact: $%.4f\n", slippage, fees, impact)
	fmt.Printf("💰 Net Cost: $%.4f\n", netCost)
	fmt.Printf("🕒 Latency: %.2f ms\n", timer.Elapsed())

	utils.LogTick(ob.Timestamp, ob.Symbol, price, quantity, slippage, fees, impact, netCost, timer.Elapsed())

}
