<template>
  <div class="page">
    <div class="card">
      <div class="card__header">My Orders</div>

      <p v-if="loading">Loading orders...</p>
      <p v-else-if="error">{{ error }}</p>

      <template v-else>
        <p v-if="orders.length === 0">You have no orders.</p>

        <div v-else v-for="o in orders" :key="o.order_id" class="order">
          <div class="order__header">
            <div>
              <div class="order__title">Order #{{ o.order_id }}</div>
              <div class="order__date">{{ formatDate(o.orderdate) }}</div>
            </div>
            <div class="order__total">Total: {{ money(o.totalprice) }}</div>
          </div>

          <div class="order__items">
            <div v-for="(it, idx) in o.order_items" :key="idx" class="item">
              <img
                v-if="it.picture_url"
                :src="it.picture_url || placeholderImg"
                class="item__image"
                alt=""
                @error="onImgError"
              />

              <div>
                <router-link class="item__name" :to="`/products/${it.product_id}`">
                  {{ it.product_name }}
                </router-link>

                <div>Qty: {{ it.quantity }} · Price: {{ money(it.price) }}</div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getOrders } from '@/services/auth'

const router = useRouter()
const loading = ref(true)
const error = ref('')
const orders = ref([])

onMounted(async () => {
  try {
    orders.value = await getOrders()
  } catch (e) {
    error.value = e?.message || 'Failed to load orders'
    router.push('/')
  } finally {
    loading.value = false
  }
})

import placeholderImg from '@/assets/placeholder.webp'

const onImgError = (e) => {
  e.target.src = placeholderImg
}

const formatDate = (d) => {
  const dt = new Date(d)
  return Number.isNaN(dt.getTime()) ? String(d) : dt.toLocaleString()
}

const money = (x) => {
  const n = Number(x)
  return Number.isNaN(n) ? String(x) : n.toFixed(2)
}
</script>

<style scoped>
/* Make it wider by increasing max-width or removing it entirely */
.page {
  max-width: 1100px; /* <- change this */
  margin: 24px auto;
  padding: 0 16px;
}

.card {
  display: flex;
  flex-direction: column;
  padding: 24px;
  border-radius: 24px;
  background-color: var(--surface);
  box-shadow: 3px 3px var(--shadow);
  gap: 24px;
}

.card__header {
  font-weight: bold;
  font-size: 2rem;
}

.order {
  padding: 24px;
  border-radius: 24px;
  background-color: var(--bg);
}

.order__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
}

.order__title {
  font-weight: bold;
}

.order__total {
  font-weight: bold;
}

.order__items {
  margin-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.item {
  display: flex;
  gap: 16px;
  align-items: center;
}

.item__image {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  border: 1px solid #eee;
  object-fit: cover;
  display: flex;
  align-items: center;
  justify-content: center;
}

.item__info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.item__name {
  text-decoration: none;
  color: inherit;
  font-weight: 700;
  text-decoration: none;
}

.item__name:hover {
  color: var(--hover);
}
</style>
