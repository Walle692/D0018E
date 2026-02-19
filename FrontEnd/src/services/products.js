const BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function getProducts() {
  const res = await fetch(`${BASE_URL}/private/products`, {
    method: 'GET',
    credentials: 'include',
  })

  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'get products failed')
  }

  return await res.json()
}

export async function getProductById(id) {
  const res = await fetch(`${BASE_URL}/private/products/${id}`, {
    method: 'GET',
    credentials: 'include',
  })

  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'get product failed')
  }

  return await res.json()
}
