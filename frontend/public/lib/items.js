import { Search } from './search.js';

export function refreshItems({ filter={}, input = null, page = 1, amount = 20 } = {}) {
    const list = document.getElementById('product-list');
  if (!list) {
    console.warn('Не найден элемент с id="product-list"');
    return;
  }
  const fetchAndRender = async (query, pageNum = page, itemsPerPage = amount) => {
    const res = await Search({ filter:filter, query:query, page: pageNum, amount: itemsPerPage });
    const products = res.results || [];

    list.innerHTML = '';

    products.forEach(product => {
      const card = document.createElement('div');
      card.className = 'product-card';

      const img = document.createElement('img');
      img.src = product.image;
      img.alt = product.title;
      img.loading = 'lazy';
      card.appendChild(img);

      const info = document.createElement('div');
      info.className = 'product-info';

      const title = document.createElement('h2');
      title.className = 'product-title';
      title.textContent = product.title;
      info.appendChild(title);

      const desc = document.createElement('p');
      desc.className = 'product-desc';
      desc.textContent = product.description;
      info.appendChild(desc);

      card.appendChild(info);
      list.appendChild(card);
    });
  };
  if (!input) input = document.getElementById('search-input');
  if (!input) {
    console.warn('Не найден элемент input с id="search-input"');
    fetchAndRender('')
    return;
  }

  let timeout = null;
  input.addEventListener('input', () => {
    const query = input.value.trim();
    if (timeout) clearTimeout(timeout);

    timeout = setTimeout(() => fetchAndRender(query), 500);
  });

  fetchAndRender(''); // первичный рендер
}
