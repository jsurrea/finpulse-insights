<template>
  <div class="stocks-page">
    <div class="mb-8">
      <h1 class="text-3xl font-weight-bold tracking-tight mb-2">Stocks</h1>
      <p class="text-medium-emphasis">Search and browse stocks tracked by FinPulse.</p>
    </div>

    <v-card elevation="4">

      <v-card-text>
        <StockFilters :loading="loading" />

        <div class="mt-6">
          <StockTable :stocks="stocks" />
        </div>
      </v-card-text>
    </v-card>

    <!-- Pagination -->
    <div class="d-flex justify-space-between align-center mt-8">
      <v-btn variant="outlined" :disabled="currentPage <= 1" @click="goToPage(currentPage - 1)">
        Previous
      </v-btn>

      <span class="text-sm text-medium-emphasis">
        Showing page {{ currentPage }} of {{ totalPages }}
      </span>

      <v-btn
        variant="outlined"
        :disabled="currentPage >= totalPages"
        @click="goToPage(currentPage + 1)"
      >
        Next
      </v-btn>
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
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStocks } from '@/stores/stocks'
import StockFilters from '@/components/stocks/StockFilters.vue'
import StockTable from '@/components/stocks/StockTable.vue'
import type { Stock } from '@/utils/types'

const route = useRoute()
const router = useRouter()
const stocksStore = useStocks()

const stocks = ref<Stock[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

const currentPage = computed(() => Number(route.query.page) || 1)
const totalPages = computed(() =>
  stocksStore.pagination
    ? Math.ceil(stocksStore.pagination.total / stocksStore.pagination.limit)
    : 1,
)

const loadStocks = async () => {
  try {
    loading.value = true
    error.value = null

    const search = (route.query.search as string) || ''
    const brokerage = (route.query.brokerage as string) || 'all'

    stocksStore.applyFilters({
      page: currentPage.value,
      limit: 10,
      search,
      brokerage,
    })

    await stocksStore.fetchStocks()
    stocks.value = stocksStore.stocks
  } catch (err: unknown) {
    if (err instanceof Error) {
      error.value = err.message
      console.error('Stocks loading error:', err)
    } else {
      error.value = 'Error loading stocks'
      console.error('Stocks loading error:', err)
    }
  } finally {
    loading.value = false
  }
}

const goToPage = (page: number) => {
  const query = { ...route.query, page: page.toString() }
  router.push({ query })
}

// Watch for route changes
watch(() => route.query, loadStocks, { immediate: true })

onMounted(() => {
  loadStocks()
})
</script>

<style scoped>
.stocks-page {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
