import React from 'react';
import './App.css';  // Import the CSS file
import Rates from './components/Rates';
import Convert from './components/Convert';

function App() {
  return (
    <div className="App">
      <h1>Currency Exchange</h1>
      <Rates />
      <Convert />
    </div>
  );
}

export default App;
