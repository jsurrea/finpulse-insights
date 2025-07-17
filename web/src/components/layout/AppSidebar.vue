<template>
  <aside
    :class="['app-sidebar', { collapsed: !ui.sidebar }]"
    aria-label="Sidebar"
  >
    <div class="sidebar-toggle-row">
      <span class="sidebar-brand" v-if="ui.sidebar">FinPulse</span>
      <v-btn
        icon
        variant="text"
        color="primary"
        aria-label="Colapsar/expandir sidebar"
        @click="ui.toggleSidebar"
        class="ml-auto sidebar-toggle-btn"
      >
        <v-icon size="22">
          <template v-if="ui.sidebar">fas fa-chevron-left</template>
          <template v-else>fas fa-bars</template>
        </v-icon>
      </v-btn>
    </div>
    <v-list nav dense class="sidebar-menu">
      <v-list-item
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        :active="$route.path === item.to"
        rounded="xl"
        class="sidebar-item"
        :class="{ 'collapsed': !ui.sidebar }"
        active-class="bg-primary text-white"
        @click="onNav"
      >
        <template #prepend>
          <v-icon size="22" color="primary" class="sidebar-icon-center">{{ item.icon }}</v-icon>
        </template>
        <transition name="fade-x">
          <v-list-item-title v-if="ui.sidebar" class="sidebar-label">{{ item.label }}</v-list-item-title>
        </transition>
      </v-list-item>
    </v-list>
  </aside>
</template>

<script setup lang="ts">
import { useUI } from '@/stores/ui'
const ui = useUI()
const navItems = [
  { to: '/dashboard', icon: 'fas fa-home', label: 'Dashboard' },
  { to: '/stocks', icon: 'fas fa-newspaper', label: 'Stocks' },
  { to: '/recommendations', icon: 'fas fa-chart-line', label: 'Recommendations' },
  { to: '/analytics', icon: 'fas fa-chart-bar', label: 'Analytics' },
  { to: '/ai-analyst', icon: 'fas fa-robot', label: 'AI Analyst' },
  { to: '/health', icon: 'fas fa-heart', label: 'Health' }
]
function onNav() { if (window.innerWidth < 768) ui.setSidebar(false) }
</script>

<style scoped>
.app-sidebar {
  grid-area: sidebar;
  width: 220px;
  min-width: 56px;
  background: var(--v-theme-surface);
  border-right: 3px solid var(--v-theme-primary);
  transition: width .24s cubic-bezier(.4,0,.2,1), background .2s;
  display: flex;
  flex-direction: column;
  align-items: stretch;
}
.app-sidebar.collapsed {
  width: 60px;
}

.sidebar-toggle-row {
  display: flex;
  align-items: center;
  height: 72px;
  padding: 0 16px;
  border-bottom: 1px solid var(--v-theme-outline, #E5E7EB);
  font-family: 'Inter', system-ui, sans-serif;
  font-weight: 700;
  font-size: 1.15rem;
  background: var(--v-theme-surface);
  letter-spacing: 0.01em;
}
.sidebar-brand {
  margin-right: 10px;
  user-select: none;
  color: var(--v-theme-primary);
}
.sidebar-toggle-btn {
  margin-left: auto;
}
.sidebar-menu {
  padding-top: 10px;
}
.sidebar-label {
  margin-left: 10px;
  font-family: 'Inter', system-ui, sans-serif;
  font-size: 1.05rem;
  font-weight: 500;
  color: var(--v-theme-primary);
  user-select: none;
}
.sidebar-icon-center {
  min-width: 28px;
  text-align: center;
}
.sidebar-item {
  margin-bottom: 8px;
  display: flex;
  align-items: center;
  transition: background .15s, color .15s;
}
.sidebar-item.v-list-item--active,
.sidebar-item:hover {
  background: rgba(var(--v-theme-primary-rgb),0.13) !important;
  color: var(--v-theme-primary) !important;
}
.fade-x-enter-active, .fade-x-leave-active {
  transition: opacity 0.21s cubic-bezier(.4,0,.2,1);
}
.fade-x-enter-from, .fade-x-leave-to {
  opacity: 0;
}
</style>
