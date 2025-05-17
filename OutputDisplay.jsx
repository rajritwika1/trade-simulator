//import { p } from "framer-motion/client";
import React from "react";

const OutputDisplay = ({ data }) => {
    if (!data) return <p>No Data yet.</p>;

    return (
        <div className="output-display">
            <h2>Simulation Output</h2>
            <p><strong>Price:</strong>${data.price}</p>
            <p><strong>Slippage:</strong>${data.slippage}</p>
            <p><strong>Fees:</strong>${data.fees}</p>
            <p><strong>Market Impact:</strong>${data.impact}</p>
            <p><strong>Net Cost:</strong>${data.netCost}</p>
            <p><strong>Latency:</strong>{data.latency} ms</p>
        </div>
    );
};

export default OutputDisplay;