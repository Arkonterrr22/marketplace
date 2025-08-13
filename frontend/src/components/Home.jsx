import React, { useState } from 'react';
import Searchbar from './Searchbar';
import ReSheet from './ReSheet';

export default function Home() {
  const [results, setResults] = useState([]);
  return (
        <div style={{ position: 'relative', maxWidth: '100%', width: '60%', margin: '0 auto' }}>
        <Searchbar setResults={setResults} />
        <ReSheet results={results} setResults={setResults} />
      </div>
  );
}