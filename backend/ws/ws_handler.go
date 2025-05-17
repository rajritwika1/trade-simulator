package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	okxConn, _, err := websocket.DefaultDialer.Dial("wss://ws.okx.com:8443/ws/v5/public", nil)
	if err != nil {
		log.Println("OKX WebSocket connection failed:", err)
		return
	}
	defer okxConn.Close()
	subscribeMsg := map[string]interface{}{
		"op": "subscribe",
		"args": []map[string]string{
			{
				"channel": "books",
				"instId":  "BTC-USDT-SWAP",
			},
		},
	}
	msgBytes, _ := json.Marshal(subscribeMsg)
	okxConn.WriteMessage(websocket.TextMessage, msgBytes)

	// Stream OKX data to client
	for {
		_, message, err := okxConn.ReadMessage()
		if err != nil {
			log.Println("Read error from OKX:", err)
			break
		}

		// Optional: Decode and filter messages here if needed

		// Forward raw message to frontend
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Write to client error:", err)
			break

		}
	}
}
