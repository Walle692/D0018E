<template>
  <div class="page">
    <div>
      <h1>Your basket</h1>
      <p>Items in your shopping cart</p>

      <p v-if="loading">Loading basket...</p>
      <p v-else-if="error">{{ error }}</p>

      <template v-else>
        <p v-if="items.length === 0">Your basket is empty.</p>

        <div v-else class="layout">
          <div class="layout__basket">
            <div v-for="it in items" :key="it.product_id" class="item_container">
              <img
                class="item_container__img"
                :src="it.picture_url || placeholder"
                alt=""
                @error="onImgError"
              />
              <div class="item_container__data">
                <router-link class="item_container__name" :to="`/products/${it.product_id}`">
                  {{ it.product_name }}
                </router-link>
                <div class="item_container__manufacturer">{{ it.manufacturer }}</div>
                <div>
                  Qty: <b>{{ it.quantity }}</b>
                  <span :class="it.available ? 'ok' : 'bad'">
                    {{ it.available ? 'Available' : 'Unavailable' }}
                  </span>
                </div>
                <div>
                  Unit: <b>{{ money(it.price) }}</b>
                </div>
                <div>
                  Total: <b>{{ money(it.price * it.quantity) }}</b>
                </div>
              </div>
            </div>
          </div>

          <div>
            <div>Subtotal</div>
            <div>
              <b>{{ money(totalprice) }}</b>
            </div>
            <div>
              <button @click="router.push('/products')">← Back</button>
              <button @click="doCheckout" :disabled="items.length === 0">Checkout</button>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { getBasket, checkoutBasket } from '@/services/basket'

const router = useRouter()

const loading = ref(true)
const error = ref('')

const basket = ref({ totalprice: 0, basket_items: [] })

const items = computed(() =>
  Array.isArray(basket.value?.basket_items) ? basket.value.basket_items : [],
)

const totalprice = computed(() => Number(basket.value?.totalprice ?? 0))

import placeholderImg from '@/assets/placeholder.webp'

const onImgError = (e) => {
  e.target.src = placeholderImg
}

function money(x) {
  const n = Number(x)
  return Number.isNaN(n) ? '-' : n.toFixed(2)
}

async function fetchBasket() {
  loading.value = true
  error.value = ''
  try {
    const data = await getBasket()

    basket.value = data && typeof data === 'object' ? data : { totalprice: 0, basket_items: [] }
  } catch (e) {
    error.value = e?.message || 'Failed to load basket'
    basket.value = { totalprice: 0, basket_items: [] }
  } finally {
    loading.value = false
  }
}

async function doCheckout() {
  error.value = ''
  try {
    await checkoutBasket() // POST /private/checkout
    await fetchBasket() // refresh basket (should be empty)
    router.push('/orders') // go to orders view
  } catch (e) {
    error.value = e?.message || 'Checkout failed'
  }
}

onMounted(fetchBasket)
</script>

<style>
.layout {
  display: grid;
  grid-template-areas: 'basket checkout';
  grid-template-columns: 2fr 1fr;
  gap: 12px;
}

.layout__basket {
  grid-area: basket;
  display: grid;
  grid-template-columns: 1fr;
  gap: 5px;
  background-color: white;
  padding: 40px;
  border-radius: 8px;
}

.layout__checkout {
  grid-area: checkout;
}

.item_container {
  display: grid;
  grid-template-areas: 'img data';
  grid-template-columns: 128px 1fr;
  gap: 20px;
  background-color: white;
  border-radius: 8px;
  border: 1px solid grey;
  padding: 20px;
}

.item_container__img {
  grid-area: img;
  width: 128px;
  height: 128px;
  object-fit: cover;
  border-radius: 8px;
}

.item_container__data {
  grid-area: data;
  display: flex;
  flex-direction: column;
}

.item_container__name {
  font-weight: bold;
  font-size: 1.4rem;
  text-decoration: none;
  color: inherit;
}
</style>
