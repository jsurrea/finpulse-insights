<template>
  <v-navigation-drawer
    :rail="!isSidebarOpen"
    :permanent="isDesktop"
    :temporary="!isDesktop"
    :width="256"
    class="border-r border-neutral-200 bg-background"
    :elevation="0"
    v-model="sidebarLocal"
    app
    disable-resize-watcher
    disable-route-watcher
  >
    <v-list nav dense>
      <!-- Logo and toggle for collapsed (desktop only) -->
      <v-list-item class="px-4 py-4 d-flex align-center" lines="one">
        <RouterLink to="/" class="d-flex align-center text-decoration-none flex-grow-1">
          <v-icon :icon="mdiTrendingUp" color="primary" size="28" class="mr-1" />
          <span v-if="isSidebarOpen" class="font-weight-bold font-headline text-body-1">FinPulse</span>
        </RouterLink>
        <v-btn
          icon
          size="small"
          color="default"
          variant="text"
          @click="toggleSidebar"
          aria-label="Toggle sidebar"
          class="ml-2"
        >
          <v-icon :icon="isSidebarOpen ? mdiChevronLeft : mdiChevronRight" />
        </v-btn>
      </v-list-item>
      <!-- Main nav items -->
      <v-list-item
        v-for="item in navItems"
        :key="item.href"
        :to="item.href"
        rounded="lg"
        :active="route.path === item.href || (item.href !== '/dashboard' && route.path.startsWith(item.href))"
        color="primary"
        @click="onNavClick"
        class="mb-1"
      >
        <template #prepend>
          <v-icon :icon="item.icon" size="22" />
        </template>
        <span v-if="isSidebarOpen" class="ml-2">{{ item.label }}</span>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { useUI } from '@/stores/ui'
import {
  mdiTrendingUp, mdiChevronLeft, mdiChevronRight, mdiHomeOutline, mdiNewspaperVariantOutline, mdiChartLine, mdiChartBar, mdiRobotOutline, mdiHeartOutline
} from '@mdi/js'

const ui = useUI()
const isSidebarOpen = computed(() => ui.sidebarOpen)
const toggleSidebar = ui.toggleSidebar
const setSidebarOpen = ui.setSidebarOpen

const route = useRoute()
const navItems = [
  { href: '/dashboard', icon: mdiHomeOutline, label: "Dashboard" },
  { href: '/stocks', icon: mdiNewspaperVariantOutline, label: "Stocks" },
  { href: '/recommendations', icon: mdiChartLine, label: "Recommendations" },
  { href: '/analytics', icon: mdiChartBar, label: "Analytics" },
  { href: '/ai-analyst', icon: mdiRobotOutline, label: "AI Analyst" },
  { href: '/health', icon: mdiHeartOutline, label: "Health" }
]

// Responsive
const isDesktop = ref(window.innerWidth >= 768)
onMounted(() => {
  window.addEventListener('resize', () => {
    isDesktop.value = window.innerWidth >= 768
  })
})

// Control sidebar open/close state with a local variable for v-model
const sidebarLocal = computed({
  get: () => ui.sidebarOpen,
  set: (val: boolean) => ui.setSidebarOpen(val)
})

const onNavClick = () => {
  if (!isDesktop.value) setSidebarOpen(false)
}
</script>
