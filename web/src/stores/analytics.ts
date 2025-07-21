import { defineStore } from 'pinia'
import type { AnalyticsSummary } from '@/utils/types'
import { getAnalyticsSummary, getTopBrokerages } from '@/utils/api'

interface AnalyticsState {
  summary: AnalyticsSummary | null
  topBrokerages: Array<{ name: string; stockCount: number }>
  loading: boolean
  error: string | null
}

export const useAnalytics = defineStore('analytics', {
  state: (): AnalyticsState => ({
    summary: null,
    topBrokerages: [],
    loading: false,
    error: null,
  }),
  actions: {
    async fetchSummary() {
      this.loading = true
      this.error = null
      try {
        this.summary = await getAnalyticsSummary()
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message
        } else {
          this.error = 'Error fetching analytics'
        }
      } finally {
        this.loading = false
      }
    },
    async fetchTopBrokerages() {
      this.loading = true
      this.error = null
      try {
        this.topBrokerages = await getTopBrokerages()
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message
        } else {
          this.error = 'Error fetching top brokerages'
        }
      } finally {
        this.loading = false
      }
    },
  },
})
