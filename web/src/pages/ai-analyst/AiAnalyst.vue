<template>
  <section class="ai-analyst-page">
    <div class="mb-8">
      <h1 class="text-h4 font-weight-bold mb-2">AI Financial Analyst</h1>
      <p class="text-medium-emphasis">
        Assess a stock's potential based on recent recommendations and market data.
      </p>
    </div>

    <v-row>
      <v-col cols="12" md="6">
        <AnalysisForm @analysis-submitted="handleAnalysisSubmit" />
      </v-col>

      <v-col cols="12" md="6">
        <AnalysisResult :loading="store.loading" :analysis="store.currentAnalysis" />
      </v-col>
    </v-row>

    <!-- Analysis History -->
    <div v-if="store.analysisHistory.length > 0" class="mt-12">
      <div class="d-flex justify-space-between align-center mb-6">
        <h2 class="text-h5 font-weight-bold">Recent Analyses</h2>
        <v-btn variant="text" size="small" @click="store.clearHistory" prepend-icon="fas fa-trash">
          Clear History
        </v-btn>
      </div>

      <v-row>
        <v-col
          v-for="analysis in store.analysisHistory.slice(0, 6)"
          :key="`${analysis.ticker}-${analysis.timestamp}`"
          cols="12"
          sm="6"
          lg="4"
        >
          <v-card
            variant="outlined"
            class="history-card"
            @click="loadHistoricalAnalysis(analysis)"
            :class="{ active: store.currentAnalysis?.timestamp === analysis.timestamp }"
          >
            <v-card-text class="pa-4">
              <div class="d-flex justify-space-between align-center mb-2">
                <div class="text-h6 font-weight-bold">{{ analysis.ticker }}</div>
                <v-chip
                  v-if="analysis.recommendation"
                  :color="getRecommendationColor(analysis.recommendation)"
                  variant="tonal"
                  size="small"
                >
                  {{ analysis.recommendation }}
                </v-chip>
              </div>
              <div class="text-body-2 text-medium-emphasis mb-3">
                {{ analysis.company }}
              </div>
              <div class="text-caption text-medium-emphasis">
                {{ formatHistoryDate(analysis.timestamp) }}
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- Error State -->
    <v-snackbar v-model="showError" color="error" timeout="5000" location="top">
      {{ store.error }}
      <template #actions>
        <v-btn variant="text" @click="showError = false" icon="fas fa-times" />
      </template>
    </v-snackbar>
  </section>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { format } from 'date-fns'
import { useAIAnalyst } from '@/stores/aiAnalyst'
import type { AIAnalysisResult } from '@/utils/types'
import AnalysisForm from '@/components/ai-analyst/AnalysisForm.vue'
import AnalysisResult from '@/components/ai-analyst/AnalysisResult.vue'
import type { AIAnalysisFormData } from '@/utils/validation'

const router = useRouter()
const route = useRoute()
const store = useAIAnalyst()

const showError = ref(false)

const handleAnalysisSubmit = async (formData: AIAnalysisFormData) => {
  try {
    // Update URL with current analysis params
    router.replace({
      query: {
        ...route.query,
        ticker: formData.ticker,
        company: formData.company,
      },
    })

    await store.analyzeStock({
      ticker: formData.ticker,
      company: formData.company,
    })
  } catch (error) {
    showError.value = true
  }
}

function loadHistoricalAnalysis(analysis: AIAnalysisResult) {
  store.currentAnalysis = analysis

  // Update URL
  router.replace({
    query: {
      ticker: analysis.ticker,
      company: analysis.company,
    },
  })
}

function getRecommendationColor(recommendation: string) {
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

function formatHistoryDate(timestamp: string): string {
  return format(new Date(timestamp), 'MMM d, yyyy HH:mm')
}

// Watch for errors
watch(
  () => store.error,
  (error) => {
    if (error) {
      showError.value = true
    }
  },
)
</script>

<style scoped>
.ai-analyst-page {
  max-width: 1200px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}

.history-card {
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}

.history-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(var(--v-theme-shadow-rgb), 0.15);
}

.history-card.active {
  border-color: rgb(var(--v-theme-primary));
  background: rgba(var(--v-theme-primary-rgb), 0.05);
}
</style>
