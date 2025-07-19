import { defineStore } from 'pinia'
import type { AIAnalysisRequest, AIAnalysisResult } from '@/utils/types'
import { analyzeStockPotential } from '@/utils/api'

interface AIAnalystState {
  currentAnalysis: AIAnalysisResult | null
  analysisHistory: AIAnalysisResult[]
  loading: boolean
  error: string | null
}

export const useAIAnalyst = defineStore('aiAnalyst', {
  state: (): AIAnalystState => ({
    currentAnalysis: null,
    analysisHistory: [],
    loading: false,
    error: null
  }),

  getters: {
    hasAnalysis: (state) => !!state.currentAnalysis,
    lastAnalysisTime: (state) => {
      if (!state.currentAnalysis) return null
      return new Date(state.currentAnalysis.timestamp)
    }
  },

  actions: {
    async analyzeStock(request: AIAnalysisRequest): Promise<AIAnalysisResult> {
      this.loading = true
      this.error = null

      try {
        const result = await analyzeStockPotential(request)

        this.currentAnalysis = result

        // Add to history (keep last 10 analyses)
        this.analysisHistory.unshift(result)
        if (this.analysisHistory.length > 10) {
          this.analysisHistory = this.analysisHistory.slice(0, 10)
        }

        return result
      } catch (err: any) {
        this.error = err.message || 'Failed to analyze stock'
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
    }
  }
})
