import { useState } from 'react';
async function addItem({ username, params = {} } = {}) {
        const res = await fetch('/api/addItem', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, params }),
        });
        const data = await res.json() || { results: [], warn: 'no data!' };
        console.log(data);
        return data;
        }

export default function AddItemWindow() {
  const [open, setOpen] = useState(false);
  const params = {}
  return (
    <>
      <button onClick={() => setOpen(true)}>Добавить объявление</button>
      {open && (
        <div className="modal" onClick={() => setOpen(false)}>
          <div className="modal-content" onClick={e => e.stopPropagation()}>
            <span className="close" onClick={() => setOpen(false)}>&times;</span>
            <h2>Модальное окно</h2>
            <button id="add-item" onClick={() => addItem({ username: window.currentUser.email, params })}>добавить товар</button>
          </div>
        </div>
      )}

      <style jsx>{`
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(26, 26, 26, 0.8); /* затемнение фона */
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
    }
    .modal-content {
      flex: 1;
    width: 100%;
    max-width: 100%;         /* ширина 80% родителя */
    margin-inline: 12%;             /* чтобы flex работал корректно */
    display: flex;
    gap: 2rem;
    flex-direction: column;
    align-items: center;
    overflow-y: auto;
    background: #1a1a1a;
    margin-top: 2rem;
    box-sizing: border-box;
    }
  .close {
    position: absolute;
    top: 0.5rem;
    right: 1rem;
    font-size: 1.5rem;
    cursor: pointer;
  }
`}</style>
    </>
  );
}
