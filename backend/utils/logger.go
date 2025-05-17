package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

var csvWriter *csv.Writer
var file *os.File

func InitCSVLogger() {
	var err error
	file, err = os.Create("logs.csv")
	if err != nil {
		panic("Failed to create CSV log file: " + err.Error())
	}
	csvWriter = csv.NewWriter(file)
	csvWriter.Write([]string{
		"timestamp", "symbol", "price", "quantity", "slippage", "fees", "impact", "net_cost", "latency_ms",
	})
}

func LogTick(timestamp, symbol string, price, qty, slippage, fees, impact, net, latency float64) {
	csvWriter.Write([]string{
		timestamp,
		symbol,
		formatFloat(price),
		formatFloat(qty),
		formatFloat(slippage),
		formatFloat(fees),
		formatFloat(impact),
		formatFloat(net),
		formatFloat(latency),
	})
	csvWriter.Flush()
}

func CloseCSVLogger() {
	if file != nil {
		file.Close()
	}
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}
