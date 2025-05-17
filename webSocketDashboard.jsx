import React, { useEffect, useState } from 'react';

function WebSocketDashboard() {
  const [tickData, setTickData] = useState(null);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws/live');

    socket.onmessage = (event) => {
    const raw = JSON.parse(event.data);
      if (raw?.arg?.channel === "books" && raw?.data?.length > 0) {
        const book = raw.data[0];
        const bestBid = book.bids[0];
        const bestAsk = book.asks[0];

        setTickData({
          bidPrice: bestBid[0],
          bidQty: bestBid[1],
          askPrice: bestAsk[0],
          askQty: bestAsk[1],
          ts: new Date().toLocaleTimeString(),
        });
      }  
    
    };

    socket.onerror = (error) => console.error("WebSocket error:", error);

    return () => socket.close();
  }, []);

  if (!tickData) return <p>Waiting for live data...</p>;

  return (
    <div className="output-panel">
      <h2>Live Trading Dashboard</h2>   
      <p><strong>Time:</strong> {tickData.ts}</p>
      <p><strong>Best Bid:</strong> {tickData.bidPrice} ({tickData.bidQty})</p>
      <p><strong>Best Ask:</strong> {tickData.askPrice} ({tickData.askQty})</p>
    
    </div>
  );
}

export default WebSocketDashboard;
