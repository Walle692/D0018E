<!-- src/views/ProductView.vue -->
<template>
  <div class="page">
    <div class="header">
      <h1>Products</h1>
      <p class="subtitle">Browse our latest products</p>
    </div>

    <div v-if="loading" class="state">Loading product...</div>
    <div v-else-if="error" class="state error">{{ error }}</div>

    <div v-else="product" class="product details"> 
        <div class="image-wrap">
          <img
            v-if="product.picture_url"
            :src="product.picture_url"
            :alt="product.product_name || 'Product image'"
            loading="lazy"
          />
          <div v-else class="image-fallback">No image</div>
        </div>
        
        <div class="Specs">
            <h3 class="name">{{ product.product_name || "Unnamed product" }}</h3>
            <div class="price">{{ formatPrice(product.price) }}</div>
            <p class="description">{{ product.description || "No description available." }}</p>
        </div>
    </div>

    <button class="linkBtn" type="button" @click="router.push('/products')">
      <- Back to products
    </button>

  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { getProductById } from "@/services/auth";

const router = useRouter();

const product = ref([]);
const loading = ref(false);
const error = ref("");

async function fetchProducts() {
  loading.value = true;
  error.value = "";

  try {
    const data = await getProductById(router.currentRoute.value.params.id);
    product.value = [data];
  } catch (e) {
    error.value = e?.message || "Failed to load product";
  } finally {
    loading.value = false;
  }
}

onMounted(fetchProducts);
</script>