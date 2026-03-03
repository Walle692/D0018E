<template>
  <div class="page">
    <div class="layout">
      <div class="search">
        <div class="search__header">Product Search</div>
        <SearchBar
          v-model="searchText"
          v-model:typeValue="searchType"
          :types="[
            { value: 'product_name', label: 'Product name' },
            { value: 'manufacturer', label: 'Manufacturer' },
          ]"
          placeholder="Search..."
          :disabled="loading"
          @search="runSearch(true)"
        />
        <div class="options-row">
          <select v-model="sortType" class="opt" :disabled="loading">
            <option value="price">Sort: Price</option>
            <option value="screen_size">Sort: Screen size</option>
          </select>
          <select v-model="sortDirection" class="opt" :disabled="loading">
            <option value="">Ascending</option>
            <option value="descending">Descending</option>
          </select>
          <button
            type="button"
            class="sb_btn sb_btn--reset"
            :disabled="loading"
            @click="resetAndSearch"
          >
            Reset
          </button>
        </div>
      </div>
      <p v-if="loading">Searching...</p>
      <p v-else-if="error">{{ error }}</p>
      <template v-else>
        <p v-if="hasSearched && products.length === 0">No products found.</p>
        <p v-else-if="!hasSearched">Enter a search and press Search.</p>
        <div v-else class="products_area">
          <ProductCard v-for="p in products" :key="p.product_id" :product="p" :imgHeight="256">
          </ProductCard>
        </div>
      </template>
    </div>
    <div class="pager">
      <button type="button" :disabled="offset === 0 || loading || !hasSearched" @click="prevPage">
        Prev
      </button>
      <button
        type="button"
        :disabled="loading || !hasSearched || products.length < limit"
        @click="nextPage"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import SearchBar from '@/components/SearchBar.vue'
import { searchProducts } from '@/services/products'
import ProductCard from '@/components/ProductCard.vue'

const loading = ref(false)
const error = ref('')
const hasSearched = ref(false)

const products = ref([])

const limit = 10
const offset = ref(0)

const searchType = ref('product_name')
const searchText = ref('')

const sortType = ref('price')
const sortDirection = ref('') // '' = ascending, 'descending' = descending

async function runSearch(resetOffset = false) {
  hasSearched.value = true
  if (resetOffset) offset.value = 0

  loading.value = true
  error.value = ''

  try {
    products.value = await searchProducts({
      search_type: searchType.value,
      search_object: searchText.value,
      sort_type: sortType.value,
      sort_direction: sortDirection.value,
      limit,
      offset: offset.value,
    })
  } catch (e) {
    error.value = e?.message || 'Search failed'
    products.value = []
  } finally {
    loading.value = false
  }
}

function resetAndSearch() {
  searchText.value = ''
  searchType.value = ''
  sortType.value = 'price'
  sortDirection.value = ''
  offset.value = 0
  runSearch(true)
}

function nextPage() {
  offset.value += limit
  runSearch(false)
}

function prevPage() {
  offset.value = Math.max(0, offset.value - limit)
  runSearch(false)
}

onMounted(resetAndSearch)
</script>

<style scoped>
.layout {
  display: grid;
  grid-template-areas:
    '. search .'
    'product product product';
  grid-template-columns: 1fr 3fr 1fr;
  gap: 36px;
  justify-content: center;
}
.search {
  grid-area: search;
  width: 100%;
  display: flex;
  flex-direction: column;
  padding: 24px;
  border-radius: 24px;
  background-color: var(--surface);
  box-shadow: 3px 3px var(--shadow);
  gap: 16px;
}

.search__header {
  font-weight: bold;
  font-size: 2rem;
}

.options-row {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* Override global button because some buffon thought that was a good idea*/
.sb_btn {
  width: auto !important;
  margin-top: 0 !important;
  display: inline-flex !important;
  align-items: center;
  justify-content: center;
  padding: 0 14px !important;
  height: 42px;
  border-radius: 12px;
  border: 2px solid var(--shadow);
  background: var(--bg);
  color: inherit;
  font-size: 1rem;
  box-sizing: border-box;
  cursor: pointer;
  white-space: nowrap;
}

.opt {
  height: 42px;
  padding: 0 12px;
  border-radius: 12px;
  border: 2px solid var(--shadow);
  background: var(--bg);
  font-size: 1rem;
  box-sizing: border-box;
}

.products_area {
  grid-area: product;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  justify-content: center;
  align-items: center;
  gap: 48px;
  background: var(--surface);
  padding: 48px;
  border-radius: 48px;
  box-shadow: 6px 6px var(--shadow);
}

.pager {
  margin-top: 8px;
  display: flex;
  gap: 12px;
}
@media (max-width: 1800px) {
  .products_area {
    grid-template-columns: repeat(3, 1fr);
  }
}
@media (max-width: 1500px) {
  .products_area {
    grid-template-columns: repeat(3, 1fr);
  }
}
@media (max-width: 1200px) {
  .products_area {
    grid-template-columns: repeat(2, 1fr);
  }
}
@media (max-width: 900px) {
  .layout {
    grid-template-columns: 1fr;
  }
  .products_area {
    grid-template-columns: repeat(1, 1fr);
  }
}
</style>
