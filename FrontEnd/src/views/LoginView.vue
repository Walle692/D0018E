<template>
  <div>
    <h1>Login Page</h1>
    <p>This is the login page.</p>
    <button @click="goToUser">  Go to user page</button>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

function goToUser() {
  router.push("/user");
}
</script>