import { jwtDecode } from "https://cdn.jsdelivr.net/npm/jwt-decode@4.0.0/+esm";
const TOKEN_KEY = 'arkontoken';

export function saveToken(token) {
  localStorage.setItem(TOKEN_KEY, token);
}

export function getToken() {
  return localStorage.getItem(TOKEN_KEY);
}

export function removeToken() {
  localStorage.removeItem(TOKEN_KEY);
}

export function getUserFromToken() {
  const token = getToken();
  if (!token) return null;

  try {
    const decoded = jwtDecode(token);

    const now = Date.now() / 1000; // в секундах
    if (!decoded.exp || decoded.exp < now) {
      removeToken();
      return null;
    }

    return decoded;
  } catch (err) {
    console.error('Ошибка при декодировании токена:', err);
    removeToken();
    return null;
  }
}

export function Session() {
  if (typeof window === 'undefined') return null; // на сервере не выполняем

  const user = getUserFromToken();

  if (!user) {
    console.log('Гость');
    return null;
  } else {
    console.log('Пользователь:', user.email);
    return user;
  }
}