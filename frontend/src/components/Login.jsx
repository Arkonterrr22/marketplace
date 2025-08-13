import React, { useState } from 'react';
import api from '../lib/api';
import { saveToken } from '../lib/session';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  async function handleLogin(e) {
    e.preventDefault();
    try {
      const res = await api.post('/login', { email, password });
      const token = res.data.token;
      if (token) {
        saveToken(token);
        window.location.href = '/';
      } else {
        alert('Токен не получен');
      }
    } catch (err) {
      alert('Ошибка входа: ' + (err.response?.data || err.message));
    }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '30%', margin: '0 auto', gap: '1rem', padding: '1vh', boxSizing: 'border-box' }}>
      <form onSubmit={handleLogin} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '100%', gap: '1rem' }}>
        <h2 style={{ margin: 0 }}>Вход</h2>
        <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} required style={{ width: '100%', padding: '0.5rem', boxSizing: 'border-box' }} />
        <input type="password" placeholder="Пароль" value={password} onChange={e => setPassword(e.target.value)} required minLength={6} style={{ width: '100%', padding: '0.5rem', boxSizing: 'border-box' }} />
        <button
          type="submit"
          style={{ width: '100%', padding: '0.5rem', cursor: 'pointer', backgroundColor: '#222', color: '#fff', border: 'none', borderRadius: '4px', transition: 'background-color 0.3s' }}
          onMouseEnter={e => (e.currentTarget.style.backgroundColor = '#444')}
          onMouseLeave={e => (e.currentTarget.style.backgroundColor = '#222')}
        >
          Войти
        </button>
      </form>

      <button
        type="button"
        onClick={() => (window.location.href = '/register')}
        style={{ width: '100%', padding: '0.5rem', cursor: 'pointer', backgroundColor: '#555', color: '#fff', border: 'none', borderRadius: '4px', transition: 'background-color 0.3s' }}
        onMouseEnter={e => (e.currentTarget.style.backgroundColor = '#777')}
        onMouseLeave={e => (e.currentTarget.style.backgroundColor = '#555')}
      >
        Регистрация
      </button>
    </div>
  );
}
