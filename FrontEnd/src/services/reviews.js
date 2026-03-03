const BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function getReviewsForProduct(id) {
  const res = await fetch(`${BASE_URL}/private/reviews/products/${id}`, {
    method: 'GET',
    credentials: 'include',
  })

  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'get reviews failed')
  }

  return await res.json()
}

export async function createReview({productID, comment, rating}) {
  const res = await fetch(`${BASE_URL}/private/review/write`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ product_id, comment, rating }),
  })
  if (!res.ok) {
    const text = await res.text().catch(() => '')
    throw new Error(text || 'error creating review')
  }
  return true
}