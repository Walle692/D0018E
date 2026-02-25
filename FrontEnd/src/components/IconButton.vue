<script setup>
const props = defineProps({
  src: {
    type: String,
    required: true,
  },
  alt: {
    type: String,
    default: 'icon button',
  },
  size: {
    type: [String, Number],
    default: 32,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['click'])

function toCssSize(value) {
  return typeof value === 'number' ? `${value}px` : value
}

function handleClick(event) {
  if (props.disabled) return
  emit('click', event)
}
</script>

<template>
  <button
    type="button"
    class="icon_btn"
    :disabled="disabled"
    @click="handleClick"
    :aria-label="alt"
  >
    <img
      :src="src"
      :alt="alt"
      class="icon_btn__img"
      :style="{ width: toCssSize(size), height: toCssSize(size) }"
    />
  </button>
</template>

<style scoped>
.icon_btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
  background: var(--surface);
  padding: 8px;
  border-radius: 999px;
  cursor: pointer;
}

.icon_btn__img {
  display: block;
  opacity: 0.6;
}

.icon_btn:hover .icon_btn__img {
  opacity: 1;
}
</style>
