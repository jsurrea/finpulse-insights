<template>
  <div class="stock-table">
    <!-- Desktop Table -->
    <div class="d-none d-md-block">
      <v-table density="compact">
        <thead>
          <tr>
            <th class="text-left font-weight-bold">Ticker</th>
            <th class="text-left font-weight-bold">Company</th>
            <th class="text-left font-weight-bold">Brokerage</th>
            <th class="text-center" style="width: 50px"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(stock, index) in stocks" :key="`${stock.ticker}-${stock.brokerage}-${index}`">
            <td class="font-weight-medium font-mono">{{ stock.ticker }}</td>
            <td>{{ stock.company }}</td>
            <td>{{ stock.brokerage }}</td>
            <td class="text-center">
              <v-btn
                icon
                variant="text"
                size="small"
                color="primary"
                :to="`/stocks/${stock.ticker}`"
              >
                <v-icon>fas fa-arrow-right</v-icon>
              </v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>
    </div>

    <!-- Mobile Cards -->
    <div class="d-block d-md-none">
      <v-card
        v-for="(stock, index) in stocks"
        :key="`${stock.ticker}-${stock.brokerage}-${index}`"
        class="mb-3"
        elevation="2"
      >
        <v-card-text>
          <div class="d-flex justify-space-between align-center">
            <div>
              <div class="font-weight-bold font-mono text-h6">{{ stock.ticker }}</div>
              <div class="text-body-2 text-medium-emphasis">{{ stock.company }}</div>
              <div class="text-body-2 text-medium-emphasis mt-1">{{ stock.brokerage }}</div>
            </div>
            <v-btn icon variant="text" color="primary" :to="`/stocks/${stock.ticker}`">
              <v-icon>fas fa-arrow-right</v-icon>
            </v-btn>
          </div>
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Stock } from '@/utils/types'

interface StockTableProps {
  stocks: Stock[]
}

defineProps<StockTableProps>()
</script>

<style scoped>
.stock-table {
  font-family: 'Inter', system-ui, sans-serif;
}

.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
