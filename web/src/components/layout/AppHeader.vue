<template>
  <v-app-bar flat height="72" density="comfortable" class="border-b">
    <v-container fluid class="px-4">
      <v-row align="center" no-gutters>
        <RouterLink to="/dashboard" class="d-flex align-center text-decoration-none">
          <v-icon size="28" color="primary" class="mr-2">fas fa-chart-line</v-icon>
          <span class="font-weight-bold d-none d-md-inline">FinPulse Insights</span>
        </RouterLink>
        <v-spacer />
        <!-- Móvil: botón abrir sidebar -->
        <v-btn
          class="d-md-none"
          icon
          variant="text"
          color="primary"
          aria-label="Abrir sidebar"
          @click="ui.setSidebar(true)"
        >
          <v-icon>fas fa-bars</v-icon>
        </v-btn>
        <!-- Dark-mode -->
        <v-btn
          icon
          variant="text"
          color="primary"
          aria-label="Toggle theme"
          @click="ui.toggleDarkMode"
        >
          <v-icon v-if="!ui.dark">fas fa-moon</v-icon>
          <v-icon v-else>fas fa-sun</v-icon>
        </v-btn>
      </v-row>
    </v-container>
  </v-app-bar>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useUI } from '@/stores/ui'
const ui = useUI()
</script>
