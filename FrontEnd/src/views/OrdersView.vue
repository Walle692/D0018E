<template>
  <div class="page">
    <div class="card">
      <h1>My Orders</h1>

      <p v-if="loading">Loading...</p>
      <p v-else-if="error" class="error">{{ error }}</p>

      <template v-else>
        <p v-if="orders.length === 0" class="muted">No orders yet.</p>

        <div v-for="o in orders" :key="o.order_id" class="order">
          <div class="orderHeader">
            <div>
              <div class="orderTitle">Order #{{ o.order_id }}</div>
              <div class="muted">{{ formatDate(o.orderdate) }}</div>
            </div>
            <div class="total">Total: {{ money(o.totalprice) }}</div>
          </div>

          <div class="items">
            <div v-for="(it, idx) in o.order_items" :key="idx" class="item">
              <img
                v-if="it.product?.picture_url"
                :src="it.product.picture_url || placeholderImg"
                class="thumb"
                alt=""
                @error="onImgError"
              />

              <div class="info">
                <router-link :to="`/products/${it.product.product_id}`" class="name">
                  {{ it.product.product_name }}
                </router-link>

                <div class="muted">Qty: {{ it.quantity }} · Price: {{ money(it.price) }}</div>
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
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 10px;
}

h1 {
  margin: 0 0 12px;
}

.error {
  color: #b00020;
}

.muted {
  color: #666;
  font-size: 13px;
}

.order {
  margin-top: 14px;
  padding: 14px;
  border: 1px solid #e5e5e5;
  border-radius: 10px;
}

.orderHeader {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}

.orderTitle {
  font-weight: 700;
}

.total {
  font-weight: 700;
}

.items {
  margin-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.item {
  display: flex;
  gap: 12px;
  align-items: center;
}

.thumb {
  width: 64px;
  height: 64px;
  border-radius: 8px;
  border: 1px solid #eee;
  object-fit: cover;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.name {
  font-weight: 700;
  text-decoration: none;
}

.name:hover {
  text-decoration: underline;
}
</style>
