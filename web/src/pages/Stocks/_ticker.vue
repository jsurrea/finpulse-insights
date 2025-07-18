<template>
  <div class="stock-detail-page">
    <div class="mb-4">
      <v-btn
        variant="outlined"
        :to="'/stocks'"
        prepend-icon="fas fa-arrow-left"
      >
        Back to Stocks
      </v-btn>
    </div>

    <div class="mb-8">
      <h1 class="text-3xl font-weight-bold tracking-tight stock-title">
        {{ stock?.company }}
      </h1>
      <p class="text-xl text-medium-emphasis font-mono">
        {{ stock?.ticker }}
      </p>
    </div>

    <v-row v-if="stock">
      <!-- Left Column - Key Info & Latest Action -->
      <v-col cols="12" md="4">
        <div class="mb-6">
          <v-card elevation="2">
            <v-card-title>Key Info</v-card-title>
            <v-card-text>
              <div class="info-item">
                <div class="d-flex align-center text-medium-emphasis">
                  <v-icon class="mr-2">fas fa-building</v-icon>
                  Brokerage
                </div>
                <div class="font-weight-semibold">{{ stock.brokerage }}</div>
              </div>
              <div class="info-item">
                <div class="d-flex align-center text-medium-emphasis">
                  <v-icon class="mr-2">fas fa-calendar</v-icon>
                  Last Update
                </div>
                <div class="font-weight-semibold">
                  {{ formatDate(stock.latest_recommendation.time) }}
                </div>
              </div>
            </v-card-text>
          </v-card>
        </div>

        <v-card elevation="2">
          <v-card-title>Latest Action</v-card-title>
          <v-card-text class="text-center">
            <p class="text-lg font-weight-semibold text-capitalize mb-4">
              {{ stock.latest_recommendation.action.toLowerCase() }}
            </p>
            <div class="d-flex justify-center align-center ga-4">
              <v-chip variant="outlined" size="small">
                {{ stock.latest_recommendation.rating_from }}
              </v-chip>
              <v-icon class="text-medium-emphasis">fas fa-arrow-right</v-icon>
              <v-chip color="primary" size="small">
                {{ stock.latest_recommendation.rating_to }}
              </v-chip>
            </div>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Right Column - Recommendation History -->
      <v-col cols="12" md="8">
        <v-card elevation="2">
          <v-card-title>Recommendation History</v-card-title>
          <v-card-subtitle>
            Recent analyst actions for {{ stock.ticker }}
          </v-card-subtitle>
          <v-card-text>
            <v-table density="compact">
              <thead>
                <tr>
                  <th class="text-left font-weight-bold">Date</th>
                  <th class="text-left font-weight-bold">Brokerage</th>
                  <th class="text-left font-weight-bold">Action</th>
                  <th class="text-left font-weight-bold">Recommendation</th>
                  <th class="text-center" style="width: 50px;"></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="rec in stock.history" :key="rec.id">
                  <td>{{ formatDate(rec.time, 'yyyy-MM-dd') }}</td>
                  <td>{{ rec.brokerage }}</td>
                  <td class="text-capitalize">{{ rec.action.toLowerCase() }}</td>
                  <td>
                    <v-chip
                      :color="getRecommendationColor(rec.recommendation)"
                      size="small"
                      class="font-weight-semibold"
                    >
                      {{ rec.recommendation }}
                    </v-chip>
                  </td>
                  <td class="text-center">
                    <v-btn
                      icon
                      variant="text"
                      size="small"
                      color="primary"
                      :to="`/recommendations/${rec.id}`"
                    >
                      <v-icon>fas fa-arrow-right</v-icon>
                    </v-btn>
                  </td>
                </tr>
              </tbody>
            </v-table>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- AI Analysis Button -->
    <div class="mt-8" v-if="stock">
      <v-btn
        color="primary"
        size="large"
        :to="`/ai-analyst?ticker=${stock.ticker}&company=${encodeURIComponent(stock.company)}`"
        append-icon="fas fa-arrow-right"
      >
        Run AI Analysis
      </v-btn>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" />
    </div>

    <!-- Error State -->
    <v-alert
      v-if="error"
      type="error"
      class="mt-4"
      :text="error"
      dismissible
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useStocks } from '@/stores/stocks'
import { format } from 'date-fns'
import type { StockDetail } from '@/utils/types'

const route = useRoute()
const stocksStore = useStocks()

const stock = ref<StockDetail | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

const ticker = route.params.ticker as string

const loadStockDetail = async () => {
  try {
    loading.value = true
    error.value = null

    await stocksStore.fetchStockDetail(ticker)
    stock.value = stocksStore.currentStock

  } catch (err: any) {
    error.value = err.message || 'Error loading stock detail'
    console.error('Stock detail error:', err)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateString: string, formatStr: string = 'MMM d, yyyy') => {
  return format(new Date(dateString), formatStr)
}

const getRecommendationColor = (recommendation: string) => {
  switch (recommendation) {
    case 'BUY':
      return 'success'
    case 'SELL':
      return 'error'
    case 'HOLD':
      return 'warning'
    default:
      return 'primary'
  }
}

onMounted(() => {
  loadStockDetail()
})
</script>

<style scoped>
.stock-detail-page {
  max-width: 1200px;
  margin: 0 auto;
}

.stock-title {
  font-family: 'Inter', system-ui, sans-serif;
  letter-spacing: -0.025em;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid rgba(var(--v-theme-on-surface), 0.08);
}

.info-item:last-child {
  border-bottom: none;
}

.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
