<template>
  <div class="page">
    <div class="header">
      <h1>Your basket</h1>
      <p class="subtitle">Items in your shopping cart</p>
    </div>

    <div v-if="loading" class="state">Loading basket...</div>
    <div v-else-if="error" class="state error">{{ error }}</div>

    <div v-else>
      <div v-if="items.length === 0" class="state">Your basket is empty.</div>
      <div v-else class="basket-list">
        <div v-for="(item, idx) in items" :key="itemKey(item, idx)" class="basket-item">
          <div class="image-wrap">
            <img
              v-if="itemImage(item)"
              :src="itemImage(item)"
              :alt="itemName(item)"
              loading="lazy"
            />
            <div v-else class="image-fallback">No image</div>
          </div>

          <div class="details">
            <div class="name">{{ itemName(item) }}</div>
            <div class="meta">{{ itemManufacturer(item) }}</div>
            <div class="qty">Qty: {{ item.quantity ?? 0 }}</div>
          </div>

          <div class="price">
            <div class="unit">Unit: {{ formatPrice(itemPrice(item)) }}</div>
            <div class="total">Total: {{ formatPrice(itemTotal(item)) }}</div>
          </div>
        </div>

        <div class="basket-summary">
          <div class="summary-label">Subtotal</div>
          <div class="summary-value">{{ formatPrice(subtotal) }}</div>
        </div>
      </div>
    </div>

    <button class="linkBtn" type="button" @click="router.push('/products')">
      <- Back to products
    </button>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getBasket } from '@/services/auth'

const router = useRouter()

const items = ref([])
const loading = ref(true)
const error = ref('')

function normalizeItems(payload) {
  if (Array.isArray(payload)) return payload
  if (Array.isArray(payload?.items)) return payload.items
  if (Array.isArray(payload?.basket_items)) return payload.basket_items
  if (Array.isArray(payload?.data)) return payload.data
  return []
}

function itemName(item) {
  return item?.product?.product_name || item?.product_name || 'Unnamed product'
}

function itemManufacturer(item) {
  return item?.product?.manufacturer || item?.manufacturer || ''
}

function itemImage(item) {
  return item?.product?.picture_url || item?.picture_url || ''
}

function itemPrice(item) {
  const price = item?.price ?? item?.product?.price ?? 0
  const num = Number(price)
  return Number.isFinite(num) ? num : 0
}

function itemTotal(item) {
  const qty = Number(item?.quantity ?? 0)
  return itemPrice(item) * (Number.isFinite(qty) ? qty : 0)
}

function itemKey(item, idx) {
  return item?.basket_item_id ?? item?.product?.product_id ?? item?.product_id ?? idx
}

const subtotal = computed(() => items.value.reduce((sum, item) => sum + itemTotal(item), 0))

function formatPrice(value) {
  const num = Number(value)
  return Number.isFinite(num) ? num.toFixed(2) : '0.00'
}

async function fetchBasket() {
  loading.value = true
  error.value = ''
  items.value = []

  try {
    const data = await getBasket()
    items.value = normalizeItems(data)
  } catch (e) {
    error.value = e?.message || 'Failed to load basket'
  } finally {
    loading.value = false
  }
}

onMounted(fetchBasket)
</script>
