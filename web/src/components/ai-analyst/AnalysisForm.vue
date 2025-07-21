<template>
  <v-card class="analysis-form">
    <v-card-title class="pb-2">
      <span class="text-lg font-weight-bold">Analyze Stock</span>
    </v-card-title>
    <v-card-subtitle class="pb-4">
      Select a stock from your portfolio to begin AI analysis.
    </v-card-subtitle>

    <v-card-text>
      <v-form @submit.prevent="onSubmit">
        <div class="form-fields">
          <v-autocomplete
            v-model="selectedStock"
            :items="stockOptions"
            item-title="displayName"
            item-value="ticker"
            label="Select Stock"
            placeholder="Search by ticker or company name..."
            variant="outlined"
            density="comfortable"
            :loading="store.loadingSearch"
            :disabled="loading"
            :error-messages="error"
            clearable
            no-data-text="No stocks found"
            class="mb-4"
            @update:search="onSearchInput"
          >
            <template #item="{ props, item }">
              <v-list-item v-bind="props">
                <template #prepend>
                  <v-avatar size="32" color="primary" class="mr-3">
                    <span class="text-caption font-weight-bold">
                      {{ item.raw.ticker.substring(0, 2) }}
                    </span>
                  </v-avatar>
                </template>
                <template #append>
                  <v-chip size="x-small" variant="outlined" color="primary">
                    {{ item.raw.brokerage }}
                  </v-chip>
                </template>
              </v-list-item>
            </template>
          </v-autocomplete>

          <!-- Información del stock seleccionado -->
          <v-card
            v-if="selectedStockInfo"
            variant="outlined"
            class="mb-6"
            color="primary"
          >
            <v-card-text class="py-3">
              <div class="d-flex align-center justify-space-between">
                <div>
                  <div class="font-weight-bold text-h6">{{ selectedStockInfo.ticker }}</div>
                  <div class="text-body-2 text-medium-emphasis">{{ selectedStockInfo.company }}</div>
                </div>
                <v-chip color="primary" variant="tonal" size="small">
                  {{ selectedStockInfo.brokerage }}
                </v-chip>
              </div>
            </v-card-text>
          </v-card>

          <v-btn
            type="submit"
            color="primary"
            block
            size="large"
            :loading="loading"
            :disabled="!selectedStock || loading"
          >
            <template #prepend>
              <v-icon>fas fa-sparkles</v-icon>
            </template>
            {{ loading ? 'Analyzing...' : 'Generate AI Analysis' }}
          </v-btn>
        </div>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAIAnalyst } from '@/stores/aiAnalyst'

const emit = defineEmits<{
  'analysis-submitted': [ticker: string]
}>()

const route = useRoute()
const store = useAIAnalyst()

const selectedStock = ref<string>('')
const error = ref<string>('')

const loading = computed(() => store.loading)

const stockOptions = computed(() =>
  store.stockOptions.map(stock => ({
    ...stock,
    displayName: `${stock.ticker} - ${stock.company}`
  }))
)

const selectedStockInfo = computed(() =>
  store.selectedStockInfo || store.stockOptions.find(s => s.ticker === selectedStock.value)
)

const onSearchInput = (searchValue: string) => {
  store.searchStocks(searchValue || '')
}

const onSubmit = async () => {
  error.value = ''

  if (!selectedStock.value) {
    error.value = 'Please select a stock to analyze'
    return
  }

  emit('analysis-submitted', selectedStock.value)
}

// Cargar opciones iniciales al montar el componente
onMounted(async () => {
  await store.loadStockOptions()

  // Set from URL params if present
  if (route.query.ticker) {
    selectedStock.value = route.query.ticker as string
    store.setSelectedStock(route.query.ticker as string)
  }
})

// Limpiar el timer de debounce al desmontar
onUnmounted(() => {
  store.clearDebounceTimer()
  store.clearSelectedStock()
})

// Watch for changes in selected stock
watch(selectedStock, (newValue) => {
  error.value = ''
  if (newValue) {
    store.setSelectedStock(newValue)
  } else {
    store.clearSelectedStock()
  }
})
</script>

<style scoped>
.analysis-form {
  font-family: 'Inter', system-ui, sans-serif;
}

.form-fields {
  display: flex;
  flex-direction: column;
}
</style>
