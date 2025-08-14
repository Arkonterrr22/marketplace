/**
 * 
 * @returns {Object.JSON}
 */
export async function Search({ filter = {}, query = '', page = 1, amount = 20 } = {}) {
  const res = await fetch('/api/search', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ query, page, amount }),
  });
  const data = await res.json() || {'results': [], 'warn':'no data!'};
  console.log(data)
  return data
}

// import { refreshProducts } from './products.js';

// export function initSearch(input) {
//   let timeout = null;

//   input.addEventListener('input', () => {
//     const query = input.value.trim();
//     console.log('пытаемся найти', query);

//     if (timeout) clearTimeout(timeout);

//     timeout = setTimeout(() => {
//       refreshProducts({ name: query, page: 1, amount: 20 });
//     }, 500);
//   });
// }


