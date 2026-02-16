<!-- src/views/CreateUserView.vue -->
<template>
  <div class="page">
    <h1>Create User</h1>
    <p class="subtitle">Create a new account and assign a role.</p>

    <div class="card">
      <form @submit.prevent="onSubmit">
        <label class="field">
          <span>Username</span>
          <input
            v-model.trim="form.username"
            type="text"
            placeholder="e.g. john123"
            autocomplete="username"
          />
        </label>

        <label class="field">
          <span>Password</span>
          <input
            v-model="form.password"
            type="password"
            placeholder="Enter a password"
            autocomplete="new-password"
          />
        </label>

        <label class="field">
          <span>Role</span>
          <select v-model="form.role">
            <option value="buyer">buyer</option>
            <option value="seller">seller</option>
            <option value="admin">admin</option>
          </select>
        </label>

        <button type="submit" :disabled="loading || !canSubmit">
          {{ loading ? 'Creating...' : 'Create user' }}
        </button>

        <p v-if="success" class="success">{{ success }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </form>
    </div>

    <button class="linkBtn" type="button" @click="router.push('/user')">← Back to account</button>
  </div>
</template>

<script setup>
import { computed, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { createUser } from '@/services/auth'

const router = useRouter()

const form = reactive({
  username: '',
  password: '',
  role: 'buyer',
})

const loading = ref(false)
const error = ref('')
const success = ref('')

const canSubmit = computed(() => form.username.length > 0 && form.password.length > 0)

async function onSubmit() {
  // Prepared for backend integration:
  // Later you will call something like:
  // await adminCreateUser({ username: form.username, password: form.password, role: form.role });

  error.value = ''
  success.value = ''
  loading.value = true

  try {
    await createUser({ username: form.username, password: form.password, role: form.role })
  } catch (e) {
    error.value = e?.message || 'User creation failed'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.page {
  padding: 24px;
  max-width: 720px;
  margin: 0 auto;
}

.subtitle {
  margin-top: 6px;
  opacity: 0.8;
}

.card {
  margin-top: 16px;
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 16px;
  background: var(--surface);
}

.field {
  display: block;
  margin: 14px 0;
}

.field span {
  display: block;
  margin-bottom: 6px;
}

input,
select {
  width: 100%;
  box-sizing: border-box;
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid #bbb;
}

button {
  width: 100%;
  margin-top: 12px;
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid #bbb;
  cursor: pointer;
}

.success {
  margin-top: 10px;
}

.error {
  margin-top: 10px;
}

.linkBtn {
  width: auto;
  margin-top: 14px;
}
</style>
