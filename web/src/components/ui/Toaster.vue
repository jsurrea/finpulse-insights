<template>
  <teleport to="body">
    <div class="toaster-container">
      <transition-group name="toast" tag="div">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          :class="toastClasses(toast)"
        >
          <v-icon
            :icon="getToastIcon(toast.type)"
            :color="toast.type"
            size="20"
            class="mr-3"
          />
          <div class="toast-content">
            <div v-if="toast.title" class="toast-title">{{ toast.title }}</div>
            <div class="toast-message">{{ toast.message }}</div>
          </div>
          <v-btn
            icon
            variant="text"
            size="small"
            @click="removeToast(toast.id)"
            class="ml-2"
          >
            <v-icon size="16">fas fa-times</v-icon>
          </v-btn>
        </div>
      </transition-group>
    </div>
  </teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Toast {
  id: string
  type: 'success' | 'error' | 'warning' | 'info'
  title?: string
  message: string
  duration?: number
}

const toasts = ref<Toast[]>([])
let toastCounter = 0

function addToast(toast: Omit<Toast, 'id'>) {
  const id = `toast-${++toastCounter}`
  const newToast: Toast = { id, ...toast }

  toasts.value.push(newToast)

  if (toast.duration !== 0) {
    setTimeout(() => {
      removeToast(id)
    }, toast.duration || 5000)
  }

  return id
}

function removeToast(id: string) {
  const index = toasts.value.findIndex(t => t.id === id)
  if (index > -1) {
    toasts.value.splice(index, 1)
  }
}

function clearAll() {
  toasts.value = []
}

function toastClasses(toast: Toast) {
  return [
    'toast',
    `toast-${toast.type}`
  ]
}

function getToastIcon(type: Toast['type']) {
  switch (type) {
    case 'success': return 'fas fa-check-circle'
    case 'error': return 'fas fa-times-circle'
    case 'warning': return 'fas fa-exclamation-triangle'
    case 'info': return 'fas fa-info-circle'
    default: return 'fas fa-info-circle'
  }
}

// Expose methods for external use
defineExpose({
  addToast,
  removeToast,
  clearAll,
  success: (message: string, title?: string) => addToast({ type: 'success', message, title }),
  error: (message: string, title?: string) => addToast({ type: 'error', message, title }),
  warning: (message: string, title?: string) => addToast({ type: 'warning', message, title }),
  info: (message: string, title?: string) => addToast({ type: 'info', message, title })
})
</script>

<style scoped>
.toaster-container {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 10000;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  max-width: 400px;
  font-family: 'Inter', system-ui, sans-serif;
}

.toast {
  display: flex;
  align-items: flex-start;
  padding: 1rem;
  background: rgb(var(--v-theme-surface));
  border: 1px solid rgb(var(--v-theme-outline));
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  backdrop-filter: blur(8px);
}

.toast-success {
  border-left: 4px solid rgb(var(--v-theme-success));
}

.toast-error {
  border-left: 4px solid rgb(var(--v-theme-error));
}

.toast-warning {
  border-left: 4px solid rgb(var(--v-theme-warning));
}

.toast-info {
  border-left: 4px solid rgb(var(--v-theme-info));
}

.toast-content {
  flex: 1;
}

.toast-title {
  font-weight: 600;
  margin-bottom: 0.25rem;
  color: rgb(var(--v-theme-on-surface));
}

.toast-message {
  color: rgb(var(--v-theme-on-surface-variant));
  font-size: 0.875rem;
  line-height: 1.4;
}

/* Transitions */
.toast-enter-active {
  transition: all 0.3s ease-out;
}

.toast-leave-active {
  transition: all 0.3s ease-in;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

@media (max-width: 768px) {
  .toaster-container {
    left: 1rem;
    right: 1rem;
    max-width: none;
  }
}
</style>
