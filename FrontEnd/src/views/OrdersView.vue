<template>
  <div class="container">
    <div class="card">
      <h1>My Orders</h1>

      <p v-if="loading">Loading...</p>
      <p v-else-if="error" class="error">{{ error }}</p>

      <div v-else>
        <p v-if="orders.length === 0">No orders yet.</p>

        <div v-for="o in orders" :key="o.order_id" class="order">
          <div class="order-header">
            <div>
              <div>
                <b>Order #{{ o.order_id }}</b>
              </div>
              <div class="muted">{{ formatDate(o.orderdate) }}</div>
            </div>
            <div class="total">
              Total: <b>{{ money(o.totalprice) }}</b>
            </div>
          </div>

          <div class="items">
            <div v-for="(it, idx) in o.order_items" :key="idx" class="item">
              <img v-if="it.product?.picture_url" :src="it.product.picture_url" class="thumb" />
              <div class="item-info">
                <router-link :to="`/products/${it.product.product_id}`" class="product-link">
                  {{ it.product.product_name }}
                </router-link>
                <div class="muted">
                  Qty: <b>{{ it.quantity }}</b> · Price: <b>{{ money(it.price) }}</b>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- debug -->
        <!-- <pre>{{ orders }}</pre> -->
      </div>
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

function formatDate(d) {
  const dt = new Date(d)
  return Number.isNaN(dt.getTime()) ? String(d) : dt.toLocaleString()
}
function money(x) {
  const n = Number(x)
  return Number.isNaN(n) ? String(x) : n.toFixed(2)
}
</script>

<style scoped>
.container {
  max-width: 900px;
  margin: 30px auto;
  padding: 0 16px;
}

.card {
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 10px;
}

.error {
  color: #b00020;
  margin-top: 10px;
}

.order {
  border: 1px solid #e5e5e5;
  border-radius: 10px;
  padding: 14px;
  margin-top: 14px;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
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
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: 8px;
  border: 1px solid #eee;
}

.item-info {
  display: flex;
  flex-direction: column;
}

.product-link {
  font-weight: 700;
  text-decoration: none;
}

.product-link:hover {
  text-decoration: underline;
}

.muted {
  color: #666;
  font-size: 13px;
  margin-top: 2px;
}

.total {
  font-size: 16px;
}
</style>
