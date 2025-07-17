import { defineStore } from 'pinia'
import type { Stock, StockDetail, Pagination } from '@/utils/types'
import { getStocks, getStockDetail } from '@/utils/api'

interface StocksState {
  stocks: Stock[]
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
    currentStock: null,
    loading: false,
    error: null,
    filters: {
      page: 1,
      limit: 20,
      search: '',
      brokerage: 'all'
    },
    pagination: null
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
          this.filters.brokerage
        )
        this.stocks = data
        this.pagination = pagination
      } catch (err: any) {
        this.error = err.message || 'Error fetching stocks'
      } finally {
        this.loading = false
      }
    },
    async fetchStockDetail(ticker: string) {
      this.loading = true
      this.error = null
      try {
        this.currentStock = await getStockDetail(ticker)
      } catch (err: any) {
        this.error = err.message || 'Error fetching stock detail'
      } finally {
        this.loading = false
      }
    },
    applyFilters(filters: Partial<StocksState['filters']>) {
      this.filters = { ...this.filters, ...filters }
      this.fetchStocks()
    },
    clearCurrentStock() {
      this.currentStock = null
    }
  }
})
