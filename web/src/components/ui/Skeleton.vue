<template>
  <div :class="skeletonClasses" :style="skeletonStyles">
    <div v-if="variant === 'text'" class="skeleton-lines">
      <div
        v-for="i in lines"
        :key="i"
        class="skeleton-line"
        :style="{ width: i === lines ? `${lastLineWidth}%` : '100%' }"
      />
    </div>
    <div v-else class="skeleton-shape" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface SkeletonProps {
  variant?: 'text' | 'rectangular' | 'circular' | 'rounded'
  width?: string | number
  height?: string | number
  lines?: number
  lastLineWidth?: number
  animation?: 'pulse' | 'wave' | 'none'
}

const props = withDefaults(defineProps<SkeletonProps>(), {
  variant: 'text',
  lines: 1,
  lastLineWidth: 60,
  animation: 'pulse',
})

const skeletonClasses = computed(() => [
  'skeleton',
  `skeleton-${props.variant}`,
  `skeleton-${props.animation}`,
])

const skeletonStyles = computed(() => {
  const styles: Record<string, string> = {}

  if (props.width) {
    styles.width = typeof props.width === 'number' ? `${props.width}px` : props.width
  }

  if (props.height) {
    styles.height = typeof props.height === 'number' ? `${props.height}px` : props.height
  }

  return styles
})
</script>

<style scoped>
.skeleton {
  background: rgb(var(--v-theme-surface-variant));
  background-color: rgba(var(--v-theme-on-surface-rgb), 0.1);
  border-radius: 4px;
}

.skeleton-text {
  width: 100%;
}

.skeleton-rectangular {
  width: 100%;
  height: 120px;
}

.skeleton-circular {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.skeleton-rounded {
  border-radius: 8px;
}

.skeleton-lines {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.skeleton-line {
  height: 1em;
  border-radius: 4px;
  background: rgb(var(--v-theme-surface-variant));
}

.skeleton-shape {
  width: 100%;
  height: 100%;
  border-radius: inherit;
  background: rgb(var(--v-theme-surface-variant));
}

/* Animations */
.skeleton-pulse {
  animation: skeleton-pulse 2s ease-in-out infinite;
}

.skeleton-wave {
  position: relative;
  overflow: hidden;
}

.skeleton-wave::after {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  transform: translateX(-100%);
  background: linear-gradient(
    90deg,
    transparent,
    rgba(var(--v-theme-on-surface-rgb), 0.2),
    transparent
  );
  animation: skeleton-wave 2s ease-in-out infinite;
}

@keyframes skeleton-pulse {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

@keyframes skeleton-wave {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}
</style>
