<template>
  <v-navigation-drawer
    v-model="ui.sidebar"
    permanent
    width="260"
    class="border-r"
    style="background: var(--v-theme-background)"
    app
  >
    <div class="d-flex justify-end px-2 py-2">
      <v-btn
        icon
        variant="text"
        color="primary"
        size="small"
        aria-label="Colapsar sidebar"
        @click="ui.toggleSidebar"
      >
        <v-icon v-if="ui.sidebar">fas fa-chevron-left</v-icon>
        <v-icon v-else>fas fa-chevron-right</v-icon>
      </v-btn>
    </div>
    <v-divider class="mb-2" />
    <!-- Navegación -->
    <v-list nav density="comfortable">
      <v-list-item
        v-for="item in items"
        :key="item.to"
        :to="item.to"
        :active="$route.path.startsWith(item.to)"
        rounded="lg"
        active-class="bg-primary text-white"
        class="px-2"
        @click="onNav"
      >
        <template #prepend>
          <v-icon size="20">{{ item.icon }}</v-icon>
        </template>
        <v-list-item-title v-if="ui.sidebar">{{ item.label }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { useUI } from '@/stores/ui'
const ui = useUI()
const items = [
  { to: '/dashboard', icon: 'fas fa-home', label: 'Dashboard' },
  { to: '/stocks', icon: 'fas fa-newspaper', label: 'Stocks' },
  { to: '/recommendations', icon: 'fas fa-chart-line', label: 'Recommendations' },
  { to: '/analytics', icon: 'fas fa-chart-bar', label: 'Analytics' },
  { to: '/ai-analyst', icon: 'fas fa-robot', label: 'AI Analyst' },
  { to: '/health', icon: 'fas fa-heart', label: 'Health' }
]
function onNav () {
  if (window.innerWidth < 768) ui.setSidebar(false)
}
</script>
