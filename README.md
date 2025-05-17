# Trade-simulator
A high-performance, real-time trade simulator that uses OKX L2 Order Book data via WebSocket to estimate transaction costs, slippage, fees, market impact, and internal latency. Built using Golang for the backend and React + Vite + JavaScript for the frontend.

![Screenshot 2025-05-17 232446](https://github.com/user-attachments/assets/3eab5cb8-8604-46e9-a99e-60bb842d4e11)

## 📌 Features
✅ Real-time WebSocket connection to OKX (L2 Order Book).  
✅ Slippage estimation using Linear/Quantile Regression.  
✅ Fee modeling based on OKX fee tiers.  
✅ Market impact prediction using Almgren-Chriss model.  
✅ Maker/Taker classification using Logistic Regression.  
✅ Internal latency tracking (per tick).  
✅ Interactive UI with parameter input + live metrics.  
✅ CSV logging for processed data.  
✅ Performance-optimized backend with concurrency.  

## 🧩 Project Structure

/trade-simulator  
├── backend (Golang)  
│   ├── main.go  
│   ├── ws/                 # WebSocket client & processor  
│   ├── models/             # Financial & statistical models  
│   ├── handlers/           # HTTP API handlers  
│   ├── utils/              # CSV logging, helpers  
├── frontend (React + Vite)  
│   ├── src/  
│   │   ├── components/     # InputForm, OutputDisplay, OrderBook  
│   │   ├── hooks/          # useWebSocket  
│   │   ├── App.jsx  
│   └── vite.config.js  

## 🧠 Models & Methodologies
| Output                | Model/Method Used                         |  
| --------------------- | ----------------------------------------- |  
| **Slippage**          | Linear or Quantile Regression             |  
| **Fees**              | Rule-based (based on OKX tiers)           |  
| **Market Impact**     | Almgren-Chriss optimal execution model    |  
| **Net Cost**          | Slippage + Fees + Market Impact           |  
| **Maker/Taker Ratio** | Logistic Regression                       |  
| **Latency**           | Time delta between tick receive & process |  

## 🛠️ Setup & Run
### 🔧 Backend (Golang)
cd backend  
go mod tidy  
go run main.go  

### 💻 Frontend (React + Vite)
cd frontend  
npm install  
npm run dev  

💡 Access the UI at: http://localhost:5173

### 🧪 WebSocket L2 Order Book Endpoint
wss://ws.gomarket-cpp.goquant.io/ws/l2-orderbook/okx/BTC-USDT-SWAP

## 🤓 References

OKX WebSocket Docs  
Almgren-Chriss Model Overview  
Regression Techniques for Market Modeling  

## 👨‍💻 Author
Ritwika Raj  
Btech Final year | Software Developer | Golang Enthusiast  
Email: rajritwika181@gmail.com


