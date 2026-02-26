const BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function getOrders({ limit = 10, offset = 0 } = {}) {
  const res = await fetch(`${BASE_URL}/private/orders?limit=${limit}&offset=${offset}`, {
    method: 'GET',
    credentials: 'include',
  })

  const text = await res.text().catch(() => '')
  if (!res.ok) throw new Error(text || 'get orders failed')
  return text ? JSON.parse(text) : []
}
