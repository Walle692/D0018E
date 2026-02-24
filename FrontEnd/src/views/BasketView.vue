<template>
  <div class="page">
    <div>
      <p v-if="loading">Loading basket...</p>
      <p v-else-if="error">{{ error }}</p>

      <template v-else>
        <p v-if="items.length === 0">Your basket is empty.</p>

        <div v-else class="layout">
          <div class="layout__basket">
            <div class="basket__header">Your basket</div>
            <div v-for="it in items" :key="it.product_id" class="item_container">
              <img
                class="item_container__img"
                :src="it.picture_url || placeholder"
                alt=""
                @error="onImgError"
              />

              <div>
                <router-link class="item_container__name" :to="`/products/${it.product_id}`">
                  {{ it.product_name }}
                </router-link>
                <div class="item_container__manufacturer">{{ it.manufacturer }}</div>
                <div>
                  <span :class="it.available ? 'ok' : 'bad'">
                    {{ it.available ? 'Available' : 'Unavailable' }}
                  </span>
                </div>
              </div>

              <div>
                <div>
                  Qty: <b>{{ it.quantity }}</b>
                </div>
                <div>
                  Unit: <b>{{ money(it.price) }}</b>
                </div>
                <div>
                  Total: <b>{{ money(it.price * it.quantity) }}</b>
                </div>
              </div>

              <div class="basket__delete">
                <img
                  src="@/assets/trash.svg"
                  alt="delete"
                  class="delete-icon"
                  @click="handleDelete(it.product_id)"
                />
              </div>
            </div>
          </div>

          <div class="checkout">
            <div class="checkout__header">Subtotal</div>
            <div class="checkout__total">${{ money(totalprice) }}</div>
            <div>
              <button @click="doCheckout" :disabled="items.length === 0">Checkout</button>
              <button @click="router.push('/products')">Continue shopping</button>
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

<style scoped>
.layout {
  display: grid;
  grid-template-areas:
    'basket checkout'
    'basket .';
  grid-template-columns: 2fr 1fr;
  gap: 12px;
}

.layout__basket {
  grid-area: basket;
  display: flex;
  flex-direction: column;
  background: var(--surface);
  padding: 24px;
  gap: 16px;
  box-shadow: 3px 3px var(--shadow);
  border-radius: 24px;
}

.basket__header {
  font-weight: bold;
  font-size: 2.3rem;
}
.item_container {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr 2fr;
  padding: 24px;
  gap: 20px;
  align-items: center;
  border-radius: 24px;
  background-color: var(--shadow);
}

.item_container__img {
  width: 128px;
  height: 128px;
  object-fit: cover;
  border-radius: 8px;
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

.checkout {
  grid-area: checkout;
  background-color: var(--surface);
  padding: 24px;
  border-radius: 24px;
  box-shadow: 3px 3px var(--shadow);
  text-align: center;
}

.checkout__header {
  font-weight: 600;
  font-size: 2.5rem;
}
.checkout__total {
  font-size: 2rem;
}

.basket__delete {
  display: flex;
  justify-content: center;
  gap: 12px;
  background-color: var(--surface);
  padding: 8px 16px;
  width: 20px;
  border-radius: 999px;
}

.delete-icon {
  width: 32px;
  height: 32px;
  cursor: pointer;
  opacity: 0.6;
}

.delete-icon:hover {
  opacity: 1;
}

@media (max-width: 900px) {
  .layout {
    grid-template-areas:
      'checkout'
      'basket';
    grid-template-columns: 1fr;
  }

  .item_container__img {
    width: 64px;
    height: 64px;
    object-fit: cover;
    border-radius: 8px;
  }

  .layout__basket {
    padding: 8px;
  }
  .item_container {
    padding: 8px;
  }
  .basket__header {
    text-align: center;
  }
}
</style>
