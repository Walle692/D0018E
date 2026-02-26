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
            <ItemContainer
              v-for="(it, idx) in o.order_items"
              :key="idk"
              :image-src="it.picture_url"
              :alt="it.product_name"
              :height="128"
              :placeholder="placeholder"
            >
              <template #left-1>
                <router-link class="item_container__name" :to="`/products/${it.product_id}`">
                  {{ it.product_name }}
                </router-link>
                <div class="item_container__manufacturer">{{ it.manufacturer }}</div>
              </template>
              <template #left-2>
                <div>
                  Unit: <b>{{ money(it.price) }}</b>
                </div>
                <div>
                  Total: <b>{{ money(it.price * it.quantity) }}</b>
                </div>
                <div>
                  Qty: <b>{{ it.quantity }}</b>
                </div>
              </template>
            </ItemContainer>
          </div>
        </div>
      </template>
    </div>
    <button type="button" :disabled="offset === 0 || loading" @click="prevPage">Prev</button>
    <button type="button" :disabled="orders.length < limit || loading" @click="nextPage">
      Next
    </button>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getOrders } from '@/services/orders'
import placeholderImg from '@/assets/placeholder.webp'
import ItemContainer from '@/components/ItemContainer.vue'

const router = useRouter()
const loading = ref(true)
const error = ref('')
const orders = ref([])

const limit = 10
const offset = ref(0)

async function loadOrders() {
  loading.value = true
  error.value = ''
  try {
    orders.value = await getOrders({ limit, offset: offset.value })
  } catch (e) {
    error.value = e?.message || 'Failed to load orders'
  } finally {
    loading.value = false
  }
}

function nextPage() {
  offset.value += limit
  loadOrders()
}

function prevPage() {
  offset.value = Math.max(0, offset.value - limit)
  loadOrders()
}

onMounted(loadOrders)

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
.item_container__manufacturer {
  font-weight: 500;
  font-size: 1.1rem;
}
.item_container__name {
  font-weight: bold;
  font-size: 1.4rem;
  text-decoration: none;
  color: inherit;
}

.item_container__name:hover {
  color: var(--hover);
}

.order__total {
  font-weight: bold;
}

.order__items {
  margin-top: 12px;
  display: grid;
  grid-template-columns: auto 0.3fr auto 1fr auto auto;
  row-gap: 12px;
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
