import React, { useState } from 'react';
import Select from 'react-select';

// Currency options for the multi-select dropdown
export const currencyOptions = [
  { value: 'USD', label: 'United States Dollar (USD)' },
  { value: 'EUR', label: 'Euro (EUR)' },
  { value: 'JPY', label: 'Japanese Yen (JPY)' },
  { value: 'GBP', label: 'British Pound Sterling (GBP)' },
  { value: 'AUD', label: 'Australian Dollar (AUD)' },
  { value: 'CAD', label: 'Canadian Dollar (CAD)' },
  { value: 'CHF', label: 'Swiss Franc (CHF)' },
  { value: 'CNY', label: 'Chinese Yuan (CNY)' },
  { value: 'SEK', label: 'Swedish Krona (SEK)' },
  { value: 'NZD', label: 'New Zealand Dollar (NZD)' }
];

function Rates() {
  const [rates, setRates] = useState({});
  const [base, setBase] = useState('USD');
  const [targets, setTargets] = useState([]);

  const fetchRates = async () => {
    const targetCurrencies = targets.map(target => target.value).join(',');
    try {
      const response = await fetch(`/rates?base=${base}&target=${targetCurrencies}`);
      const data = await response.json();
      console.log('Data:', data);
      //console.log(data);
      setRates(data.rates);
    } catch (error) {
      console.error('Error fetching rates:', error);
    }
  };

  return (
    <div>
      <h2>Exchange Rates</h2>
      <div>
        <label>
          Base Currency:
          <input type="text" value={base} onChange={(e) => setBase(e.target.value)} />
        </label>
        <label>
          Target Currencies:
          <Select
            isMulti
            options={currencyOptions}
            value={targets}
            onChange={setTargets}
            className="basic-multi-select"
            classNamePrefix="select"
          />
        </label>
        <button onClick={fetchRates}>Fetch Rates</button>
      </div>
      {Object.keys(rates).length > 0 && (
        <div>
          <h3>Rates:</h3>
          <table>
            <thead>
              <tr>
                <th>Currency</th>
                <th>Rate</th>
              </tr>
            </thead>
            <tbody>
              {Object.entries(rates).map(([currency, rate]) => (
                <tr key={currency}>
                  <td>{currency}</td>
                  <td>{rate.toFixed(4)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

export default Rates;
