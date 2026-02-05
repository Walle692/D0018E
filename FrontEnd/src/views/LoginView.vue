<template>
  <div class="container">

    <div class="infoCard">
      <h1>MonitorB2B</h1>
      <p>For bizzness by bizzness the best place to sell and buy monitor for bizznesses</p>
    </div>


    <div class="loginCard">
      <h1>Login</h1>
      <form @submit.prevent="onSubmit">
        <label class="inputField">
          <input v-model.trim="form.username" type="text" placeholder="Username">
        </label>
        
        <label class="inputField">
          <input v-model.trim="form.password" type="password" placeholder="Password">
        </label>

        <button type="submit" :disabled="loading">
          {{ loading ? "Logging in..." : "Log in" }}
        </button>
        <p v-if="error" class="error">{{ error }}</p>
      </form> 
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

import { reactive } from "vue";
import { login } from "@/services/auth";

const form = reactive({ username: "", password: "" });

async function onSubmit() {
  try{
    await login({username: form.username, password: form.password})
    router.push("/user")
  }catch(error){
    error.value = e?.message || "Login failed";
  }
}

const router = useRouter();

function goToUser() {
  router.push("/user");
}
</script>

<style>

.container{
  display: flex;
  justify-content: center;
  background: var(--bg);
  gap: 25px;
  padding: 25px 25px;
}


.infoCard,
.loginCard {
  width: 100%;
  max-width: 400px;
  padding: 25px;
  border: 1px solid var(--border);
  border-radius: 10px;
  background: var(--surface);

}

button {
  width: 100%;
  margin-top: 12px;
  padding: 10px 12px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: var(--primary);
  color: var(--text);
  cursor: pointer;
}

.inputField {
  display: block;
  margin: 20px 0px;
}

input{
  display: block;
  width: 100%;
  box-sizing: border-box;
  margin-top: 6px;
  border-radius: 8px;
  padding: 8px 10px;
  border: 1px solid #bbb;
  border-radius: px;
}

</style>