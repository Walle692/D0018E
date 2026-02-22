<template>
  <nav class="nav">
    <div class="left">
      <span class="brand">MonitorB2B</span>
    </div>

    <div class="right">
      <RouterLink v-if="loggedIn" to="/user" class="link"> My profile </RouterLink>
      <RouterLink to="/user" class="link"> About </RouterLink>
      <href v-if="loggedIn" @click="handleLogout" class="link">Logout</href>
    </div>
  </nav>
</template>

<script setup>
import router from '@/router'
import { getMe, logout } from '@/services/auth'
import { onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const loggedIn = ref(false)
const route = useRoute()

async function checkAuth() {
  try {
    await getMe()
    loggedIn.value = true
  } catch (e) {
    loggedIn.value = false
  }
}

async function handleLogout() {
  try {
    await logout()
    loggedIn.value = false
    router.push('/')
  } catch (e) {
    console.error(e)
  }
}

onMounted(checkAuth)
//Rerun script when route changes
watch(route, checkAuth)
</script>

<style scoped>
.nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid var(--border);
  background: var(--nav);
}

.brand {
  font-weight: 700;
  font-size: 18px;
  color: var(--nav-text);
}

.right {
  display: flex;
  gap: 14px;
  align-items: center;
}

.link {
  text-decoration: none;
  color: var(--nav-text);
  font-size: 14px;
}

.link:hover {
  text-decoration: underline;
  color: var(--hover);
  cursor: pointer;
}
</style>
