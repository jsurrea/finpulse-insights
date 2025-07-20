export { default as LoadingSpinner } from './LoadingSpinner.vue'
export { default as EmptyState } from './EmptyState.vue'
export { default as Pagination } from './Pagination.vue'
export { default as ErrorBoundary } from './ErrorBoundary.vue'
export { default as Skeleton } from './Skeleton.vue'
export { default as Toaster } from './Toaster.vue'

// Types
export interface PaginationEmits {
  'page-change': [page: number]
  'page-size-change': [size: number]
}

export interface ToasterMethods {
  addToast: (toast: {
    type: 'success' | 'error' | 'warning' | 'info'
    title?: string
    message: string
    duration?: number
  }) => string
  removeToast: (id: string) => void
  clearAll: () => void
  success: (message: string, title?: string) => string
  error: (message: string, title?: string) => string
  warning: (message: string, title?: string) => string
  info: (message: string, title?: string) => string
}
