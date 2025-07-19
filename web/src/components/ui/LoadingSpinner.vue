<template>
  <div :class="containerClasses">
    <v-progress-circular
      :indeterminate="!progress"
      :model-value="progress"
      :size="spinnerSize"
      :width="spinnerWidth"
      :color="spinnerColor"
      class="loading-spinner"
    />
    <div v-if="text" :class="textClasses">
      {{ text }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface LoadingSpinnerProps {
  size?: 'small' | 'medium' | 'large'
  color?: 'primary' | 'secondary' | 'success' | 'warning' | 'error'
  text?: string
  progress?: number
  overlay?: boolean
  centered?: boolean
}

const props = withDefaults(defineProps<LoadingSpinnerProps>(), {
  size: 'medium',
  color: 'primary',
  overlay: false,
  centered: true
})

const spinnerSize = computed(() => {
  switch (props.size) {
    case 'small': return 24
    case 'large': return 64
    default: return 40
  }
})

const spinnerWidth = computed(() => {
  switch (props.size) {
    case 'small': return 3
    case 'large': return 6
    default: return 4
  }
})

const spinnerColor = computed(() => props.color)

const containerClasses = computed(() => [
  'loading-spinner-container',
  {
    'loading-overlay': props.overlay,
    'loading-centered': props.centered,
    'loading-inline': !props.centered
  }
])

const textClasses = computed(() => [
  'loading-text',
  {
    'text-caption': props.size === 'small',
    'text-body-2': props.size === 'medium',
    'text-body-1': props.size === 'large'
  }
])
</script>

<style scoped>
.loading-spinner-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  font-family: 'Inter', system-ui, sans-serif;
}

.loading-centered {
  justify-content: center;
  min-height: 120px;
}

.loading-inline {
  display: inline-flex;
}

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(var(--v-theme-background-rgb), 0.8);
  backdrop-filter: blur(2px);
  z-index: 9999;
  justify-content: center;
  align-items: center;
}

.loading-text {
  color: rgb(var(--v-theme-on-surface-variant));
  text-align: center;
  font-weight: 500;
}

.loading-spinner {
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.7; }
}
</style>
