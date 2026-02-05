<template>
  <div class="page">
    <h1>Your Account</h1>

    <p v-if="loading">Loading...</p>
    <p v-else-if="error" class="error">{{ error }}</p>

    <div v-else>
      <p>Welcome, <b>{{ me.user }}</b> ({{ me.role }})</p>

      <div class="grid">

        <AccountCard
          v-show="!isSellerOrAdmin"
          title="Your Orders"
          subtitle="Track, return or buy again"
          @click="router.push('/orders')"
        />
        <AccountCard
          title="Login & Security"
          subtitle="Manage password and security settings"
          @click="router.push('/security')"
        />
        <AccountCard
          v-show="!isAdmin"
          title="Your Addresses"
          subtitle="Edit addresses and preferences"
          @click="router.push('/addresses')"
        />


        <AccountCard
          v-if="isSeller"
          title="My Listings"
          subtitle="Create and manage listings"
          @click="router.push('/listings')"
        />
        <AccountCard
          v-if="isSeller"
          title="Pending Orders"
          subtitle="Orders you need to handle"
          @click="router.push('/orders/pending')"
        />

        <!-- Admin only -->
        <AccountCard
          v-if="isAdmin"
          title="Create User"
          subtitle="Create new accounts"
          @click="() => {router.push('/admin/create-user') }"
        />
        <AccountCard
          v-if="isAdmin"
          title="All Orders"
          subtitle="View all orders"
          @click="router.push('/admin/orders')"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { getMe } from "@/services/auth";
import AccountCard from "@/components/AccountCard.vue";

const router = useRouter();

const loading = ref(true);
const error = ref("");
const me = ref(null);

const isAdmin = computed(() => me.value?.role === "admin");
const isSellerOrAdmin = computed(() => ["seller", "admin"].includes(me.value?.role));
const isSeller = computed(() => me.value?.role == "seller")

onMounted(async () => {
  try {
    me.value = await getMe();
  } catch (e) {
    error.value = e?.message || "Failed to load user";
    router.push("/");
  } finally {
    loading.value = false;
  }
});
</script>

<style>
.page {
  padding: 25px;

}

.grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(220px, 1fr));
  gap: 16px;
}
.error { margin-top: 10px; }


@media (max-width: 900px) {
  .grid {
    grid-template-columns: repeat(2, minmax(220px, 1fr));
  }
}

/* Phone -> 1 column */
@media (max-width: 520px) {
  .grid {
    grid-template-columns: repeat(1, minmax(220px, 1fr));
  }
}

</style>

