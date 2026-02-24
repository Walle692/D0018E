<template>
  <div class="page">
    <div class="layout">
      <div class="product-form">
        <div class="form__header">Create a new product</div>

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

          <button class="submit-btn" type="button" :disabled="loading" @click="submit">
            {{ loading ? 'Creating...' : 'Create product' }}
          </button>
        </div>
      </div>
      <div class="products">
        <div class="products__header">Your products</div>
        <!-- Place holder for sellers products-->
        <div class="item">
          <img class="item__image" src="https://placehold.co/400x400" alt="image" />
          <div class="item__info">
            <div class="item__name">Placeholder</div>
            <div class="item__price">Price : 399</div>
            <div class="item__stock">Stock : 200</div>
            <div class="item__available">Available: No/Yes</div>
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
import { ref } from 'vue'
import { createProduct } from '@/services/products'

const loading = ref(false)
const error = ref('')
const success = ref(false)

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

async function submit() {
  loading.value = true
  error.value = ''
  success.value = false

  try {
    await createProduct(form.value)
    success.value = true
    form.value = getDefaultForm()
  } catch (e) {
    error.value = e?.message || 'Failed to create product'
  } finally {
    loading.value = false
  }
}
</script>
