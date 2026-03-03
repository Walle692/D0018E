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

          <p v-if="reviewLoading" class="state">Loading Reviews...</p>
          <p v-else-if="reviewError" class="state error">{{ error }}</p>
          <template v-else>
            <div v-if="reviews.length === 0" class="state">No reviews yet.</div>
            <div v-else>
              <div v-for="review in reviews" :key="review.review_id" class="reviews__review">
                <div class="review__rating">Rating: {{ review.rating }} / 5</div>
                <div class="review__comment">{{ review.comment }}</div>
              </div>
            </div>
          </template>
        </div>

        <div class="write-review">
            <div class="card">
              <form @submit.prevent="onSubmit">
                <label class="field">
                  <span>Comment</span>
                  <input
                    v-model.trim="form.comment"
                    type="text"
                    placeholder="Write your review here"
                    autocomplete=""
                  />
                </label>

                <label class="field">
                  <span>Rating</span>
                  <select v-model="form.rating">
                    <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                  </select>
                </label>

                <button type="submit" :disabled="loading || !canSubmit">
                  {{ loading ? 'Creating...' : 'Creating review' }}
                </button>

                <p v-if="success" class="success">{{ success }}</p>
                <p v-if="error" class="error">{{ error }}</p>
              </form>
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
import { computed, onMounted, ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { getProductById } from '@/services/products'
import { addToBasket as addToBasketApi } from '@/services/basket'
import { getReviewsForProduct, createReview } from '@/services/reviews'
import placeholderImg from '@/assets/placeholder.webp'

const router = useRouter()
const product = ref(null)
const loading = ref(true)
const error = ref('')
const success = ref('')
const reviews = ref([])
const reviewLoading = ref(true)
const reviewError = ref('')

const form = reactive({
  comment: '',
  rating: 5,
})

const canSubmit = computed(() => form.rating >= 1 && form.rating <= 5 && product.value !== null)

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
    const productId = router.currentRoute.value.params.id
    console.log('Fetching product with ID:', productId)
    const data = await getProductById(productId)
    console.log('Received product data:', data)
    product.value = data
  } catch (e) {
    console.error('Error fetching product:', e)
    error.value = e?.message || 'Failed to load product'
  } finally {
    loading.value = false
  }
}

async function fetchReviews() {
  reviewLoading.value = true
  reviewError.value = ''
  reviews.value = []

  try {
    const data = await getReviewsForProduct(router.currentRoute.value.params.id)
    reviews.value = data
  } catch (e) {
    console.error('Failed to fetch reviews:', e)
  } finally {
    reviewLoading.value = false
  }
}

async function onSubmit() {
  error.value = ''
  success.value = ''
  loading.value = true

  try {
    if (!product.value) {
      throw new Error('Product not loaded')
    }
    await createReview({ Product_id: product.value.product_id, Comment: form.comment, Rating: form.rating })
    success.value = 'Review created successfully!'
    form.comment = ''
    form.rating = 5
    fetchReviews() // Refresh reviews after submission
  } catch (e) {
    error.value = e?.message || 'Failed to create review'
  } finally {
    loading.value = false
  }
}

onMounted(fetchProduct)
onMounted(fetchReviews)
</script>
