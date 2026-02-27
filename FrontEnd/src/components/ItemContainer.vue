<script setup>
import placeholderImg from '@/assets/placeholder.webp'

const props = defineProps({
  imageSrc: {
    type: String,
    default: '',
  },
  alt: {
    type: String,
    default: '',
  },
  height: {
    type: [String, Number],
    default: 96,
  },
  placeholder: {
    type: String,
    default: placeholderImg,
  },
})

const emit = defineEmits(['img-error'])

function onError(e) {
  e.target.src = props.placeholder
  emit('img-error', e)
}

function toCssSize(value) {
  return typeof value === 'number' ? `${value}px` : value
}
</script>

<template>
  <div class="item_container" :style="{ '--item-height': toCssSize(height) }">
    <img class="item_container__img" :src="imageSrc || placeholder" :alt="alt" @error="onError" />

    <div class="item_container__col item_container__col--left1">
      <slot name="left-1" />
    </div>

    <div class="item_container__col item_container__col--left2">
      <slot name="left-2" />
    </div>

    <div class="item_container__col item_container__col--right1">
      <slot name="right-1" />
    </div>

    <div class="item_container__col item_container__col--right2">
      <slot name="right-2" />
    </div>
  </div>
</template>

<style scoped>
.item_container {
  grid-column: 1/-1;
  display: grid;
  grid-template-columns: subgrid;
  grid-template-areas: 'a b c . d e';
  gap: 20px;
  align-items: center;
  padding: 24px;
  padding-right: 96px;
  border-radius: 24px;
  background: var(--shadow);
  border: solid 1px var(--border);
}

.item_container__img {
  grid-area: a;
  width: var(--item-height);
  height: var(--item-height);
  object-fit: cover;
  border-radius: 12px;
  display: block;
}

.item_container__col--right1,
.item_container__col--right2 {
  text-align: right;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
}

.item_container__col--left1 {
  grid-area: b;
}

.item_container__col--left2 {
  grid-area: c;
}
.item_container__col--right1 {
  grid-area: d;
}
.item_container__col--right2 {
  grid-area: e;
}

@media (max-width: 900px) {
  .item_container {
    padding-right: 24px;
    grid-template-areas:
      'a a . . e d'
      'b c . . . .';
  }

  .item_container__img {
    width: calc(var(--item-height) / 2);
    height: calc(var(--item-height) / 2);
  }
}

@media (max-width: 350px) {
}
</style>
