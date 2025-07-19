<template>
  <div class="empty-state">
    <div class="empty-state-content">
      <v-icon
        :icon="computedIcon"
        :size="iconSize"
        :color="iconColor"
        class="empty-icon mb-4"
      />
      <h3 class="text-h6 font-weight-bold mb-2">{{ title }}</h3>
      <p class="text-body-2 text-medium-emphasis mb-6">{{ description }}</p>

      <div v-if="$slots.actions" class="empty-actions">
        <slot name="actions" />
      </div>

      <div v-else-if="actionText && actionHandler" class="empty-actions">
        <v-btn
          :color="actionColor"
          :variant="actionVariant"
          :prepend-icon="actionIcon"
          @click="actionHandler"
        >
          {{ actionText }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface EmptyStateProps {
  title: string
  description: string
  icon?: string
  iconColor?: string
  size?: 'small' | 'medium' | 'large'
  actionText?: string
  actionIcon?: string
  actionColor?: string
  actionVariant?: 'text' | 'outlined' | 'flat' | 'elevated' | 'tonal' | 'plain'
  actionHandler?: () => void
}

const props = withDefaults(defineProps<EmptyStateProps>(), {
  icon: 'fas fa-inbox',
  iconColor: 'grey-darken-1',
  size: 'medium',
  actionColor: 'primary',
  actionVariant: 'elevated'
})

const iconSize = computed(() => {
  switch (props.size) {
    case 'small': return 48
    case 'large': return 80
    default: return 64
  }
})

const computedIcon = computed(() => props.icon)
</script>

<style scoped>
.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
  padding: 2rem;
  font-family: 'Inter', system-ui, sans-serif;
}

.empty-state-content {
  text-align: center;
  max-width: 400px;
}

.empty-icon {
  opacity: 0.6;
}

.empty-actions {
  display: flex;
  justify-content: center;
  gap: 12px;
  flex-wrap: wrap;
}
</style>
