<template>
  <v-card class="h-100">
    <v-card-title class="pa-4">
      <span class="text-lg font-weight-bold">Recommendation Distribution</span>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="chart-container" style="width: 100%; height: 320px">
        <Doughnut :data="chartData" :options="chartOptions" :plugins="chartPlugins" />
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
        'rgb(var(--v-theme-success))',
        'rgb(var(--v-theme-warning))',
        'rgb(var(--v-theme-info))',
      ],
      borderColor: [
        'rgb(var(--v-theme-success))',
        'rgb(var(--v-theme-warning))',
        'rgb(var(--v-theme-info))',
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
      position: 'bottom' as const,
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
      backgroundColor: 'rgb(var(--v-theme-surface))',
      titleColor: 'rgb(var(--v-theme-on-surface))',
      bodyColor: 'rgb(var(--v-theme-on-surface))',
      borderColor: 'rgb(var(--v-theme-outline))',
      borderWidth: 1,
    },
  },
  cutout: '60%',
}

const chartPlugins = [
  {
    id: 'centerText',
    beforeDraw: (chart: ChartJS) => {
      const { ctx, width, height } = chart
      ctx.restore()

      const fontSize = Math.min(width, height) / 20
      ctx.font = `${fontSize}px Inter, system-ui, sans-serif`
      ctx.textBaseline = 'middle'
      ctx.fillStyle = 'rgb(var(--v-theme-on-surface))'

      const text = 'Recommendations'
      const textX = Math.round((width - ctx.measureText(text).width) / 2)
      const textY = height / 2

      ctx.fillText(text, textX, textY)
      ctx.save()
    },
  },
]
</script>

<style scoped>
.chart-container {
  position: relative;
}
</style>
