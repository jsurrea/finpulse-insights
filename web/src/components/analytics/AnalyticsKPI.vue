<template>
  <v-card class="h-100 analytics-kpi">
    <v-card-title class="pa-4 pb-2">
      <div class="d-flex align-center justify-space-between">
        <span class="text-sm font-medium text-medium-emphasis">{{ title }}</span>
        <v-icon :icon="icon" size="16" class="text-medium-emphasis" />
      </div>
    </v-card-title>
    <v-card-text class="pa-4 pt-0">
      <div class="text-3xl font-weight-bold mb-1" style="font-family: 'Inter', system-ui, sans-serif;">
        {{ formattedValue }}
      </div>
      <div v-if="change !== undefined" class="d-flex align-center">
        <v-icon
          :icon="change >= 0 ? 'fas fa-arrow-up' : 'fas fa-arrow-down'"
          :color="change >= 0 ? 'success' : 'error'"
          size="12"
          class="mr-1"
        />
        <span
          :class="change >= 0 ? 'text-success' : 'text-error'"
          class="text-caption font-weight-medium"
        >
          {{ Math.abs(change).toFixed(1) }}%
        </span>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface AnalyticsKPIProps {
  title: string
  value: number
  icon: string
  format?: 'number' | 'percentage' | 'currency'
  change?: number
}

const props = withDefaults(defineProps<AnalyticsKPIProps>(), {
  format: 'number'
})

const formattedValue = computed(() => {
  switch (props.format) {
    case 'percentage':
      return `${props.value.toFixed(1)}%`
    case 'currency':
      return `$${props.value.toLocaleString()}`
    default:
      return props.value.toLocaleString()
  }
})
</script>

<style scoped>
.analytics-kpi {
  transition: all 0.2s;
}

.analytics-kpi:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(var(--v-theme-shadow-rgb), 0.15);
}
</style>
