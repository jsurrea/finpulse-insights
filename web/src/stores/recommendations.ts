import { defineStore } from 'pinia'
import type { Recommendation, Pagination } from '@/utils/types'
import { getRecommendations, getRecommendationDetail } from '@/utils/api'

interface RecommendationsState {
  recommendations: Recommendation[]
  currentRec: Recommendation | null
  loading: boolean
  error: string | null
  filters: {
    page: number
    limit: number
    ticker?: string
    company?: string
    brokerage?: string
    action?: string
    rating?: string
    date_from?: string
    date_to?: string
    sort?: string
  }
  pagination: Pagination | null
}

export const useRecommendations = defineStore('recommendations', {
  state: (): RecommendationsState => ({
    recommendations: [],
    currentRec: null,
    loading: false,
    error: null,
    filters: {
      page: 1,
      limit: 20,
      sort: 'time:desc'
    },
    pagination: null
  }),
  actions: {
    async fetchRecommendations() {
      this.loading = true
      this.error = null
      try {
        const { data, pagination } = await getRecommendations(
          this.filters.page,
          this.filters.limit,
          this.filters.ticker,
          this.filters.company,
          this.filters.brokerage,
          this.filters.action,
          this.filters.rating,
          this.filters.date_from,
          this.filters.date_to,
          this.filters.sort
        )
        this.recommendations = data
        this.pagination = pagination
      } catch (err: any) {
        this.error = err.message || 'Error fetching recommendations'
      } finally {
        this.loading = false
      }
    },
    async fetchRecommendationDetail(id: string) {
      this.loading = true
      this.error = null
      try {
        this.currentRec = await getRecommendationDetail(id)
      } catch (err: any) {
        this.error = err.message || 'Error fetching recommendation detail'
      } finally {
        this.loading = false
      }
    },
    applyFilters(filters: Partial<RecommendationsState['filters']>) {
      this.filters = { ...this.filters, ...filters }
      this.fetchRecommendations()
    },
    clearCurrent() {
      this.currentRec = null
    }
  }
})
