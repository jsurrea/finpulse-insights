import { defineStore } from 'pinia'
import type { AIAnalysisRequest, AIAnalysisResult, StockOption } from '@/utils/types'
import { getStocks } from '@/utils/api'

interface AIAnalystState {
  currentAnalysis: AIAnalysisResult | null
  analysisHistory: AIAnalysisResult[]
  stockOptions: StockOption[]
  loading: boolean
  loadingSearch: boolean
  error: string | null
  searchQuery: string
  debounceTimer: number | null
}

export const useAIAnalyst = defineStore('aiAnalyst', {
  state: (): AIAnalystState => ({
    currentAnalysis: null,
    analysisHistory: [],
    stockOptions: [],
    loading: false,
    loadingSearch: false,
    error: null,
    searchQuery: '',
    debounceTimer: null
  }),

  getters: {
    hasAnalysis: (state) => !!state.currentAnalysis,
    lastAnalysisTime: (state) => {
      if (!state.currentAnalysis) return null
      return new Date(state.currentAnalysis.timestamp)
    },
    isSearching: (state) => state.loadingSearch
  },

  actions: {
    // Buscar stocks con debouncing
    searchStocks(query: string) {
      this.searchQuery = query

      // Limpiar el timer anterior
      if (this.debounceTimer) {
        clearTimeout(this.debounceTimer)
      }

      // Establecer nuevo timer con debounce de 100ms
      this.debounceTimer = window.setTimeout(async () => {
        await this.fetchStockOptions(query)
      }, 100)
    },

    // Cargar opciones de stocks desde la API
    async fetchStockOptions(search = '') {
      this.loadingSearch = true
      try {
        const response = await getStocks(1, 30, search) // Límite de 30 para el autocomplete

        this.stockOptions = response.data?.map((stock) => ({
          ticker: stock.ticker,
          company: stock.company,
          brokerage: stock.brokerage,
          id: stock.ticker
        })) || []
      } catch (error) {
        console.error('Failed to load stock options:', error)
        this.stockOptions = []
      } finally {
        this.loadingSearch = false
      }
    },

    // Cargar opciones iniciales (sin búsqueda)
    async loadStockOptions() {
      await this.fetchStockOptions('')
    },

    async analyzeStock(request: AIAnalysisRequest): Promise<AIAnalysisResult> {
      this.loading = true
      this.error = null

      try {
        const n8nWebhookUrl = import.meta.env.VITE_N8N_WEBHOOK_URL
        const n8nSecret = import.meta.env.VITE_N8N_SECRET

        if (!n8nWebhookUrl) {
          throw new Error('N8N webhook URL not configured')
        }

        // Encontrar la información completa del stock
        const stockInfo = this.stockOptions.find(s => s.ticker === request.ticker)

        if (!stockInfo) {
          throw new Error(`Stock ${request.ticker} not found`)
        }

        // Llamada directa al webhook de n8n
        const response = await fetch(n8nWebhookUrl, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            ...(n8nSecret && { 'Authorization': `Bearer ${n8nSecret}` })
          },
          body: JSON.stringify({
            ticker: stockInfo.ticker,
            company: stockInfo.company,
            brokerage: stockInfo.brokerage,
            timestamp: new Date().toISOString()
          })
        })

        if (!response.ok) {
          throw new Error(`Analysis failed: ${response.statusText}`)
        }

        const result: AIAnalysisResult = await response.json()

        // Validar estructura de respuesta
        if (!result.ticker || !result.summary || !result.recommendation) {
          throw new Error('Invalid response format from AI analysis')
        }

        this.currentAnalysis = result

        // Agregar al historial (mantener últimos 10)
        this.analysisHistory.unshift(result)
        if (this.analysisHistory.length > 10) {
          this.analysisHistory = this.analysisHistory.slice(0, 10)
        }

        // Guardar en localStorage para persistencia
        localStorage.setItem('finpulse_ai_history', JSON.stringify(this.analysisHistory))

        return result
      } catch (err: unknown) {
        if (err instanceof Error) {
          this.error = err.message
        } else {
          this.error = 'Failed to analyze stock'
        }
        throw err
      } finally {
        this.loading = false
      }
    },

    clearCurrentAnalysis() {
      this.currentAnalysis = null
      this.error = null
    },

    clearHistory() {
      this.analysisHistory = []
      localStorage.removeItem('finpulse_ai_history')
    },

    loadHistoryFromStorage() {
      try {
        const stored = localStorage.getItem('finpulse_ai_history')
        if (stored) {
          this.analysisHistory = JSON.parse(stored)
        }
      } catch (error) {
        console.error('Failed to load analysis history:', error)
      }
    },

    // Limpiar el timer de debounce
    clearDebounceTimer() {
      if (this.debounceTimer) {
        clearTimeout(this.debounceTimer)
        this.debounceTimer = null
      }
    }
  }
})
