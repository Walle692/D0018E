<template>
  <div class="container">
    <div class="card">
      <h1>User page</h1>

      <p v-if="loading">Loading...</p>
      <p v-else-if="error" class="error">{{ error }}</p>
      <p v-else>Welcome, <b>{{ me?.user }}</b>, your role is <b>{{ me?.role }}</b>!</p>
    </div>
  </div>
</template>


<script setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { getMe } from "@/services/auth";

const router = useRouter();

const loading = ref(true);
const error = ref("");
const me = ref(null);

onMounted(async () => {
  try {
    me.value = await getMe();
  } catch (e) {
    error.value = e?.message || "Failed to load user";
    // optional: if not logged in, kick back to login
    router.push("/");
  } finally {
    loading.value = false;
  }
});
</script>

<style>
.card { padding: 20px; border: 1px solid #ccc; border-radius: 10px; }
.error { margin-top: 10px; }
</style>
