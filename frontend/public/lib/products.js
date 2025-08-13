export async function refreshProducts({ name = '', page = 1, amount = 20 } = {}) {
  const res = await fetch('/api/search', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, page, amount }),
  });

  const data = await res.json();
  const products = data.results || [];

  const list = document.getElementById('product-list');
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
}
