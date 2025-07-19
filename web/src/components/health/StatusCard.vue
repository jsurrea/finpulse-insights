<template>
  <v-card class="h-100 status-card">
    <v-card-title class="d-flex justify-space-between align-center pa-4 pb-2">
      <span class="text-sm font-medium">{{ title }}</span>
      <v-icon
        :icon="icon"
        :color="iconColor"
        size="20"
      />
    </v-card-title>
    <v-card-text class="pa-4 pt-0">
      <div
        :class="[
          'text-2xl font-weight-bold text-uppercase mb-2',
          statusTextClass
        ]"
        style="font-family: 'Inter', system-ui, sans-serif;"
      >
        {{ value }}
      </div>
      <p class="text-caption text-medium-emphasis">
        {{ description }}
      </p>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface StatusCardProps {
  title: string
  value: string | number
  description: string
  status: 'ok' | 'error' | 'connected' | 'disconnected' | 'info'
  icon: string
}

const props = defineProps<StatusCardProps>()

const iconColor = computed(() => {
  switch (props.status) {
    case 'ok':
    case 'connected':
      return 'success'
    case 'error':
    case 'disconnected':
      return 'error'
    default:
      return 'primary'
  }
})

const statusTextClass = computed(() => {
  switch (props.status) {
    case 'ok':
    case 'connected':
      return 'text-success'
    case 'error':
    case 'disconnected':
      return 'text-error'
    default:
      return 'text-primary'
  }
})
</script>

<style scoped>
.status-card {
  transition: all 0.2s ease-in-out;
}

.status-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(var(--v-theme-shadow-rgb), 0.15);
}
</style>
