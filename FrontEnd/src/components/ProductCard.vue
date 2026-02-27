<template>
  <router-link
    class="products"
    :to="`/products/${product.product_id}`"
    :style="{ '--img-h': toCss(imgHeight) }"
  >
    <img class="product__img" :src="product.picture_url" :alt="alt" @error="onError" />

    <div class="product__name">{{ product.product_name }}</div>
    <div class="product__manufacturer">{{ product.manufacturer }}</div>

    <div class="product__screen_size">Screen size: {{ product.screen_size }}"</div>
    <div class="product__price">{{ product.price }}:-</div>
  </router-link>
</template>

<script setup>
import placeholderImg from '@/assets/placeholder.webp'

const props = defineProps({
  product: { type: Object, required: true },
  imgHeight: { type: [Number, String], default: 140 },
  placeholder: { type: String, default: placeholderImg },
})

function toCss(v) {
  return typeof v === 'number' ? `${v}px` : v
}

function onError(e) {
  e.target.src = props.placeholder
}
</script>

<style scoped>
.products,
.products:link,
.products:visited,
.products:hover,
.products:active {
  color: inherit;
  text-decoration: none;
}

.products {
  display: flex;
  flex-direction: column;
  background: var(--bg);
  padding: 24px;
  border-radius: 24px;
  text-decoration: none;
  max-width: 360px;
  width: 100%;
}

.product__img {
  width: 100%;
  max-width: calc(var(--img-h) * 1.778);
  min-width: var(--img-h);
  height: var(--img-h);
  object-fit: cover;
  border-radius: 24px;
  align-self: center;
}

.product__name {
  font-weight: bold;
  font-size: 1.4rem;
  text-decoration: none;
  color: inherit;
  margin-top: 16px;
}

.product__manufacturer {
  font-weight: 500;
  font-size: 1.1rem;
  opacity: 80%;
}

.product__screen_size {
  font-weight: 500;
  font-size: 0.9rem;
  opacity: 70%;
  margin-top: 8px;
}

.product__price {
  margin-top: 20px;
  font-weight: bold;
  font-size: 1.3rem;
  color: red;
}
</style>
