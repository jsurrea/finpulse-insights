<template>
  <div class="header-right d-flex align-center justify-end pr-8 pl-4" style="height: 100%">
    <span
      class="header-app-name mr-6 font-weight-bold text-primary"
      style="font-family: Inter, system-ui, sans-serif"
    >
      FinPulse Insights
    </span>
    <v-btn icon variant="text" color="primary" aria-label="Toggle theme" @click="ui.toggleDarkMode">
      <v-icon v-if="!ui.dark">fas fa-moon</v-icon>
      <v-icon v-else>fas fa-sun</v-icon>
    </v-btn>
  </div>
</template>
<script setup lang="ts">
import { useUI } from '@/stores/ui'
const ui = useUI()
</script>
<style scoped>
.header-right {
  font-size: 1.2rem;
  font-family: 'Inter', system-ui, sans-serif;
  font-weight: 500;
  color: var(--v-theme-primary);
}
.header-app-name {
  font-size: 1.19rem;
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
