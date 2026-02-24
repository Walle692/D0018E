<!-- src/views/ProductView.vue -->
<template>
  <div class="page">
    <div>
      <p v-if="loading" class="state">Loading product...</p>
      <p v-else-if="error" class="state error">{{ error }}</p>
      <template v-else>
        <div v-if="product" class="layout">
          <div class="layout__product">
            <div class="product__header">
              <div class="product__name">{{ product.product_name || 'Unnamed product' }}</div>
              <div class="product__manufacture">
                {{ product.manufacturer || 'Unnamed product' }}
              </div>
              <div>★★★★</div>
            </div>
            <div class="product__image_area">
              <img
                class="product__image"
                v-if="product.picture_url"
                :src="product.picture_url"
                alt="product image"
                @error="onImgError"
              />
            </div>
            <div class="product__desc">
              <div class="product__price">Price: <br />${{ product.price }}</div>
              <div class="product__stock">Stock <br />{{ product.stock }}</div>
              <div class="product__size">Screen size: <br />{{ product.screen_size }}"</div>
              <div class="description">
                Description: <br />{{ product.description || 'No description available.' }}
              </div>
            </div>
          </div>

          <div class="purchase_card">
            <div class="purchase__unit_price">
              Unit Price: <br />
              ${{ product.price }}
            </div>

            <div class="purchase__info">
              <div class="info__left">
                Units: <br />
                <input type="number" v-model="quantity" min="1" @input="preventBadInput" />
              </div>

              <div class="info__right">
                Total Price: <br />
                ${{ totalPrice }}
              </div>
            </div>

            <div>
              <button type="button" @click="addToBasket(product.product_id, quantity)">
                Add to basket
              </button>
              <button class="linkBtn" type="button" @click="router.push('/products')">
                Browse other products
              </button>
            </div>
          </div>
        </div>
        <div class="reviews">
          <div class="reviews__header">Reviews</div>

          <!-- Place holder until fetch orders API is done-->
          <div class="reviews__review">
            <div>
              <div class="review__username">testcompany.net</div>
              <div class="review__rating">★★★★</div>
            </div>

            <div class="review__test">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque rhoncus fermentum est
              vitae auctor. Duis pulvinar, metus ac sagittis convallis, erat dui pulvinar tellus, at
              scelerisque quam tellus eget augue. Nulla facilisi. Donec efficitur pretium rutrum.
              Praesent aliquet interdum accumsan. In placerat lectus nibh, mollis pellentesque
              lectus egestas nec. Aenean a congue magna. Suspendisse feugiat, orci nec ultricies
              efficitur, purus nulla sagittis tortor, sit amet tempus tellus leo a ligula. Integer
              luctus, tortor in sagittis hendrerit, tellus elit interdum turpis, eu dictum dolor
              lorem eget enim. Etiam eget feugiat purus, eget aliquam neque. Ut a tempus ligula, vel
              congue mauris. Aenean feugiat magna orci, quis accumsan sapien imperdiet vitae. Proin
              sit amet lacus eu libero sollicitudin euismod. In hac habitasse platea dictumst.
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped>
.reviews {
  margin-top: 24px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  background-color: var(--surface);
  border-radius: 24px;
  gap: 24px;
  box-shadow: 5px 5px var(--shadow);
}

.reviews__header {
  font-weight: bold;
  font-size: 2rem;
}

.reviews__review {
  display: flex;
  flex-direction: column;
  background-color: var(--bg);
  padding: 24px;
  border-radius: 24px;
  gap: 16px;
}

.review__username {
  font-weight: 600;
  font-size: 1.3rem;
}
.review__rating {
  font-size: 1.2rem;
}
.layout {
  display: grid;
  grid-template-areas:
    'product buy'
    'product .';
  grid-template-columns: 3fr 2fr;
  gap: 24px;
}

.layout__product {
  display: grid;
  grid-area: product;
  grid-template-areas:
    'head head'
    'image desc';
  grid-template-columns: 2fr 1fr;
  background-color: var(--surface);
  padding: 24px;
  border-radius: 24px;
  box-shadow: 5px 5px var(--shadow);
  gap: 24px;
}

.product__header {
  grid-area: head;
}

.product__name {
  font-weight: bold;
  font-size: 3rem;
}

.product__manufacture {
  font-weight: 400;
  font-size: 1.5rem;
}
.product__image_area {
  grid-area: image;
  display: flex;
}
.product__desc {
  grid-area: desc;
  display: flex;
  flex-direction: column;
  font-weight: 500;
  font-size: 1.4rem;
  gap: 24px;
  background-color: var(--bg);
  padding: 24px;
  border-radius: 24px;
}
.product__image {
  width: 100%;
  object-fit: cover;
}
.purchase_card {
  grid-area: buy;
  display: flex;
  flex-direction: column;
  background-color: var(--surface);
  padding: 24px;
  border-radius: 24px;
  box-shadow: 5px 5px var(--shadow);
  gap: 24px;
}

.purchase__unit_price {
  font-size: 1.2rem;
  font-weight: 600;
}

.purchase__info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: var(--bg);
  padding: 16px;
  border-radius: 16px;
  gap: 16px;
}

.info__left,
.info__right {
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-weight: 500;
  font-size: 1rem;
}

.info__right {
  text-align: right;
  font-weight: 700;
  font-size: 1.2rem;
}

input[type='number'] {
  width: 80px;
  padding: 8px 12px;
  border: 2px solid var(--shadow);
  border-radius: 10px;
  font-size: 1rem;
  font-family: inherit;
  background-color: var(--surface);
  color: inherit;
  text-align: center;
  outline: none;
}

input:focus {
  outline-style: double;
  outline-color: var(--hover);
}

@media (max-width: 900px) {
  .layout {
    grid-template-areas:
      'product'
      'buy';
    grid-template-columns: 1fr;
  }

  .layout__product {
    grid-template-areas:
      'head'
      'image'
      'desc';
    grid-template-columns: 1fr;
  }
}
</style>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getProductById } from '@/services/products'
import { addToBasket as addToBasketApi } from '@/services/basket'
import placeholderImg from '@/assets/placeholder.webp'

const router = useRouter()
const product = ref(null)
const loading = ref(true)
const error = ref('')

const onImgError = (e) => {
  e.target.src = placeholderImg
}

const quantity = ref(1)
const totalPrice = computed(() => (product.value?.price * quantity.value).toFixed(2))

function preventBadInput() {
  const val = parseInt(quantity.value)
  if (isNaN(val) || val < 1) {
    quantity.value = 1
  } else if (val >= product.value?.stock) {
    quantity.value = product.value?.stock
  } else {
    quantity.value = val
  }
}
//Api calls
async function addToBasket(productId, quantity) {
  try {
    await addToBasketApi(productId, quantity)
    router.push('/basket')
  } catch (e) {
    alert(e?.message || 'Failed to add product to basket')
  }
}

async function fetchProduct() {
  loading.value = true
  error.value = ''
  product.value = null

  try {
    const data = await getProductById(router.currentRoute.value.params.id)
    product.value = data
  } catch (e) {
    error.value = e?.message || 'Failed to load product'
  } finally {
    loading.value = false
  }
}

onMounted(fetchProduct)
</script>
