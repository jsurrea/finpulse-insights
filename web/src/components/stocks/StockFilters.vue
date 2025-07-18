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
          :loading="loading"
          hide-details
        />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <v-select
          v-model="brokerage"
          :items="brokerageOptions"
          placeholder="Filter by brokerage"
          variant="outlined"
          density="compact"
          :loading="loading"
          hide-details
        />
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useDebounceFn } from '@vueuse/core'

interface StockFiltersProps {
  loading?: boolean
}

const props = defineProps<StockFiltersProps>()
const emit = defineEmits<{
  'update:search': [value: string]
  'update:brokerage': [value: string]
}>()

const router = useRouter()
const route = useRoute()

const search = ref(route.query.search as string || '')
const brokerage = ref(route.query.brokerage as string || 'all')

const brokerageOptions = [
  { title: 'All Brokerages', value: 'all' },
  { title: 'Evercore ISI', value: 'Evercore ISI' },
  { title: 'B. Riley', value: 'B. Riley' },
  { title: 'Barclays', value: 'Barclays' },
  { title: 'Royal Bank Of Canada', value: 'Royal Bank Of Canada' },
  { title: 'Rosenblatt Securities', value: 'Rosenblatt Securities' },
  { title: 'Sidoti', value: 'Sidoti' },
  { title: 'Mizuho', value: 'Mizuho' },
  { title: 'Wedbush', value: 'Wedbush' },
  { title: 'HC Wainwright', value: 'HC Wainwright' },
  { title: 'Morgan Stanley', value: 'Morgan Stanley' },
  { title: 'The Goldman Sachs Group', value: 'The Goldman Sachs Group' },
  { title: 'Citigroup', value: 'Citigroup' },
  { title: 'Raymond James', value: 'Raymond James' },
  { title: 'KeyCorp', value: 'KeyCorp' },
  { title: 'UBS Group', value: 'UBS Group' },
  { title: 'Bank of America', value: 'Bank of America' },
  { title: 'Piper Sandler', value: 'Piper Sandler' },
  { title: 'Truist Financial', value: 'Truist Financial' },
  { title: 'Wells Fargo & Company', value: 'Wells Fargo & Company' }
]

const updateFilters = useDebounceFn(() => {
  const query = { ...route.query, page: '1' }

  if (search.value) {
    query.search = search.value
  } else {
    delete query.search
  }

  if (brokerage.value && brokerage.value !== 'all') {
    query.brokerage = brokerage.value
  } else {
    delete query.brokerage
  }

  router.push({ query })
}, 500)

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
