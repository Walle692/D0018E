<!-- src/views/ProductsView.vue -->
<template>
  <div class="page">
    <div class="header">
      <h1>Products</h1>
      <p class="subtitle">Browse our latest products</p>
    </div>

    <div v-if="loading" class="state">Loading products...</div>
    <div v-else-if="error" class="state error">{{ error }}</div>
    <div v-else-if="products.length === 0" class="state">No products found.</div>

    <div v-else class="grid">
      <article
        v-for="p in products"
        :key="p.product_id || p.product_id || p.picture_url"
        class="product-card"
        @click="router.push(`/products/${p.product_id}`)"
      >
        <div class="image-wrap">
          <img
            v-if="p.picture_url"
            :src="p.picture_url"
            :alt="p.product_name || 'Product image'"
            loading="lazy"
          />
          <div v-else class="image-fallback">No image</div>
        </div>

        <div class="info">
          <h3 class="name">{{ p.product_name || 'Unnamed product' }}</h3>
          <div class="price">{{ formatPrice(p.price) }}</div>
        </div>
      </article>
    </div>

    <button class="linkBtn" type="button" @click="router.push('/user')"><- Back to account</button>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getProducts } from '@/services/products'

const router = useRouter()

const products = ref([])
const loading = ref(false)
const error = ref('')

function formatPrice(price) {
  if (price === null || price === undefined || price === '') return 'Price unavailable'
  if (typeof price === 'number') return `${price} kr`
  return `${price}`
}

async function fetchProducts() {
  loading.value = true
  error.value = ''

  try {
    const data = await getProducts()
    products.value = Array.isArray(data) ? data : data?.products || []
  } catch (e) {
    error.value = e?.message || 'Failed to load products'
  } finally {
    loading.value = false
  }
}

onMounted(fetchProducts)
</script>

<style scoped>
.page {
  padding: 24px;
  max-width: 1000px;
  margin: 0 auto;
}

.header {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.subtitle {
  opacity: 0.8;
}

.state {
  margin-top: 16px;
}

.state.error {
  color: #b00020;
}

.grid {
  margin-top: 18px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 14px;
}

.product-card {
  border: 1px solid var(--border, #ddd);
  border-radius: 10px;
  background: var(--surface, #fff);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  min-height: 220px;
}

.image-wrap {
  width: 100%;
  height: 140px;
  background: #f4f4f4;
  display: flex;
  align-items: center;
  justify-content: center;
}

.image-wrap img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.image-fallback {
  font-size: 12px;
  opacity: 0.6;
}

.info {
  padding: 10px 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.name {
  font-size: 14px;
  line-height: 1.2;
  margin: 0;
}

.price {
  font-weight: 600;
}

.linkBtn {
  width: auto;
  margin-top: 18px;
}
</style>
