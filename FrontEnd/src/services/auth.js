const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function login({ username, password }) {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    // NO credentials yet
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    // Try to read a useful error message
    const text = await res.text().catch(() => "");
    throw new Error(text || "Login failed");
  }

  return true;
}
