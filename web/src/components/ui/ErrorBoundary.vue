<template>
  <div v-if="!hasError">
    <slot />
  </div>
  <div v-else class="error-boundary">
    <v-alert
      type="error"
      variant="tonal"
      class="error-alert"
      :title="errorTitle"
      :text="errorMessage"
    >
      <template #prepend>
        <v-icon size="24">fas fa-exclamation-triangle</v-icon>
      </template>
      <template #append>
        <div class="error-actions">
          <v-btn variant="text" size="small" @click="showDetails = !showDetails">
            {{ showDetails ? 'Hide' : 'Show' }} Details
          </v-btn>
          <v-btn variant="outlined" size="small" color="primary" @click="handleRetry">
            <v-icon start size="16">fas fa-redo</v-icon>
            Try Again
          </v-btn>
        </div>
      </template>
    </v-alert>

    <v-expand-transition>
      <v-card v-if="showDetails" class="mt-4" variant="outlined">
        <v-card-title class="text-h6">Error Details</v-card-title>
        <v-card-text>
          <div class="error-details">
            <div class="mb-3"><strong>Error:</strong> {{ errorInfo.message }}</div>
            <div class="mb-3"><strong>Component:</strong> {{ errorInfo.component }}</div>
            <div class="mb-3"><strong>Time:</strong> {{ errorInfo.timestamp }}</div>
            <div v-if="errorInfo.stack">
              <strong>Stack Trace:</strong>
              <pre class="error-stack">{{ errorInfo.stack }}</pre>
            </div>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-btn variant="text" size="small" @click="copyErrorDetails" prepend-icon="fas fa-copy">
            Copy Error Info
          </v-btn>
          <v-btn
            v-if="reportHandler"
            variant="text"
            size="small"
            @click="reportHandler(errorInfo)"
            prepend-icon="fas fa-bug"
          >
            Report Issue
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-expand-transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onErrorCaptured, computed } from 'vue'
import { format } from 'date-fns'

interface ErrorBoundaryProps {
  fallbackTitle?: string
  fallbackMessage?: string
  reportHandler?: (error: ErrorInfo) => void
  onRetry?: () => void
}

interface ErrorInfo {
  message: string
  component: string
  stack?: string
  timestamp: string
}

const props = withDefaults(defineProps<ErrorBoundaryProps>(), {
  fallbackTitle: 'Something went wrong',
  fallbackMessage:
    'An unexpected error occurred. Please try again or contact support if the problem persists.',
})

const emit = defineEmits<{
  'error-captured': [error: Error, info: ErrorInfo]
}>()

const hasError = ref(false)
const showDetails = ref(false)
const errorInfo = ref<ErrorInfo>({
  message: '',
  component: '',
  timestamp: '',
})

const errorTitle = computed(() => props.fallbackTitle)
const errorMessage = computed(() => props.fallbackMessage)

onErrorCaptured((error: Error, instance) => {
  console.error('Error captured by ErrorBoundary:', error)

  hasError.value = true
  errorInfo.value = {
    message: error.message,
    component: instance?.type?.name || 'Unknown Component',
    stack: error.stack,
    timestamp: format(new Date(), 'PPpp'),
  }

  emit('error-captured', error, errorInfo.value)

  // Return false to prevent the error from propagating further
  return false
})

function handleRetry() {
  hasError.value = false
  showDetails.value = false
  errorInfo.value = {
    message: '',
    component: '',
    timestamp: '',
  }

  if (props.onRetry) {
    props.onRetry()
  }
}

async function copyErrorDetails() {
  const details = `
Error: ${errorInfo.value.message}
Component: ${errorInfo.value.component}
Time: ${errorInfo.value.timestamp}
Stack: ${errorInfo.value.stack || 'No stack trace available'}
  `.trim()

  try {
    await navigator.clipboard.writeText(details)
    // You could show a toast here
  } catch (err) {
    console.error('Failed to copy error details:', err)
  }
}
</script>

<style scoped>
.error-boundary {
  margin: 1rem 0;
  font-family: 'Inter', system-ui, sans-serif;
}

.error-alert {
  margin-bottom: 0;
}

.error-actions {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.error-details {
  font-size: 0.875rem;
}

.error-stack {
  background: rgb(var(--v-theme-surface-variant));
  border-radius: 4px;
  padding: 0.75rem;
  font-size: 0.8rem;
  overflow-x: auto;
  white-space: pre-wrap;
  word-break: break-all;
  max-height: 200px;
  overflow-y: auto;
}

@media (max-width: 768px) {
  .error-actions {
    flex-direction: column;
    width: 100%;
  }
}
</style>
