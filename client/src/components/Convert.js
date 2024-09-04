import React, { useState } from 'react';
import Select from 'react-select';

// Currency options for the dropdowns
const currencyOptions = [
  { value: 'USD', label: 'United States Dollar (USD)' },
  { value: 'EUR', label: 'Euro (EUR)' },
  { value: 'GBP', label: 'British Pound (GBP)' },
  { value: 'JPY', label: 'Japanese Yen (JPY)' },
  { value: 'CAD', label: 'Canadian Dollar (CAD)' },
  { value: 'AUD', label: 'Australian Dollar (AUD)' },
];

function Convert() {
  const [from, setFrom] = useState({ value: 'USD', label: 'United States Dollar (USD)' });
  const [to, setTo] = useState({ value: 'EUR', label: 'Euro (EUR)' });
  const [amount, setAmount] = useState(100);
  const [convertedAmount, setConvertedAmount] = useState(null);

  const convertCurrency = async () => {
    try {
      const response = await fetch(`/convert?from=${from.value}&to=${to.value}&amount=${amount}`);
      const data = await response.json();
      setConvertedAmount(data.convertedAmount);
    } catch (error) {
      console.error('Error converting currency:', error);
    }
  };

  return (
    <div>
      <h2>Currency Converter</h2>
      <div>
        <label>
          From Currency:
          <Select
            options={currencyOptions}
            value={from}
            onChange={setFrom}
            className="basic-single-select"
            classNamePrefix="select"
          />
        </label>
        <label>
          To Currency:
          <Select
            options={currencyOptions}
            value={to}
            onChange={setTo}
            className="basic-single-select"
            classNamePrefix="select"
          />
        </label>
        <label>
          Amount:
          <input type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
        </label>
        <button onClick={convertCurrency}>Convert</button>
      </div>
      {convertedAmount !== null && (
        <div>
          <h3>Converted Amount: ${convertedAmount}</h3>
        </div>
      )}
    </div>
  );
}

export default Convert;
