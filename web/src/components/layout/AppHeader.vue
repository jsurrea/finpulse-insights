<template>
  <v-app-bar
    flat
    height="64"
    class="border-b border-neutral-200 bg-background/80"
    elevation="0"
    density="comfortable"
    scroll-behavior="none"
  >
    <!-- Left: Logo and name -->
    <v-row class="align-center flex-nowrap w-100 pa-0 ma-0">
      <!-- Logo and name: show 'FinPulse' on all widths, show 'FinPulse Insights' only md+ -->
      <v-col class="px-0 flex-grow-0 flex-shrink-0 d-flex align-center" cols="auto">
        <RouterLink to="/" class="d-flex align-center text-decoration-none">
          <v-icon :icon="mdiTrendingUp" color="primary" size="28" class="mr-1" />
          <span class="font-weight-bold font-headline text-body-1 d-none d-md-inline">FinPulse Insights</span>
          <span class="font-weight-bold font-headline text-body-1 d-inline d-md-none">FinPulse</span>
        </RouterLink>
      </v-col>
      <!-- Mobile: sidebar trigger -->
      <v-col class="px-0 flex-grow-0 flex-shrink-0 d-flex align-center d-md-none">
        <v-btn icon @click="() => setSidebarOpen(true)" color="default" variant="text" aria-label="Open sidebar">
          <v-icon :icon="mdiMenu" />
        </v-btn>
      </v-col>
      <!-- Spacer -->
      <v-spacer />
      <!-- Right: theme toggle icon -->
      <v-col class="px-0 flex-grow-0 flex-shrink-0">
        <v-btn icon color="default" variant="text" aria-label="Toggle theme" @click="toggleTheme">
          <v-icon :icon="theme === 'light' ? mdiMoonWaningCrescent : mdiWhiteBalanceSunny" />
        </v-btn>
      </v-col>
    </v-row>
  </v-app-bar>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useUI } from '@/stores/ui'
import { mdiTrendingUp, mdiMenu, mdiMoonWaningCrescent, mdiWhiteBalanceSunny } from '@mdi/js'

const ui = useUI()
const theme = ui.darkMode ? 'dark' : 'light'
const toggleTheme = ui.toggleDarkMode
const setSidebarOpen = ui.setSidebarOpen
</script>
