<template>
  <v-card class="h-100">
    <v-card-title class="pa-4">
      <span class="text-lg font-weight-bold">{{ title }}</span>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="chart-container" style="width: 100%; height: 400px">
        <Bar v-if="chartData" :data="chartData" :options="chartOptions" />
        <div v-else class="d-flex justify-center align-center h-100">
          <v-progress-circular indeterminate color="primary" />
        </div>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  type ChartOptions,
} from 'chart.js'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

interface BarChartProps {
  title: string
  data: Array<{
    name: string
    buy: number
    sell: number
    hold: number
  }> | null
}

const props = defineProps<BarChartProps>()

const chartData = computed(() => {
  if (!props.data) return null

  return {
    labels: props.data.map((item) => item.name),
    datasets: [
      {
        label: 'Buy',
        data: props.data.map((item) => item.buy),
        backgroundColor: 'rgba(var(--v-theme-success-rgb), 0.8)',
        borderColor: 'rgb(var(--v-theme-success))',
        borderWidth: 1,
      },
      {
        label: 'Sell',
        data: props.data.map((item) => item.sell),
        backgroundColor: 'rgba(var(--v-theme-warning-rgb), 0.8)',
        borderColor: 'rgb(var(--v-theme-warning))',
        borderWidth: 1,
      },
      {
        label: 'Hold',
        data: props.data.map((item) => item.hold),
        backgroundColor: 'rgba(var(--v-theme-info-rgb), 0.8)',
        borderColor: 'rgb(var(--v-theme-info))',
        borderWidth: 1,
      },
    ],
  }
})

const chartOptions: ChartOptions<'bar'> = {
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    x: {
      stacked: true,
      grid: {
        display: false,
      },
      ticks: {
        font: {
          family: 'Inter, system-ui, sans-serif',
        },
      },
    },
    y: {
      stacked: true,
      beginAtZero: true,
      ticks: {
        font: {
          family: 'Inter, system-ui, sans-serif',
        },
      },
    },
  },
  plugins: {
    legend: {
      position: 'top' as const,
      labels: {
        font: {
          family: 'Inter, system-ui, sans-serif',
          size: 12,
        },
        usePointStyle: true,
      },
    },
    tooltip: {
      backgroundColor: 'rgb(var(--v-theme-surface))',
      titleColor: 'rgb(var(--v-theme-on-surface))',
      bodyColor: 'rgb(var(--v-theme-on-surface))',
      borderColor: 'rgb(var(--v-theme-outline))',
      borderWidth: 1,
      titleFont: {
        family: 'Inter, system-ui, sans-serif',
      },
      bodyFont: {
        family: 'Inter, system-ui, sans-serif',
      },
    },
  },
}
</script>
