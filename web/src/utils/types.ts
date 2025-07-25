export interface Pagination {
  page: number
  limit: number
  total: number
}

export interface Stock {
  ticker: string
  company: string
  brokerage: string
}

export interface StockRecommendation {
  action: string
  rating_from: string
  rating_to: string
  target_from?: number
  target_to?: number
  time: string
}

export interface StockDetail {
  ticker: string
  company: string
  brokerage: string
  latest_recommendation: StockRecommendation
  history: Recommendation[]
}

export interface Recommendation {
  id: string
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from?: number
  target_to?: number
  time: string
  recommendation: 'BUY' | 'SELL' | 'HOLD'
  confidence: number
  reason: string
}

export interface AnalyticsSummary {
  total_stocks: number
  total_recommendations: number
  buy_percentage: number
  sell_percentage: number
  hold_percentage: number
  average_confidence: number
}

export interface HealthStatus {
  status: 'ok' | 'error' | 'degraded'
  database: 'connected' | 'disconnected'
  uptime: number
}

export interface BrokerageSummary {
  name: string
  stockCount: number
}

export interface StockOption {
  ticker: string
  company: string
  brokerage: string
  id: string
}

export interface AIAnalysisRequest {
  ticker: string
}

export interface AIAnalysisResult {
  ticker: string
  company: string
  brokerage: string
  summary: string
  recommendation: 'BUY' | 'SELL' | 'HOLD'
  confidence_score: number
  key_factors: string[]
  risk_assessment: string
  news_summary: string
  price_trend: string
  timestamp: string
}

export interface AIAnalysisError {
  message: string
  code?: string
}
