<template>
  <div class="dashboard-page">
    <div class="mb-8">
      <h1 class="text-3xl font-weight-bold tracking-tight dashboard-title">Dashboard</h1>
    </div>

    <!-- KPI Cards Grid -->
    <div class="mb-8">
      <v-row>
        <v-col cols="12" sm="6" lg="3">
          <KpiCard
            title="Total Stocks"
            :value="summary?.total_stocks || 0"
            icon="fas fa-chart-bar"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <KpiCard
            title="Buy Recommendations"
            :value="`${summary?.buy_percentage?.toFixed(1) || 0}%`"
            icon="fas fa-percentage"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <KpiCard
            title="Avg. Confidence"
            :value="summary?.average_confidence?.toFixed(2) || 0"
            icon="fas fa-shield-check"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <KpiCard title="AI Analyst" value="Ready" icon="fas fa-robot" link="/ai-analyst" />
        </v-col>
      </v-row>
    </div>

    <!-- Charts and Top Brokerages -->
    <div>
      <v-row>
        <v-col cols="12" lg="8">
          <DistributionChart v-if="summary" :data="summary" />
        </v-col>
        <v-col cols="12" lg="4">
          <TopBrokerages v-if="topBrokerages" :data="topBrokerages" />
        </v-col>
      </v-row>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" />
    </div>

    <!-- Error State -->
    <v-alert v-if="error" type="error" class="mt-4" :text="error" dismissible />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAnalytics } from '@/stores/analytics'
import KpiCard from '@/components/dashboard/KpiCard.vue'
import DistributionChart from '@/components/dashboard/DistributionChart.vue'
import TopBrokerages from '@/components/dashboard/TopBrokerages.vue'
import type { AnalyticsSummary, BrokerageSummary } from '@/utils/types'

const analytics = useAnalytics()

const summary = ref<AnalyticsSummary | null>(null)
const topBrokerages = ref<BrokerageSummary[] | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

const loadDashboardData = async () => {
  try {
    loading.value = true
    error.value = null

    // Load analytics summary
    await analytics.fetchSummary()
    summary.value = analytics.summary

    // Load top brokerages
    await analytics.fetchTopBrokerages()
    topBrokerages.value = analytics.topBrokerages
  } catch (err: any) {
    error.value = err.message || 'Error loading dashboard data'
    console.error('Dashboard error:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadDashboardData()
})
</script>

<style scoped>
.dashboard-page {
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-title {
  font-family: 'Inter', system-ui, sans-serif;
  color: var(--v-theme-on-surface);
  letter-spacing: -0.025em;
}
</style>
