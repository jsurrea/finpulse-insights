<template>
  <div class="recommendation-filters">
    <v-row class="pt-4">
      <v-col cols="12" sm="6" md="3">
        <v-text-field
          v-model="ticker"
          label="Ticker"
          placeholder="e.g. AAPL"
          prepend-inner-icon="fas fa-search"
          clearable
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-text-field
          v-model="company"
          label="Company"
          placeholder="Company name"
          clearable
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-select
          v-model="brokerage"
          :items="brokerageOptions"
          item-title="title"
          item-value="value"
          label="Brokerage"
          variant="outlined"
          density="compact"
          :loading="loadingBrokerages"
          hide-details
          clearable
        />
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-select
          v-model="action"
          :items="actionOptions"
          item-title="title"
          item-value="value"
          label="Action"
          variant="outlined"
          density="compact"
          hide-details
          clearable
        />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-text-field
          v-model="dateFrom"
          label="Date From"
          type="date"
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-text-field
          v-model="dateTo"
          label="Date To"
          type="date"
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-select
          v-model="sort"
          :items="sortOptions"
          item-title="title"
          item-value="value"
          label="Sort By"
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useDebounceFn } from '@vueuse/core'
import { useStocks } from '@/stores/stocks'

type RecommendationFilters = {
  ticker: string
  company: string
  brokerage: string
  action: string
  date_from: string
  date_to: string
  sort: string
}

const emit = defineEmits<{
  'filters-changed': [filters: RecommendationFilters]
}>()

const stocksStore = useStocks()
const router = useRouter()
const route = useRoute()

const ticker = ref((route.query.ticker as string) || '')
const company = ref((route.query.company as string) || '')
const brokerage = ref((route.query.brokerage as string) || '')
const action = ref((route.query.action as string) || '')
const dateFrom = ref((route.query.date_from as string) || '')
const dateTo = ref((route.query.date_to as string) || '')
const sort = ref((route.query.sort as string) || 'time:desc')

const loadingBrokerages = ref(false)

const brokerageOptions = computed(() => [
  { title: 'All Brokerages', value: '' },
  ...stocksStore.brokerages.map((b) => ({ title: b, value: b })),
])

const actionOptions = [
  { title: 'All Actions', value: '' },
  { title: 'Upgrade', value: 'upgrade' },
  { title: 'Downgrade', value: 'downgrade' },
  { title: 'Initiate', value: 'initiate' },
  { title: 'Maintain', value: 'maintain' },
  { title: 'Reiterate', value: 'reiterate' },
]

const sortOptions = [
  { title: 'Date (Newest)', value: 'time:desc' },
  { title: 'Date (Oldest)', value: 'time:asc' },
  { title: 'Ticker (A-Z)', value: 'ticker:asc' },
  { title: 'Ticker (Z-A)', value: 'ticker:desc' },
  { title: 'Confidence (High)', value: 'confidence:desc' },
  { title: 'Confidence (Low)', value: 'confidence:asc' },
]

const updateFilters = useDebounceFn(() => {
  const query: Record<string, string> = { page: '1' }

  if (ticker.value) query.ticker = ticker.value
  if (company.value) query.company = company.value
  if (brokerage.value) query.brokerage = brokerage.value
  if (action.value) query.action = action.value
  if (dateFrom.value) query.date_from = dateFrom.value
  if (dateTo.value) query.date_to = dateTo.value
  if (sort.value) query.sort = sort.value

  router.push({ query })

  emit('filters-changed', {
    ticker: ticker.value,
    company: company.value,
    brokerage: brokerage.value,
    action: action.value,
    date_from: dateFrom.value,
    date_to: dateTo.value,
    sort: sort.value,
  })
}, 500)

onMounted(async () => {
  if (!stocksStore.brokerages.length) {
    loadingBrokerages.value = true
    await stocksStore.fetchBrokerages?.()
    loadingBrokerages.value = false
  }
})

watch([ticker, company, brokerage, action, dateFrom, dateTo, sort], () => {
  updateFilters()
})
</script>

<style scoped>
.recommendation-filters {
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
