import { defineStore } from 'pinia'
import type { HealthStatus } from '@/utils/types'
import { getHealth } from '@/utils/api'

interface HealthState {
  healthStatus: HealthStatus | null
  loading: boolean
  error: string | null
  lastUpdated: Date | null
}

export const useHealth = defineStore('health', {
  state: (): HealthState => ({
    healthStatus: null,
    loading: false,
    error: null,
    lastUpdated: null,
  }),

  getters: {
    isHealthy: (state) => state.healthStatus?.status === 'ok',
    isDatabaseConnected: (state) => state.healthStatus?.database === 'connected',
    systemStatus: (state) => {
      if (!state.healthStatus) return 'unknown'
      if (state.healthStatus.status === 'ok' && state.healthStatus.database === 'connected') {
        return 'healthy'
      }
      if (state.healthStatus.status === 'error' || state.healthStatus.database === 'disconnected') {
        return 'critical'
      }
      return 'degraded'
    },
  },

  actions: {
    async fetchHealth() {
      this.loading = true
      this.error = null
      try {
        this.healthStatus = await getHealth()
        this.lastUpdated = new Date()
      } catch (err: any) {
        this.error = err.message || 'Error fetching health data'
        console.error('Health check error:', err)
      } finally {
        this.loading = false
      }
    },

    startAutoRefresh(intervalMs: number = 30000) {
      // Auto-refresh every 30 seconds by default
      this.fetchHealth()
      return setInterval(() => {
        this.fetchHealth()
      }, intervalMs)
    },
  },
})
