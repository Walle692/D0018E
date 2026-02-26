<template>
  <div class="page">
    <div class="card">
      <div class="card__header">Search by all orders</div>

      <SearchBar
        v-model="searchText"
        v-model:typeValue="searchType"
        placeholder="Search orders..."
        :disabled="loading"
        @search="runSearch(true)"
      />

      <p v-if="loading">Loading orders...</p>
      <p v-else-if="error">{{ error }}</p>

      <template v-else>
        <p v-if="hasSearched && orders.length === 0">No orders found.</p>

        <div v-else v-for="o in orders" :key="o.order_id" class="order">
          <div class="order__header">
            <div>
              <div class="order__title">Order #{{ o.order_id }}</div>
              <div class="order__date">{{ formatDate(o.orderdate) }}</div>
              <div class="order__sub">User ID: {{ o.order_user_id }}</div>
            </div>
            <div class="order__total">Total: {{ money(o.totalprice) }}</div>
          </div>

          <div class="order__items">
            <ItemContainer
              v-for="(it, idx) in o.order_items"
              :key="`${o.order_id}-${idx}`"
              :image-src="it.picture_url"
              :alt="it.product_name"
              :height="128"
              :placeholder="placeholderImg"
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

    <div class="pager">
      <button type="button" :disabled="offset === 0 || loading" @click="prevPage">Prev</button>
      <button type="button" :disabled="orders.length < limit || loading" @click="nextPage">
        Next
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ItemContainer from '@/components/ItemContainer.vue'
import SearchBar from '@/components/SearchBar.vue'
import placeholderImg from '@/assets/placeholder.webp'
import { searchOrders } from '@/services/orders'

const loading = ref(false)
const error = ref('')
const orders = ref([])

const limit = 10
const offset = ref(0)

const searchType = ref('username') // username | orderID | userID
const searchText = ref('')

const hasSearched = ref(false)

async function runSearch(resetOffset = false) {
  if (!searchText.value.trim()) {
    error.value = 'Please enter a search value.'
    orders.value = []
    return
  }

  hasSearched.value = true
  if (resetOffset) offset.value = 0

  loading.value = true
  error.value = ''

  try {
    orders.value = await searchOrders({
      search_type: searchType.value,
      search_object: searchText.value,
      limit,
      offset: offset.value,
    })
  } catch (e) {
    error.value = e?.message || 'Failed to load orders'
    orders.value = []
  } finally {
    loading.value = false
  }
}

function nextPage() {
  offset.value += limit
  runSearch(false)
}

function prevPage() {
  offset.value = Math.max(0, offset.value - limit)
  runSearch(false)
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
.page {
  max-width: 1100px;
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
  gap: 16px;
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

.order__sub {
  opacity: 0.7;
  font-size: 0.95rem;
}

.order__total {
  font-weight: bold;
}

.order__items {
  margin-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;
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

.pager {
  margin-top: 12px;
  display: flex;
  gap: 12px;
}
</style>
