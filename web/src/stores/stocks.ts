import { defineStore } from 'pinia'
import type { Stock, StockDetail, Pagination } from '@/utils/types'
import { getStocks, getStockDetail, getTopBrokerages } from '@/utils/api'

interface StocksState {
  stocks: Stock[]
  brokerages: string[]
  currentStock: StockDetail | null
  loading: boolean
  error: string | null
  filters: {
    page: number
    limit: number
    search: string
    brokerage: string
  }
  pagination: Pagination | null
}

export const useStocks = defineStore('stocks', {
  state: (): StocksState => ({
    stocks: [],
    brokerages: [],
    currentStock: null,
    loading: false,
    error: null,
    filters: {
      page: 1,
      limit: 20,
      search: '',
      brokerage: 'all',
    },
    pagination: null,
  }),
  actions: {
    async fetchStocks() {
      this.loading = true
      this.error = null
      try {
        const { data, pagination } = await getStocks(
          this.filters.page,
          this.filters.limit,
          this.filters.search,
          this.filters.brokerage,
        )
        this.stocks = data
        this.pagination = pagination
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message || 'Error fetching stocks'
        } else {
          this.error = 'Error fetching stocks'
        }
      } finally {
        this.loading = false
      }
    },
    async fetchStockDetail(ticker: string) {
      this.loading = true
      this.error = null
      try {
        this.currentStock = await getStockDetail(ticker)
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message || 'Error fetching stock detail'
        } else {
          this.error = 'Error fetching stock detail'
        }
      } finally {
        this.loading = false
      }
    },
    async fetchBrokerages() {
      try {
        const arr = await getTopBrokerages()
        this.brokerages = arr.map((b) => b.name)
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message || 'Error fetching brokerages'
        } else {
          this.error = 'Error fetching brokerages'
        }
        console.error('Brokerages loading error:', err)
        this.brokerages = []
      }
    },
    applyFilters(filters: Partial<StocksState['filters']>) {
      this.filters = { ...this.filters, ...filters }
      this.fetchStocks()
    },
    clearCurrentStock() {
      this.currentStock = null
    },
  },
})
