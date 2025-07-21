<template>
  <div class="analysis-result">
    <!-- Loading State - Sin cambios -->
    <div v-if="loading" class="loading-state">
      <div class="d-flex flex-column align-center ga-6">
        <v-progress-circular
          indeterminate
          size="56"
          color="primary"
          width="3"
        />
        <div class="text-center">
          <div class="text-h6 font-weight-medium mb-2">Generating insights...</div>
          <div class="text-body-2 text-disabled">
            Analyzing market data and recommendations
          </div>
        </div>
      </div>
    </div>

    <!-- Analysis Result -->
    <v-card
      v-else-if="analysis"
      class="analysis-card elevation-2"
      :class="analysisCardClasses"
    >
      <!-- Header con confidence score agregado -->
      <v-card-title class="d-flex align-center pa-6 pb-4">
        <div class="d-flex align-center ga-3 flex-grow-1">
          <v-avatar color="primary" size="40" class="analysis-avatar">
            <v-icon color="white" size="20" style="position: relative; left: -2px;">fas fa-robot</v-icon>
          </v-avatar>
          <div>
            <div class="text-h6 font-weight-bold">{{ analysis.ticker }}</div>
            <div class="text-body-2 text-disabled">{{ analysis.company }}</div>
          </div>
        </div>

        <div class="d-flex align-center ga-2">
          <!-- Confidence score movido aquí -->
          <v-chip
            v-if="analysis.confidence_score"
            :color="confidenceColor"
            size="small"
            variant="tonal"
            class="font-weight-bold"
          >
            {{ (analysis.confidence_score * 100).toFixed(0) }}% Confidence
          </v-chip>

          <v-chip
            v-if="analysis.recommendation"
            :color="recommendationColor"
            variant="elevated"
            size="large"
            class="font-weight-bold recommendation-chip"
          >
            {{ analysis.recommendation }}
          </v-chip>
        </div>
      </v-card-title>

      <v-card-text class="pa-6 pt-0">
        <!-- Summary Section - Sin cambios -->
        <div class="analysis-summary">
          <div class="summary-header mb-3">
            <v-icon color="primary" size="18" class="mr-2">fas fa-file-alt</v-icon>
            <span class="text-subtitle-1 font-weight-medium">Executive Summary</span>
          </div>
          <div class="text-body-1 summary-text">
            {{ analysis.summary }}
          </div>
        </div>

        <!-- Key Factors - Cambiado a columna y mismo estilo que summary -->
        <div v-if="analysis.key_factors" class="key-factors-section">
          <div class="summary-header mb-3">
            <v-icon color="primary" size="18" class="mr-2">fas fa-tags</v-icon>
            <span class="text-subtitle-1 font-weight-medium">Key Factors</span>
          </div>
          <div class="d-flex flex-wrap ga-2">
            <v-chip
              v-for="factor in analysis.key_factors"
              :key="factor"
              size="default"
              variant="outlined"
              color="primary"
              class="factor-chip"
            >
              {{ factor }}
            </v-chip>
          </div>
        </div>

        <!-- Recent News - Cambiado a columna y mismo estilo -->
        <div v-if="analysis.news_summary" class="news-section">
          <div class="summary-header mb-3">
            <v-icon color="primary" size="18" class="mr-2">fas fa-newspaper</v-icon>
            <span class="text-subtitle-1 font-weight-medium">Recent News</span>
          </div>
          <div class="text-body-1 summary-text">
            {{ analysis.news_summary }}
          </div>
        </div>

        <!-- Price Outlook - Cambiado a columna y mismo estilo -->
        <div v-if="analysis.price_trend" class="price-section">
          <div class="summary-header mb-3">
            <v-icon color="primary" size="18" class="mr-2">fas fa-chart-line</v-icon>
            <span class="text-subtitle-1 font-weight-medium">Price Outlook</span>
          </div>
          <div class="text-body-1 summary-text">
            {{ analysis.price_trend }}
          </div>
        </div>

        <!-- Risk Assessment - Sin cambios -->
        <div v-if="analysis.risk_assessment" class="risk-section mb-6">
          <v-alert
            type="warning"
            variant="tonal"
            prominent
            border="start"
            class="risk-alert"
          >
            <template #prepend>
              <v-icon>fas fa-shield-exclamation</v-icon>
            </template>
            <v-alert-title class="text-body-1 font-weight-medium">Risk Assessment</v-alert-title>
            <div class="text-body-2 mt-2">{{ analysis.risk_assessment }}</div>
          </v-alert>
        </div>

        <!-- Footer - Sin cambios -->
        <div class="analysis-footer">
          <div class="d-flex align-center justify-center text-disabled">
            <v-icon size="12" class="mr-1">fas fa-clock</v-icon>
            <span class="text-caption">Generated {{ formatTimestamp(analysis.timestamp) }}</span>
          </div>
        </div>
      </v-card-text>
    </v-card>

    <!-- Empty State - Sin cambios -->
    <div v-else class="empty-state">
      <div class="text-center">
        <v-avatar size="80" color="primary" variant="tonal" class="mb-4">
          <v-icon size="40" color="primary">fas fa-robot</v-icon>
        </v-avatar>
        <div class="text-h6 font-weight-medium mb-2">Ready for AI Analysis</div>
        <div class="text-body-2 text-disabled max-width-300">
          Select a stock from the dropdown to generate comprehensive AI-powered insights
        </div>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { computed } from 'vue'
import { formatDistanceToNow } from 'date-fns'
import type { AIAnalysisResult } from '@/utils/types'

interface AnalysisResultProps {
  loading: boolean
  analysis: AIAnalysisResult | null
}

const props = defineProps<AnalysisResultProps>()

const analysisCardClasses = computed(() => [
  'analysis-result-card',
  {
    'analysis-positive': props.analysis?.recommendation === 'BUY',
    'analysis-negative': props.analysis?.recommendation === 'SELL',
    'analysis-neutral': props.analysis?.recommendation === 'HOLD',
  },
])

const recommendationColor = computed(() => {
  switch (props.analysis?.recommendation) {
    case 'BUY': return 'success'
    case 'SELL': return 'error'
    case 'HOLD': return 'warning'
    default: return 'primary'
  }
})

const confidenceColor = computed(() => {
  const score = props.analysis?.confidence_score || 0
  if (score >= 0.8) return '#219653'
  if (score >= 0.6) return '#497aff'
  if (score >= 0.4) return '#F2C94C'
  return '#D9534F'
})

function formatTimestamp(timestamp: string): string {
  return formatDistanceToNow(new Date(timestamp), { addSuffix: true })
}
</script>

<style scoped>
.analysis-result--wrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  min-height: 400px;
  font-family: 'Inter', system-ui, sans-serif;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 420px;
  gap: 20px;
}

.loading-message {
  text-align: center;
}

.analysis-card {
  width: 100%;
  max-width: 500px;
  margin: 0 auto;
  box-shadow: 0 2px 10px 0 rgba(60, 60, 100, 0.06);
  border-radius: 18px;
}

.header {
  padding-bottom: 0.5rem !important;
}

.chip-upper {
  min-width: 62px;
  justify-content: center;
  font-size: 1rem;
  letter-spacing: 0.03em;
}

.section {
  width: 100%;
  margin-bottom: 8px;
}

.section-title {
  display: flex;
  align-items: center;
  font-size: 1rem;
}

.summary-text {
  font-size: 1rem;
}

.confidence-score-box {
  background: #F5F7FF;
  border-radius: 6px;
  padding: 8px 14px 7px 14px;
  margin-right: 12px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 70px;
  min-height: 48px;
}

.keyfactors-group {
  flex: 1 1 0;
}

.keyfactors-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  width: 100%;
  margin-top: 2px;
}

.keyfactor-chip {
  font-size: 1rem;
  min-width: 0;
  max-width: 300px;
  text-align: left;
  white-space: normal;
  word-break: break-word;
}

.news-block,
.trend-block {
  background: #F5F7FF;
  border-radius: 10px;
  padding: 14px 20px 12px 16px;
  margin: 0 0 2px 0;
}

.news-text, .trend-text {
  color: #23263b;
  font-size: 1rem;
  letter-spacing: 0.01em;
  line-height: 1.5;
}

.riskalert {
  font-size: 1rem;
  border-radius: 12px;
  background: #FFF8F0;
  color: #D35400;
  box-shadow: none;
  margin-top: 8px;
}

.analysis-result-card {
  border-top: 4px solid transparent;
}
.analysis-positive {
  border-top-color: #27ae60 !important;
}
.analysis-negative {
  border-top-color: #d9534f !important;
}
.analysis-neutral {
  border-top-color: #f2c94c !important;
}

.empty-state {
  width: 100%;
  min-height: 350px;
  padding: 40px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #AAA;
}

.max-width-300 {
  max-width: 300px;
  margin: 0 auto;
}

/* Nuevos estilos que siguen el patrón del executive summary */
.analysis-summary,
.key-factors-section,
.news-section,
.price-section {
  background: rgba(var(--v-theme-primary-rgb), 0.03);
  border: 1px solid rgba(var(--v-theme-primary-rgb), 0.08);
  border-radius: 12px;
  padding: 10px;
}

/* Chips de factors más grandes */
.factor-chip {
  font-size: 0.875rem;
  height: auto;
  min-height: 32px;
  padding: 8px 12px;
  white-space: normal;
  word-wrap: break-word;
  transition: transform 0.1s ease;
}

.factor-chip:hover {
  transform: scale(1.05);
}


@media (max-width: 840px) {
  .analysis-card {
    max-width: 100vw;
    border-radius: 0;
    margin: 0;
  }

  .section, .news-block, .trend-block, .keyfactors-group {
    padding-left: 8px !important;
    padding-right: 8px !important;
  }
}
</style>
