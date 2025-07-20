<template>
  <v-card class="h-100">
    <v-card-title class="pa-4">
      <span class="text-lg font-weight-bold">{{ title }}</span>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="chart-container" style="width: 100%; height: 300px">
        <Line v-if="chartData" :data="chartData" :options="chartOptions" />
        <div v-else class="d-flex justify-center align-center h-100">
          <v-progress-circular indeterminate color="primary" />
        </div>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Line } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  type ChartOptions,
} from 'chart.js'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
)

interface LineChartProps {
  title: string
  data: Array<{
    date: string
    confidence: number
    volume: number
  }> | null
}

const props = defineProps<LineChartProps>()

const chartData = computed(() => {
  if (!props.data) return null

  return {
    labels: props.data.map((item) => item.date),
    datasets: [
      {
        label: 'Average Confidence',
        data: props.data.map((item) => item.confidence),
        borderColor: 'rgb(var(--v-theme-primary))',
        backgroundColor: 'rgba(var(--v-theme-primary-rgb), 0.1)',
        fill: true,
        tension: 0.4,
        pointBackgroundColor: 'rgb(var(--v-theme-primary))',
        pointBorderColor: 'white',
        pointBorderWidth: 2,
        pointRadius: 4,
      },
      {
        label: 'Recommendation Volume',
        data: props.data.map((item) => item.volume),
        borderColor: 'rgb(var(--v-theme-secondary))',
        backgroundColor: 'rgba(var(--v-theme-secondary-rgb), 0.1)',
        fill: true,
        tension: 0.4,
        pointBackgroundColor: 'rgb(var(--v-theme-secondary))',
        pointBorderColor: 'white',
        pointBorderWidth: 2,
        pointRadius: 4,
        yAxisID: 'y1',
      },
    ],
  }
})

const chartOptions: ChartOptions<'line'> = {
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    x: {
      grid: {
        display: true,
        color: 'rgba(var(--v-theme-outline-rgb), 0.2)',
      },
      ticks: {
        font: {
          family: 'Inter, system-ui, sans-serif',
        },
      },
    },
    y: {
      type: 'linear',
      display: true,
      position: 'left' as const,
      beginAtZero: true,
      max: 1,
      ticks: {
        callback: (value) => `${(Number(value) * 100).toFixed(0)}%`,
        font: {
          family: 'Inter, system-ui, sans-serif',
        },
      },
    },
    y1: {
      type: 'linear',
      display: true,
      position: 'right' as const,
      beginAtZero: true,
      grid: {
        drawOnChartArea: false,
      },
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
    },
  },
}
</script>
