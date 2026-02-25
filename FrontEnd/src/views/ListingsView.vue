<template>
  <div class="page">
    <div class="layout">
      <div class="product-form">
        <div class="form__header">
          {{ formMode === 'create' ? 'Create a new product' : 'Edit product' }}
        </div>

        <div class="pair">
          <div class="form__group">
            <label class="form__label">Product Name</label>
            <input
              class="product_input"
              type="text"
              v-model="form.product_name"
              placeholder='27" Monitor'
            />
          </div>
          <div class="form__group">
            <label class="form__label">Manufacturer</label>
            <input
              class="product_input"
              type="text"
              v-model="form.manufacturer"
              placeholder="Monitor Builder"
            />
          </div>
        </div>
        <div class="pair">
          <div class="form__group">
            <label class="form__label">Price</label>
            <input
              class="product_input"
              type="number"
              step="0.01"
              v-model="form.price"
              placeholder="399"
            />
          </div>
          <div class="form__group">
            <label class="form__label">Stock</label>
            <input class="product_input" type="number" v-model="form.stock" placeholder="200" />
          </div>
        </div>
        <div class="pair">
          <div class="form__group">
            <label class="form__label">Screen size</label>
            <input
              class="product_input"
              type="number"
              v-model="form.screen_size"
              placeholder="27"
            />
          </div>
          <div class="form__group">
            <label class="form__label">Product image URL</label>
            <input
              class="product_input"
              type="text"
              v-model="form.picture_url"
              placeholder="https://example.com/monitor1.jpg"
            />
          </div>
        </div>
        <div class="form__group">
          <label class="form__label">Product description</label>
          <textarea
            class="product_desc"
            v-model="form.description"
            placeholder="4k Monitor 144hz 1920x1080"
          />
          <button
            v-if="formMode === 'create'"
            class="submit-btn"
            type="button"
            :disabled="loading"
            @click="submitCreate"
          >
            {{ loading ? 'Creating...' : 'Create product' }}
          </button>

          <button v-else class="submit-btn" type="button" :disabled="loading" @click="submitEdit">
            {{ loading ? 'Saving...' : 'Save changes' }}
          </button>

          <button v-if="formMode === 'edit'" type="button" :disabled="loading" @click="cancelEdit">
            Go back
          </button>
        </div>
      </div>
      <div class="products">
        <div class="products__header">Your products</div>
        <div v-if="productsLoading">Loading products...</div>
        <div v-else-if="productsError">{{ productsError }}</div>
        <div v-else-if="products.length === 0">No products yet.</div>

        <div class="item" v-else v-for="product in products" :key="product.product_id">
          <img class="item__image" :src="product.picture_url" alt="image" />
          <div class="item__info">
            <div class="item__name">{{ product.product_name }}</div>
            <div class="item__price">Price: {{ product.price }}</div>
            <div class="item__stock">Stock: {{ product.stock }}</div>
            <div class="item__available">Available: {{ product.active ? 'Yes' : 'No' }}</div>
          </div>

          <div class="btnwrap"><button type="button" @click="openEdit(product)">Edit</button></div>
          <div class="btnwrap">
            <button type="button" @click="submitDelist(product.product_id)">Delist</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.item {
  display: flex;
  background: var(--bg);
  max-height: 96px;
  gap: 24px;
  padding: 16px;
  border-radius: 24px;
}

.item__image {
  object-fit: cover;
  width: 96px;
  border-radius: 24px;
}

.item__name {
  font-weight: bold;
  font-size: 1.2rem;
}

.product-form,
.products {
  display: flex;
  flex-direction: column;
  align-self: start;
  background: var(--surface);
  padding: 24px;
  border-radius: 24px;
  box-shadow: 5px 5px var(--shadow);
  gap: 24px;
}

.form__header,
.products__header {
  font-weight: bold;
  font-size: 2rem;
}

.form__group {
  display: flex;
  flex-direction: column;
}

.form__label {
  padding-left: 8px;
  font-weight: 600;
  font-size: 1.1rem;
}

.pair {
  display: flex;
  gap: 24px;
}

.form__group {
  flex: 1;
}

.product_input,
.product_desc {
  padding: 8px 16px;
  border: 2px solid var(--shadow);
  border-radius: 12px;
  font-size: 1rem;
  background: var(--bg);
}

.product_desc {
  height: 10vh;
}

.product_input:focus {
  outline: solid 2px var(--hover);
}

@media (max-width: 900px) {
  .pair {
    flex-direction: column;
    gap: 16px;
  }
}
</style>

<script setup>
import { onMounted, ref } from 'vue'
import { createProduct, delistProduct, getSellerProducts } from '@/services/products'

const formMode = ref('create') // 'create' | 'edit'
const editingProductId = ref(null)

const loading = ref(false)
const error = ref('')
const success = ref(false)

const productsLoading = ref(false)
const productsError = ref('')
const products = ref([])

const getDefaultForm = () => ({
  product_name: '',
  price: null,
  description: '',
  stock: null,
  screen_size: null,
  manufacturer: '',
  picture_url: '',
})
const form = ref(getDefaultForm())

//Manage state
function resetFormToCreateMode() {
  formMode.value = 'create'
  editingProductId.value = null
  form.value = getDefaultForm()
}

function openEdit(product) {
  formMode.value = 'edit'
  editingProductId.value = product.product_id
  error.value = ''
  success.value = false

  form.value = {
    product_name: product.product_name ?? '',
    price: product.price ?? null,
    description: product.description ?? '',
    stock: product.stock ?? null,
    screen_size: product.screen_size ?? null,
    manufacturer: product.manufacturer ?? '',
    picture_url: product.picture_url ?? '',
  }
}

function cancelEdit() {
  error.value = ''
  success.value = false
  resetFormToCreateMode()
}

// API calls
async function submitCreate() {
  loading.value = true
  error.value = ''
  success.value = false

  try {
    await createProduct(form.value)
    success.value = true
    form.value = getDefaultForm()
    await loadSellerProducts()
  } catch (e) {
    error.value = e?.message || 'Failed to create product'
  } finally {
    loading.value = false
  }
}

async function submitDelist(productID) {
  loading.value = true
  error.value = ''
  success.value = false

  try {
    await delistProduct(productID)
    success.value = true
    await loadSellerProducts()
  } catch (e) {
    error.value = e?.message || 'Failed to create product'
  } finally {
    loading.value = false
  }
}

async function submitEdit() {
  return
}

async function loadSellerProducts() {
  productsLoading.value = true
  productsError.value = ''

  try {
    const data = await getSellerProducts()

    // If your backend returns an array directly:
    products.value = Array.isArray(data) ? data : []

    // If backend returns { products: [...] } instead, use:
    //products.value = Array.isArray(data.products) ? data.products : []
  } catch (e) {
    productsError.value = e?.message || 'Failed to load products'
    products.value = []
  } finally {
    productsLoading.value = false
  }
}

onMounted(() => {
  loadSellerProducts()
})
</script>
