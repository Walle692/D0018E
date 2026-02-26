<template>
  <div class="sb_row">
    <select
      class="sb_select"
      :disabled="disabled"
      :value="typeValue"
      @change="emit('update:typeValue', $event.target.value)"
    >
      <option v-for="t in types" :key="t.value" :value="t.value">
        {{ t.label }}
      </option>
    </select>

    <button type="button" class="sb_btn" :disabled="disabled" @click="triggerSearch">Search</button>

    <input
      class="sb_input"
      type="text"
      :disabled="disabled"
      :placeholder="placeholder"
      :value="modelValue"
      @input="emit('update:modelValue', $event.target.value)"
      @keyup.enter="triggerSearch"
    />
  </div>
</template>

<script setup>
const props = defineProps({
  modelValue: { type: String, default: '' },
  typeValue: { type: String, default: 'username' },
  placeholder: { type: String, default: 'Search...' },
  disabled: { type: Boolean, default: false },
  types: {
    type: Array,
    default: () => [
      { value: 'username', label: 'Username' },
      { value: 'orderID', label: 'Order ID' },
      { value: 'userID', label: 'User ID' },
    ],
  },
})

const emit = defineEmits(['update:modelValue', 'update:typeValue', 'search'])

function triggerSearch() {
  emit('search')
}
</script>

<style scoped>
.sb_row {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: nowrap;
  width: 100%;
}

.sb_select,
.sb_input,
.sb_btn {
  height: 42px;
  border-radius: 12px;
  border: 2px solid var(--shadow);
  background: var(--bg);
  font-size: 1rem;
  box-sizing: border-box;
}

.sb_row .sb_btn {
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
  cursor: pointer;
  flex: 0 0 auto;
  white-space: nowrap;
}

.sb_select {
  padding: 0 12px;
  flex: 0 0 170px;
}

.sb_btn {
  width: auto !important;
  display: inline-flex !important;
  flex: 0 0 auto !important;
  white-space: nowrap;
}

.sb_input {
  padding: 0 12px;
  flex: 1 1 auto;
  min-width: 0;
}

.sb_select:disabled,
.sb_btn:disabled,
.sb_input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
