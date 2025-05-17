import React,{useState} from 'react';
import InputForm from "./components/InputForm";
import OutputDisplay from "./components/OutputDisplay";
import WebSocketDashboard from './components/webSocketDashboard';
import './App.css';

function App() {
  const [outputData, setOutputData] = useState(null);

  const handleSimulation = async (params) => {

   try
   {
      const res = await fetch('http://localhost:8080/simulate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
      });
      const data = await res.json();
      setOutputData(data);
    } catch (err) {
      console.error('Failed to fetch simulation:', err);
    }
    
   };
  return (
    <div className="app-container">
      <InputForm onSubmit={handleSimulation}/>
      <OutputDisplay data = {outputData}/>
      <WebSocketDashboard/>
    </div>
  );
}

export default App;