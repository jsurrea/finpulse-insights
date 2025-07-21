<template>
  <div class="stock-filters">
    <v-row class="pt-4">
      <v-col cols="12" sm="6" md="4">
        <v-text-field
          v-model="search"
          placeholder="Search by ticker or company..."
          prepend-inner-icon="fas fa-search"
          clearable
          variant="outlined"
          density="compact"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-select
          v-model="brokerage"
          :items="brokerageItems"
          item-title="title"
          item-value="value"
          placeholder="Filter by brokerage"
          variant="outlined"
          density="compact"
          :loading="loadingBrokerages"
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

const emit = defineEmits<{
  'update:search': [value: string]
  'update:brokerage': [value: string]
}>()

const stocksStore = useStocks()
const router = useRouter()
const route = useRoute()

const search = ref((route.query.search as string) || '')
const brokerage = ref((route.query.brokerage as string) || 'all')

const loadingBrokerages = ref(false)
const brokerageItems = computed(() => [
  { title: 'All Brokerages', value: 'all' },
  ...stocksStore.brokerages.map((b) => ({ title: b, value: b })),
])

const updateFilters = useDebounceFn(() => {
  const query: Record<string, string> = { ...route.query, page: '1' }

  if (search.value) query.search = search.value
  else delete query.search

  if (brokerage.value && brokerage.value !== 'all') query.brokerage = brokerage.value
  else delete query.brokerage

  router.push({ query })
}, 500)

onMounted(async () => {
  if (!stocksStore.brokerages.length) {
    loadingBrokerages.value = true
    await stocksStore.fetchBrokerages?.()
    loadingBrokerages.value = false
  }
})

watch([search, brokerage], () => {
  emit('update:search', search.value)
  emit('update:brokerage', brokerage.value)
  updateFilters()
})
</script>

<style scoped>
.stock-filters {
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
