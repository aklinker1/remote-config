const storage = window.sessionStorage;
const STORAGE_KEY = "token";

export function provideAuthToken() {
  const token = ref(getAuthToken());
  const setToken = (newValue: string | undefined) => {
    token.value = newValue;
    setAuthToken(newValue);
  };
  window.addEventListener("storage", () => {
    const newValue = getAuthToken();
    if (newValue !== token.value) setToken(newValue);
  });
  provide(STORAGE_KEY, token);
  provide(STORAGE_KEY + "-setter", setToken);
}

export function getAuthToken(): string | undefined {
  return storage.getItem(STORAGE_KEY) ?? undefined;
}

export function setAuthToken(token: string | undefined) {
  if (token) storage.setItem(STORAGE_KEY, token);
  else storage.removeItem(STORAGE_KEY);
}

export function useAuthToken() {
  return inject<string | undefined>(STORAGE_KEY, undefined);
}
