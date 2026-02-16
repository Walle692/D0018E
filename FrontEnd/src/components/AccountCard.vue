<template>
  <div class="card" @click="$emit('click')" role="button" tabindex="0">
    <h2 class="cardTitle">{{ title }}</h2>
    <p class="cardSub">{{ subtitle }}</p>
  </div>
</template>

<script setup>
defineProps({
  title: String,
  subtitle: String,
})
const emit = defineEmits(['click'])
</script>

<style scoped>
.card {
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 16px;
  background: var(--surface);
}
.cardTitle {
  margin: 0 0 6px;
}
.cardSub {
  margin: 0;
  opacity: 1;
}

.card:hover {
  border-color: var(--primary-hover);
  cursor: pointer;
}
</style>
