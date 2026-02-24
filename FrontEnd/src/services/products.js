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

export async function createProduct(product) {
  const res = await fetch(`${BASE_URL}/seller/create-product`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(product),
  })
  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'error creating product')
  }
  return true
}
