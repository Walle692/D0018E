const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function login({ username, password }) {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    // Try to read a useful error message
    const text = await res.text().catch(() => "");
    throw new Error(text || "Login failed");
  }

  return true;
}

export async function getMe() {
  const res = await fetch(`${BASE_URL}/private/me`,
    {
      method: "GET",
      credentials: "include"
    }
  );

  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(text || "Get me failed");
  }

  return await res.json();
}
