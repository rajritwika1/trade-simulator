import React, { useState } from 'react';


const InputForm = ({ onSubmit }) => {
    const [symbol , setSymbol] = useState("BTC-USDT-SWAP");
    const [usdAmount, setUsdAmount] = useState(100);
    const [volatility, setVolatility] = useState(0.02);
    const [feeRate, setFeeRate] = useState(0.001);
    
    

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit({symbol,usdAmount, volatility,feeRate});
    };

    return(
     <form onSubmit={handleSubmit} className="input-form">
      <h2>Trade Parameters</h2>
      <label>Symbol:
        <input value={symbol} onChange={(e) => setSymbol(e.target.value)} />
      </label>
      <label>USD Amount:
        <input type="number" value={usdAmount} onChange={(e) => setUsdAmount(+e.target.value)} />
      </label>
      <label>Volatility:
        <input type="number" step="0.001" value={volatility} onChange={(e) => setVolatility(+e.target.value)} />
      </label>
      <label>Fee Rate:
        <input type="number" step="0.0001" value={feeRate} onChange={(e) => setFeeRate(+e.target.value)} />
      </label>
      <button type="submit">Start Simulation</button>
    </form>
  );

};

export default InputForm;