import { defineStore } from 'pinia'
import type { AnalyticsSummary } from '@/utils/types'
import { getAnalyticsSummary, getAnalyticsTrends, getTopBrokerages } from '@/utils/api'

interface TrendData {
  date: string
  confidence: number
  volume: number
}

interface QuarterlyData {
  name: string
  buy: number
  sell: number
  hold: number
}

interface AnalyticsState {
  summary: AnalyticsSummary | null
  trends: TrendData[]
  quarterlyData: QuarterlyData[]
  topBrokerages: Array<{ name: string; stockCount: number }>
  loading: boolean
  error: string | null
}

export const useAnalytics = defineStore('analytics', {
  state: (): AnalyticsState => ({
    summary: null,
    trends: [],
    quarterlyData: [],
    topBrokerages: [],
    loading: false,
    error: null
  }),
  actions: {
    async fetchSummary() {
      this.loading = true
      this.error = null
      try {
        this.summary = await getAnalyticsSummary()
      } catch (err: any) {
        this.error = err.message || 'Error fetching analytics'
      } finally {
        this.loading = false
      }
    },
    async fetchTopBrokerages() {
      this.loading = true
      this.error = null
      try {
        this.topBrokerages = await getTopBrokerages()
      } catch (err: any) {
        this.error = err.message || 'Error fetching top brokerages'
      } finally {
        this.loading = false
      }
    },
    async fetchTrends() {
      try {
        const response = await getAnalyticsTrends()
        this.trends = response
      } catch (err: any) {
        console.error('Error fetching trends:', err)
        this.trends = []
      }
    },
    async fetchQuarterlyData() {
      try {
        // This is a placeholder
        this.quarterlyData = [
          { name: 'Q1 2024', buy: 450, sell: 120, hold: 230 },
          { name: 'Q2 2024', buy: 520, sell: 95, hold: 285 },
          { name: 'Q3 2024', buy: 380, sell: 180, hold: 340 },
          { name: 'Q4 2024', buy: 610, sell: 75, hold: 215 }
        ]
      } catch (err: any) {
        console.error('Error fetching quarterly data:', err)
        this.quarterlyData = []
      }
    },
    async fetchAllAnalytics() {
      await Promise.all([
        this.fetchSummary(),
        this.fetchTrends(),
        this.fetchQuarterlyData()
      ])
    },
  },
})
