import React, { useState, useEffect, useRef } from 'react';

function Search({ setResults }) {
  const [query, setQuery] = useState('');
  const [loading, setLoading] = useState(false);
  const timeoutRef = useRef(null);

  useEffect(() => {
    if (!query.trim()) {
      setResults([]);
      setLoading(false);
      return;
    }

    if (timeoutRef.current) clearTimeout(timeoutRef.current);

    setLoading(true);
    timeoutRef.current = setTimeout(() => {
      fetch(`/api/search?q=${encodeURIComponent(query)}`)
        .then(res => res.json())
        .then(data => {
          setResults(Array.isArray(data.results) ? data.results : []);
          setLoading(false);
        })
        .catch(() => {
          setResults([]);
          setLoading(false);
        });
    }, 500);

    return () => clearTimeout(timeoutRef.current);
  }, [query, setResults]);

  return (
    <div id='searchbar' style={{ position: 'relative', marginBottom: '5%' }}>
        <input
          type="text"
          placeholder="Поиск..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          style={{
            width: '100%',
            padding: '0.75rem 2.5rem 0.75rem 1rem',
            minWidth: 0,
            backgroundColor: '#1f1f1f',
            color: '#fff',
            border: '1px solid #444',
            borderRadius: '0.5rem',
            fontSize: '1rem',
            outline: 'none',
            boxSizing: 'border-box',
          }}
        />
        <svg
          style={{
            position: 'absolute',
            right: '0.75rem',
            top: '50%',
            transform: 'translateY(-50%)',
            color: '#888',
            pointerEvents: 'none',
            width: 18,
            height: 18,
            strokeWidth: 2,
            stroke: 'currentColor',
            fill: 'none',
          }}
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <circle cx="11" cy="11" r="8" />
          <line x1="21" y1="21" x2="16.65" y2="16.65" />
        </svg>
      </div>
  );
}
export default Search;
