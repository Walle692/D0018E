const BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function addToBasket(productId, quantity) {
  const res = await fetch(`${BASE_URL}/private/basket/add`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ product_id: productId, quantity }),
  })

  return res.ok
}

export async function getBasket() {
  const res = await fetch(`${BASE_URL}/private/basket`, {
    method: 'GET',
    credentials: 'include',
  })

  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'get basket failed')
  }

  return await res.json()
}

export async function checkoutBasket() {
  const res = await fetch(`${BASE_URL}/private/checkout`, {
    method: 'POST',
    credentials: 'include',
  })

  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'checkout failed')
  }

  return await res.json().catch(() => ({}))
}
