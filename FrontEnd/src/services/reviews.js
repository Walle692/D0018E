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