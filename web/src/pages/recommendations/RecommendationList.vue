<template>
  <section class="recommendations-page">
    <div class="mb-8">
      <h1 class="text-h4 font-weight-bold mb-2">Broker Recommendations</h1>
      <p class="text-medium-emphasis">Browse the latest stock recommendations from top brokerages.</p>
    </div>

    <v-card>
      <v-card-title class="pb-0">Latest Recommendations</v-card-title>
      <v-card-subtitle v-if="store.pagination">
        Page {{ store.pagination.page }} of {{ totalPages }}
      </v-card-subtitle>

      <RecommendationFilters :loading="store.loading" @filters-changed="handleFiltersChanged" />

      <v-card-text>
        <v-data-table
          :items="store.recommendations"
          :headers="tableHeaders"
          :loading="store.loading"
          item-key="id"
          density="comfortable"
          no-data-text="No recommendations found"
          loading-text="Loading recommendations..."
        >
          <template #item.ticker="{ item }">
            <RouterLink :to="`/stocks/${item.ticker}`" class="ticker-link font-mono">
              {{ item.ticker }}
            </RouterLink>
          </template>

          <template #item.recommendation="{ item }">
            <RecommendationBadge :recommendation="item.recommendation" />
          </template>

          <template #item.confidence="{ item }">
            <span class="font-mono">{{ (item.confidence * 100).toFixed(0) }}%</span>
          </template>

          <template #item.time="{ item }">
            {{ formatDate(item.time) }}
          </template>

          <template #item.actions="{ item }">
            <RouterLink :to="`/recommendations/${item.id}`">
              <v-btn icon variant="text" color="primary" size="small">
                <v-icon size="18">fas fa-arrow-right</v-icon>
              </v-btn>
            </RouterLink>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Pagination Controls -->
    <div class="d-flex justify-space-between align-center mt-6">
      <v-btn variant="outlined" :disabled="store.filters.page === 1" @click="changePage(-1)">
        Previous
      </v-btn>

      <span class="text-body-2">
        Showing {{ showingStart }} - {{ showingEnd }} of {{ store.pagination?.total || 0 }}
      </span>

      <v-btn variant="outlined" :disabled="!canGoNext" @click="changePage(1)"> Next </v-btn>
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
import { onMounted, computed } from 'vue'
import { useRecommendations } from '@/stores/recommendations'
import { format } from 'date-fns'
import RecommendationFilters from '@/components/recommendations/RecommendationFilters.vue'
import RecommendationBadge from '@/components/recommendations/RecommendationBadge.vue'

const store = useRecommendations()

const tableHeaders = [
  { title: 'Ticker', value: 'ticker', width: 100 },
  { title: 'Brokerage', value: 'brokerage' },
  { title: 'Recommendation', value: 'recommendation', width: 140 },
  { title: 'Confidence', value: 'confidence', align: 'end' as const, width: 100 },
  { title: 'Date', value: 'time', align: 'end' as const, width: 120 },
  { title: '', value: 'actions', width: 60, sortable: false },
]

const totalPages = computed(() =>
  store.pagination ? Math.ceil(store.pagination.total / store.pagination.limit) : 1,
)

const showingStart = computed(() =>
  store.pagination ? (store.pagination.page - 1) * store.pagination.limit + 1 : 0,
)

const showingEnd = computed(() =>
  store.pagination
    ? Math.min(store.pagination.page * store.pagination.limit, store.pagination.total)
    : 0,
)

const canGoNext = computed(() =>
  store.pagination ? store.filters.page * store.filters.limit < store.pagination.total : false,
)

function formatDate(dateString: string): string {
  return format(new Date(dateString), 'MMM d, yyyy')
}

function handleFiltersChanged(filters: any) {
  store.applyFilters({ ...filters, page: 1 })
}

function changePage(delta: number) {
  store.filters.page += delta
  store.fetchRecommendations()
}

onMounted(() => {
  if (!store.recommendations.length) {
    store.fetchRecommendations()
  }
})
</script>

<style scoped>
.recommendations-page {
  max-width: 1200px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}

.ticker-link {
  color: var(--v-theme-primary);
  text-decoration: none;
  font-weight: 600;
}

.ticker-link:hover {
  text-decoration: underline;
}

.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
