<template>
  <nav class="nav">
    <div class="left">
      <span class="brand">MonitorB2B</span>
    </div>

    <div class="right">
      <a href="#" class="link">My profile</a>
      <a href="#" class="link">Settings</a>
      <a href="#" class="link">About</a>
    </div>
  </nav>
</template>

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
}
</style>
