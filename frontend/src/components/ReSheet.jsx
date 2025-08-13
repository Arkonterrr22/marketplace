function ReSheet({ results = [], setResults }) {
  const columns = results.length > 0 ? Object.keys(results[0]) : [];

  const cellStyle = {
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    whiteSpace: 'nowrap',
    maxWidth: 150,
  };

  return (
    <table style={{ width: '100%', tableLayout: 'fixed' }}>
      <thead>
        <tr>
          {columns.map(col => (
            <th key={col} style={cellStyle}>{col}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {results.map((row, idx) => (
          <tr key={idx}>
            {columns.map(col => (
              <td key={col} style={cellStyle} title={row[col]}>
                {row[col]}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default ReSheet;