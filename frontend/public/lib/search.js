import { refreshProducts } from './products.js';

export function initSearch(input) {
  let timeout = null;

  input.addEventListener('input', () => {
    const query = input.value.trim();
    console.log('пытаемся найти', query);

    if (timeout) clearTimeout(timeout);

    timeout = setTimeout(() => {
      refreshProducts({ name: query, page: 1, amount: 20 });
    }, 500);
  });
}
