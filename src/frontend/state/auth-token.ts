const storage = window.localStorage;
const STORAGE_KEY = 'token';

export function getAuthToken(): string | undefined {
  return storage.getItem(STORAGE_KEY) ?? undefined;
}

export function setAuthToken(token: string | undefined) {
  if (token) storage.setItem(STORAGE_KEY, token);
  else storage.removeItem(STORAGE_KEY);
}
