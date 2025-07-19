import axios from 'axios';
import type {
  Stock,
  StockDetail,
  Recommendation,
  AnalyticsSummary,
  HealthStatus,
  Pagination,
  BrokerageSummary
} from './types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    Accept: 'application/json'
  }
});

// --- STOCKS ---

export async function getStocks(
  page = 1,
  limit = 20,
  search = '',
  brokerage = 'all'
): Promise<{ data: Stock[]; pagination: Pagination }> {
  const params: any = {
    page,
    limit
  };
  if (search) params.search = search;
  if (brokerage && brokerage !== 'all') params.brokerage = brokerage;
  const res = await api.get('/stocks', { params });
  return res.data;
}

export async function getStockDetail(ticker: string): Promise<StockDetail> {
  const res = await api.get(`/stocks/${ticker}`);
  return res.data;
}

// --- RECOMMENDATIONS ---

export async function getRecommendations(
  page = 1,
  limit = 20,
  ticker?: string,
  company?: string,
  brokerage?: string,
  action?: string,
  rating?: string,
  date_from?: string,
  date_to?: string,
  sort: string = 'time:desc'
): Promise<{ data: Recommendation[]; pagination: Pagination }> {
  const params: any = { page, limit, sort };
  if (ticker) params.ticker = ticker;
  if (company) params.company = company;
  if (brokerage) params.brokerage = brokerage;
  if (action) params.action = action;
  if (rating) params.rating = rating;
  if (date_from) params.date_from = date_from;
  if (date_to) params.date_to = date_to;

  const res = await api.get('/recommendations', { params });
  return res.data;
}

export async function getRecommendationDetail(id: string): Promise<Recommendation> {
  const res = await api.get(`/recommendations/${id}`);
  return res.data;
}

// --- ANALYTICS ---

export async function getAnalyticsSummary(): Promise<AnalyticsSummary> {
  const res = await api.get('/analytics/summary');
  return res.data;
}

export async function getAnalyticsTrends(): Promise<Array<{
  date: string
  confidence: number
  volume: number
}>> {
  const res = await api.get('/analytics/trends')
  return res.data
}

export async function getAnalyticsQuarterly(): Promise<Array<{
  name: string
  buy: number
  sell: number
  hold: number
}>> {
  const res = await api.get('/analytics/quarterly')
  return res.data
}

// --- HEALTH ---

export async function getHealth(): Promise<HealthStatus> {
  const res = await api.get('/health');
  return res.data;
}

// --- BROKERAGES ---

export async function getTopBrokerages(): Promise<BrokerageSummary[]> {
  const res = await api.get('/analytics/brokerages');
  return res.data as BrokerageSummary[];
}

// --- AI ANALYST ---

export async function analyzeStockPotential(ticker: string): Promise<any> {
  const res = await api.get(`/ai-analyst/analyze?ticker=${encodeURIComponent(ticker)}`)
  return res.data
}

