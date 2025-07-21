<template>
  <v-card class="h-100">
    <v-card-title class="pa-4 mb-8">
      <span class="text-lg font-weight-bold">Recommendation Distribution</span>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="chart-container">
        <Doughnut :data="chartData" :options="chartOptions" />
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Doughnut } from 'vue-chartjs'
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend,
  type ChartOptions,
  type TooltipItem,
} from 'chart.js'
import type { AnalyticsSummary } from '@/utils/types'

ChartJS.register(ArcElement, Tooltip, Legend)

interface DistributionChartProps {
  data: AnalyticsSummary
}

const props = defineProps<DistributionChartProps>()

const chartData = computed(() => ({
  labels: ['Buy', 'Sell', 'Hold'],
  datasets: [
    {
      data: [props.data.buy_percentage, props.data.sell_percentage, props.data.hold_percentage],
      backgroundColor: [
        '#22C55E', // Buy
        '#F59E42', // Sell
        '#3B82F6', // Hold
      ],
      borderColor: [
        '#22C55E',
        '#F59E42',
        '#3B82F6',
      ],
      borderWidth: 2,
    },
  ],
}))

const chartOptions: ChartOptions<'doughnut'> = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      position: 'bottom',
      labels: {
        padding: 20,
        usePointStyle: true,
        font: {
          family: 'Inter, system-ui, sans-serif',
          size: 14,
        },
      },
    },
    tooltip: {
      callbacks: {
        label: (context: TooltipItem<'doughnut'>) => {
          const label = context.label || ''
          const value = context.parsed
          return `${label}: ${value.toFixed(1)}%`
        },
      },
      backgroundColor: '#fff',
      titleColor: '#111',
      bodyColor: '#111',
      borderColor: '#999',
      borderWidth: 1,
    },
  },
  cutout: '60%',
}
</script>

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  height: 450px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
