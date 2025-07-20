<template>
  <section class="analytics-page">
    <div class="mb-8">
      <h1 class="text-h4 font-weight-bold mb-2">Advanced Analytics</h1>
      <p class="text-body-2 text-medium-emphasis">
        Deep dive into market trends and recommendation data.
      </p>
    </div>

    <!-- KPI Cards -->
    <v-row class="mb-8">
      <v-col cols="12" sm="6" lg="3">
        <AnalyticsKPI
          title="Total Stocks Tracked"
          :value="store.summary?.total_stocks || 0"
          icon="fas fa-chart-bar"
        />
      </v-col>
      <v-col cols="12" sm="6" lg="3">
        <AnalyticsKPI
          title="Buy Recommendations"
          :value="store.summary?.buy_percentage || 0"
          icon="fas fa-arrow-trend-up"
          format="percentage"
          :change="2.4"
        />
      </v-col>
      <v-col cols="12" sm="6" lg="3">
        <AnalyticsKPI
          title="Sell Recommendations"
          :value="store.summary?.sell_percentage || 0"
          icon="fas fa-arrow-trend-down"
          format="percentage"
          :change="1.2"
        />
      </v-col>
      <v-col cols="12" sm="6" lg="3">
        <AnalyticsKPI
          title="Avg. Confidence"
          :value="(store.summary?.average_confidence || 0) * 100"
          icon="fas fa-shield-check"
          format="percentage"
          :change="-0.8"
        />
      </v-col>
    </v-row>

    <!-- Charts Section -->
    <v-row class="mb-8">
      <v-col cols="12" lg="8">
        <BarChart title="Quarterly Recommendation Trends" :data="store.quarterlyData" />
      </v-col>
      <v-col cols="12" lg="4">
        <LineChart title="Confidence & Volume Trends" :data="store.trends" />
      </v-col>
    </v-row>

    <!-- Additional Charts Row -->
    <v-row>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="pa-4">
            <span class="text-lg font-weight-bold">Top Performing Sectors</span>
          </v-card-title>
          <v-card-text class="pa-4">
            <v-list density="compact">
              <v-list-item v-for="(sector, index) in topSectors" :key="sector.name" class="px-0">
                <template #prepend>
                  <v-avatar size="32" :color="getSectorColor(index)" class="mr-3">
                    <span class="text-caption font-weight-bold">{{ index + 1 }}</span>
                  </v-avatar>
                </template>
                <v-list-item-title>{{ sector.name }}</v-list-item-title>
                <template #append>
                  <v-chip size="small" color="success" variant="tonal">
                    +{{ sector.growth }}%
                  </v-chip>
                </template>
              </v-list-item>
            </v-list>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="pa-4">
            <span class="text-lg font-weight-bold">Market Sentiment</span>
          </v-card-title>
          <v-card-text class="pa-4">
            <div class="mb-4">
              <div class="d-flex justify-space-between align-center mb-2">
                <span class="text-body-2">Bullish Sentiment</span>
                <span class="text-body-2 font-weight-bold">68%</span>
              </div>
              <v-progress-linear model-value="68" color="success" height="8" rounded />
            </div>
            <div class="mb-4">
              <div class="d-flex justify-space-between align-center mb-2">
                <span class="text-body-2">Bearish Sentiment</span>
                <span class="text-body-2 font-weight-bold">22%</span>
              </div>
              <v-progress-linear model-value="22" color="error" height="8" rounded />
            </div>
            <div>
              <div class="d-flex justify-space-between align-center mb-2">
                <span class="text-body-2">Neutral Sentiment</span>
                <span class="text-body-2 font-weight-bold">10%</span>
              </div>
              <v-progress-linear model-value="10" color="warning" height="8" rounded />
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Loading State -->
    <div v-if="store.loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" size="64" />
    </div>

    <!-- Error State -->
    <v-alert
      v-if="store.error"
      type="error"
      class="mt-4"
      :text="store.error"
      dismissible
      @click:close="store.error = null"
    />
  </section>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAnalytics } from '@/stores/analytics'
import AnalyticsKPI from '@/components/analytics/AnalyticsKPI.vue'
import BarChart from '@/components/analytics/BarChart.vue'
import LineChart from '@/components/analytics/LineChart.vue'

const store = useAnalytics()

// Mock data for additional sections - replace with real API data
const topSectors = [
  { name: 'Technology', growth: 12.4 },
  { name: 'Healthcare', growth: 8.7 },
  { name: 'Financial Services', growth: 6.2 },
  { name: 'Consumer Goods', growth: 4.1 },
  { name: 'Energy', growth: 3.8 },
]

function getSectorColor(index: number): string {
  const colors = ['primary', 'secondary', 'success', 'warning', 'info']
  return colors[index] || 'primary'
}

onMounted(() => {
  store.fetchAllAnalytics()
})
</script>

<style scoped>
.analytics-page {
  max-width: 1200px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
