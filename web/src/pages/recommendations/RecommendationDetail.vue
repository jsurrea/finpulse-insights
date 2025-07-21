<template>
  <section class="recommendation-detail" v-if="store.currentRec">
    <div class="mb-6">
      <RouterLink to="/recommendations">
        <v-btn variant="outlined" class="mb-4" prepend-icon="fas fa-arrow-left">
          Back to Recommendations
        </v-btn>
      </RouterLink>
      <h1 class="text-h4 font-weight-bold">Recommendation Details</h1>
    </div>

    <v-card class="mb-6">
      <v-card-title class="d-flex justify-space-between align-center">
        <div>
          <span class="text-h5">{{ store.currentRec.company }}</span>
          <span class="font-mono ml-2">({{ store.currentRec.ticker }})</span>
        </div>
        <RecommendationBadge :recommendation="store.currentRec.recommendation" />
      </v-card-title>

      <v-card-subtitle>
        Recommendation issued because: {{ store.currentRec.reason }}
      </v-card-subtitle>

      <v-card-text>
        <v-row class="mb-2">
          <v-col cols="12" md="6" lg="4">
            <InfoItem
              icon="fas fa-building"
              label="Brokerage"
              :value="store.currentRec.brokerage"
            />
          </v-col>
          <v-col cols="12" md="6" lg="4">
            <InfoItem
              icon="fas fa-calendar-alt"
              label="Date"
              :value="formatDetailDateTime(store.currentRec.time)"
            />
          </v-col>
          <v-col cols="12" md="6" lg="4">
            <InfoItem icon="fas fa-tag" label="Action" :value="store.currentRec.action" />
          </v-col>
          <v-col cols="12" md="6" lg="4">
            <InfoItem
              icon="fas fa-chart-line"
              label="Rating Change"
              :value="`${store.currentRec.rating_from} → ${store.currentRec.rating_to}`"
            />
          </v-col>
          <v-col cols="12" md="6" lg="4">
            <InfoItem
              icon="fas fa-bullseye"
              label="Target Price"
              :value="
                store.currentRec.target_to ? `$${store.currentRec.target_to.toFixed(2)}` : 'N/A'
              "
            />
          </v-col>
          <v-col cols="12" md="6" lg="4">
            <InfoItem
              icon="fas fa-info-circle"
              label="Recommendation"
              :value="store.currentRec?.recommendation"
            />
          </v-col>
        </v-row>

        <!-- Confidence Gauge -->
        <v-row>
          <v-col cols="12">
            <ConfidenceGauge :confidence="store.currentRec.confidence" />
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <RouterLink :to="`/stocks/${store.currentRec.ticker}`">
      <v-btn color="primary" append-icon="fas fa-arrow-right"> View Stock Details </v-btn>
    </RouterLink>

    <!-- Loading State -->
    <div v-if="store.loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" />
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
import { onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useRecommendations } from '@/stores/recommendations'
import { format } from 'date-fns'
import RecommendationBadge from '@/components/recommendations/RecommendationBadge.vue'
import ConfidenceGauge from '@/components/recommendations/ConfidenceGauge.vue'
import InfoItem from '@/components/recommendations/InfoItem.vue'

const route = useRoute()
const store = useRecommendations()
const id = route.params.id as string

function formatDetailDateTime(dateString: string): string {
  return format(new Date(dateString), 'PPP p')
}

onMounted(() => {
  store.fetchRecommendationDetail(id)
})

onUnmounted(() => {
  store.clearCurrent()
})
</script>

<style scoped>
.recommendation-detail {
  max-width: 1000px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}

.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
