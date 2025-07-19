<template>
  <div class="layout-grid">
    <AppSidebar />
    <header class="header-grid">
      <AppHeaderLeft />
      <AppHeaderRight />
    </header>
    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import AppSidebar from '@/components/layout/AppSidebar.vue'
import AppHeaderLeft from '@/components/layout/AppHeaderLeft.vue'
import AppHeaderRight from '@/components/layout/AppHeaderRight.vue'
</script>

<style scoped>
.layout-grid {
  min-height: 100vh;
  display: grid;
  grid-template-columns: auto 1fr;
  grid-template-rows: 72px 1fr;
  grid-template-areas:
    "sidebar header"
    "sidebar main";
  background: var(--v-theme-background);
}
.header-grid {
  grid-area: header;
  display: flex;
  height: 72px;
  border-bottom: 1.5px solid var(--v-theme-outline, #E5E7EB);
  background: var(--v-theme-surface);
  align-items: stretch;
}
.main-content {
  grid-area: main;
  padding: 2rem;
  overflow: auto;
  background: var(--v-theme-surface);
}
</style>
