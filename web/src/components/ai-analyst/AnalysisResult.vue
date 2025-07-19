<template>
  <div class="analysis-result">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state">
      <div class="d-flex flex-column align-center ga-4">
        <v-progress-circular
          indeterminate
          size="48"
          color="primary"
          width="4"
        />
        <div class="text-center">
          <div class="text-h6 mb-1">Generating insights...</div>
          <div class="text-body-2 text-medium-emphasis">
            Analyzing market data and recommendations
          </div>
        </div>
      </div>
    </div>

    <!-- Analysis Result -->
    <v-card
      v-else-if="analysis"
      class="analysis-card"
      :class="analysisCardClasses"
    >
      <v-card-title class="d-flex align-center ga-2 pb-2">
        <v-icon color="primary" size="24">fas fa-sparkles</v-icon>
        <span class="text-lg font-weight-bold">
          Analysis for {{ analysis.ticker }}
        </span>
        <v-spacer />
        <v-chip
          v-if="analysis.recommendation"
          :color="recommendationColor"
          variant="tonal"
          size="small"
          class="font-weight-bold"
        >
          {{ analysis.recommendation }}
        </v-chip>
      </v-card-title>

      <v-card-subtitle v-if="analysis.company" class="pb-4">
        {{ analysis.company }}
      </v-card-subtitle>

      <v-card-text>
        <!-- Summary -->
        <div class="analysis-summary mb-6">
          <div class="text-body-1 leading-relaxed">
            {{ analysis.summary }}
          </div>
        </div>

        <!-- Key Metrics -->
        <div v-if="analysis.confidence_score || analysis.key_factors" class="analysis-metrics">
          <v-row>
            <v-col v-if="analysis.confidence_score" cols="12" sm="6">
              <div class="metric-card">
                <div class="d-flex justify-space-between align-center mb-2">
                  <span class="text-body-2 font-weight-medium">Confidence Score</span>
                  <span class="text-body-2 font-weight-bold">
                    {{ (analysis.confidence_score * 100).toFixed(0) }}%
                  </span>
                </div>
                <v-progress-linear
                  :model-value="analysis.confidence_score * 100"
                  :color="confidenceColor"
                  height="8"
                  rounded
                />
              </div>
            </v-col>

            <v-col v-if="analysis.key_factors" cols="12" sm="6">
              <div class="metric-card">
                <div class="text-body-2 font-weight-medium mb-2">Key Factors</div>
                <div class="d-flex flex-wrap ga-1">
                  <v-chip
                    v-for="factor in analysis.key_factors"
                    :key="factor"
                    size="small"
                    variant="outlined"
                    color="primary"
                  >
                    {{ factor }}
                  </v-chip>
                </div>
              </div>
            </v-col>
          </v-row>
        </div>

        <!-- Risk Assessment -->
        <div v-if="analysis.risk_assessment" class="risk-assessment mt-6">
          <v-divider class="mb-4" />
          <div class="d-flex align-start ga-3">
            <v-icon color="warning" size="20">fas fa-exclamation-triangle</v-icon>
            <div>
              <div class="text-body-2 font-weight-medium mb-1">Risk Assessment</div>
              <div class="text-body-2 text-medium-emphasis">
                {{ analysis.risk_assessment }}
              </div>
            </div>
          </div>
        </div>

        <!-- Timestamp -->
        <div class="analysis-timestamp mt-6">
          <v-divider class="mb-3" />
          <div class="d-flex align-center ga-2 text-caption text-medium-emphasis">
            <v-icon size="14">fas fa-clock</v-icon>
            Generated {{ formatTimestamp(analysis.timestamp) }}
          </div>
        </div>
      </v-card-text>
    </v-card>

    <!-- Empty State -->
    <div v-else class="empty-state">
      <div class="text-center text-medium-emphasis">
        <v-icon size="48" class="mb-3">fas fa-chart-line</v-icon>
        <div class="text-h6 mb-1">Ready for Analysis</div>
        <div class="text-body-2">
          Your AI-powered analysis will appear here.
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
    'analysis-neutral': props.analysis?.recommendation === 'HOLD'
  }
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
  if (score >= 0.8) return 'success'
  if (score >= 0.6) return 'primary'
  if (score >= 0.4) return 'warning'
  return 'error'
})

function formatTimestamp(timestamp: string): string {
  return formatDistanceToNow(new Date(timestamp), { addSuffix: true })
}
</script>

<style scoped>
.analysis-result {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  font-family: 'Inter', system-ui, sans-serif;
}

.loading-state {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 300px;
}

.analysis-card {
  width: 100%;
}

.analysis-positive {
  border-left: 4px solid rgb(var(--v-theme-success));
}

.analysis-negative {
  border-left: 4px solid rgb(var(--v-theme-error));
}

.analysis-neutral {
  border-left: 4px solid rgb(var(--v-theme-warning));
}

.analysis-summary {
  background: rgba(var(--v-theme-primary-rgb), 0.05);
  border-radius: 8px;
  padding: 1rem;
  border: 1px solid rgba(var(--v-theme-primary-rgb), 0.1);
}

.metric-card {
  background: rgb(var(--v-theme-surface-variant));
  border-radius: 8px;
  padding: 1rem;
}

.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 300px;
}

.leading-relaxed {
  line-height: 1.6;
}
</style>
